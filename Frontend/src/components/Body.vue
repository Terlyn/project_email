<template>
  <div class="container mx-auto p-6 bg-pink-100 min-h-screen">
    <div class="flex flex-col bg-white shadow-md rounded-lg overflow-hidden">
      <!-- Search Bar -->
      <div class="flex items-center p-4 border-b border-pink-200 bg-pink-50">
        <input
          v-model="searchQuery"
          type="text"
          placeholder="Search"
          class="border border-pink-300 p-2 flex-grow rounded-lg focus:outline-none focus:ring-2 focus:ring-pink-500 text-sm"
        />
        <button @click="submitHandler" class="bg-pink-500 text-white p-2 ml-2 rounded-lg hover:bg-pink-600 focus:outline-none focus:ring-2 focus:ring-pink-500 text-sm">
          Search
        </button>
      </div>

      <!-- Emails Table -->
      <div class="overflow-x-auto">
        <table class="min-w-full border border-pink-200">
          <thead class="bg-pink-50">
            <tr>
              <th class="px-6 py-3 border border-pink-200 text-left text-xs font-medium text-pink-600 uppercase tracking-wider">Subject</th>
              <th class="px-6 py-3 border border-pink-200 text-left text-xs font-medium text-pink-600 uppercase tracking-wider">From</th>
              <th class="px-6 py-3 border border-pink-200 text-left text-xs font-medium text-pink-600 uppercase tracking-wider">To</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-pink-200 text-sm">
            <tr
              v-for="(email, index) in emails"
              :key="index"
              class="hover:bg-pink-100 cursor-pointer"
              @click="showEmailBody(index)"
            >
              <td class="px-6 py-4 border border-pink-200 whitespace-nowrap">{{ email.subject }}</td>
              <td class="px-6 py-4 border border-pink-200 whitespace-nowrap">{{ email.from }}</td>
              <td class="px-6 py-4 border border-pink-200 whitespace-nowrap">{{ email.to }}</td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Display Email Body -->
      <div v-if="selectedEmailBody" class="p-6 bg-white shadow-md rounded-lg mt-4">
        <h1 class="text-lg font-bold mb-2">Email Body</h1>
        <p>{{ selectedEmailBody }}</p>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'search',
  props: {
    Emails: {
      type: Array,
      default: () => []
    }
  },
  data() {
    return {
      searchQuery: "",
      emails: this.Emails,
      selectedEmailBody: ""
    };
  },
  mounted() {
    if (this.emails.length === 0) {
      this.fetchEmails();
    }
  },
  methods: {
    async fetchEmails() {
      try {
        const response = await axios.get('http://localhost:8081/emails');
        this.emails = response.data;
      } catch (error) {
        console.error("There was an error fetching the emails:", error);
      }
    },
    async submitHandler() {
      console.log("submitHandler called - success!");

      const payload = {
        searchQuery: this.searchQuery,
      }

      try {
        const response = await fetch("http://localhost:8081/search", {
          method: "POST",
          body: JSON.stringify(payload),
          headers: {
            "Content-Type": "application/json"
          }
        });
        const data = await response.json();
        if (data.error) {
          console.log("Error:", data.message);
        } else {
          console.log(data);
        }
      } catch (error) {
        console.error("There was an error with the search:", error);
      }
    },
    showEmailBody(index) {
      this.selectedEmailBody = this.emails[index].body;
    }
  }
}
</script>

<style scoped>
.container {
  max-width: 100%;
}
</style>
