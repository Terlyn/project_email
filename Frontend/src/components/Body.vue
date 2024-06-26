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
        <button @click.prevent="submitHandler" class="bg-pink-500 text-white p-2 ml-2 rounded-lg hover:bg-pink-600 focus:outline-none focus:ring-2 focus:ring-pink-500 text-sm">
          Search
        </button>
      </div>

      <!-- Emails Table -->
      <div class="overflow-x-auto" v-if="searched && emailsToShow.length > 0">
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
              v-for="(email, index) in emailsToShow"
              :key="index"
              :class="{'bg-pink-200': selectedEmailIndex === index}"
              class="hover:bg-pink-100 cursor-pointer"
              @click="showEmailBody(index)"
            >
              <td class="px-6 py-4 border border-pink-200 whitespace-nowrap">{{ email.subject }}</td>
              <td class="px-6 py-4 border border-pink-200 whitespace-nowrap">{{ email.from }}</td>
              <td class="px-6 py-4 border border-pink-200 whitespace-nowrap">{{ email.to }}</td>
            </tr>
          </tbody>
        </table>
        <div class="flex justify-center">
          <button v-if="remainingEmails > 0" @click="loadMore" class="bg-pink-500 text-white p-2 rounded-lg mt-4 hover:bg-pink-600 focus:outline-none focus:ring-2 focus:ring-pink-500 text-sm">
            <span class="loader"></span>
          </button>
        </div>
        <div class="flex justify-center">
          <button v-if="emailsToShow.length > 3" @click="showLess" class="bg-pink-500 text-white p-2 rounded-lg mt-4 hover:bg-pink-600 focus:outline-none focus:ring-2 focus:ring-pink-500 text-sm">
            Mostrar menos resultados
          </button>
        </div>
      </div>

      <!-- Display Email Body -->
      <div v-if="selectedEmailBody" class="p-6 bg-white shadow-md rounded-lg mt-4">
        <h1 class="text-lg font-bold mb-2">Email Body</h1>
        <p>{{ selectedEmailBody }}</p>
      </div>
      <div v-else-if="searched && emails.length === 0" class="p-6 bg-white shadow-md rounded-lg mt-4">
        <p>No se encontraron resultados.</p>
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
      selectedEmailBody: "",
      selectedEmailIndex: null,
      searched: false,
      emailsToShow: [] // Nuevo array para almacenar los emails mostrados
    };
  },
  computed: {
    remainingEmails() {
      return this.emails.length - this.emailsToShow.length;
    }
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

      if (this.searchQuery.trim() === "") {
        this.emails = [];
        this.searched = false;
        this.selectedEmailBody = "";
        this.selectedEmailIndex = null;
        this.emailsToShow = [];
        return;
      }

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
        console.log("Response from backend:", data); // Verificar la estructura de los datos

        if (data.error) {
          console.log("Error:", data.message);
        } else {
          this.emails = Array.isArray(data) ? data : []; // Asegurarse de que los datos sean un array
          this.searched = true;
          this.selectedEmailBody = "";
          this.selectedEmailIndex = null;
          this.showInitialResults();
        }
      } catch (error) {
        console.error("There was an error with the search:", error);
      }
    },
    showInitialResults() {
      this.emailsToShow = this.emails.slice(0, 3); // Mostrar solo los primeros 3 resultados
    },
    loadMore() {
      const currentLength = this.emailsToShow.length;
      const remainingEmails = this.emails.slice(currentLength, currentLength + 3);
      this.emailsToShow = [...this.emailsToShow, ...remainingEmails];
    },
    showLess() {
      this.emailsToShow = this.emails.slice(0, 3); // Mostrar solo los primeros 3 resultados después de cargar más
    },
    showEmailBody(index) {
      this.selectedEmailBody = this.emailsToShow[index].body;
      this.selectedEmailIndex = index;
    }
  }
}
</script>

<style scoped>
.container {
  max-width: 100%;
}
.bg-pink-200 {
  background-color: #fbcfe8; /* Ajusta este color según sea necesario */
}
.loader {
  border: 4px solid #f3f3f3; /* Light grey */
  border-top: 4px solid #3498db; /* Blue */
  border-radius: 50%;
  width: 16px;
  height: 16px;
  animation: spin 2s linear infinite;
}
@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}
</style>
