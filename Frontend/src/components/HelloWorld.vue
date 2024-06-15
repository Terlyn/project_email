<script setup>
import { ref } from 'vue'

defineProps({
  msg: String,
})

const count = ref(0)


export default {
  data() {
    return {
      emails: [
        {
          subject: "changes made on 12/04 to 11/30",
          from: "valarie.sabo@enron.com",
          to: "stacey.white@enron.com",
          body: `The gas book liquidated twice, so our QTD and YTD was understated...`
        },
        {
          subject: "Enron Mentions - 03-04-01",
          from: "john.arnold@enron.com",
          to: "greg.whalley@enron.com",
          body: `In numbers, LT books QTD and YTD increased on 11/30 by approximately...`
        }
      ],
      selectedEmail: {}
    };
  },
  methods: {
    selectEmail(email) {
      this.selectedEmail = email;
    }
  },
  mounted() {
    this.selectedEmail = this.emails[0];
  }
};
</script>

<template>
<div class="container mx-auto p-6 bg-pink-100 min-h-screen">
    <div class="flex flex-col bg-white shadow-md rounded-lg overflow-hidden">
      <!-- Search Bar -->
      <div class="flex items-center p-4 border-b border-pink-200 bg-pink-50">
        <input
          type="text"
          placeholder="Search"
          class="border border-pink-300 p-2 flex-grow rounded-lg focus:outline-none focus:ring-2 focus:ring-pink-500 text-sm"
        />
        <button class="bg-pink-500 text-white p-2 ml-2 rounded-lg hover:bg-pink-600 focus:outline-none focus:ring-2 focus:ring-pink-500 text-sm">
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
              @click="selectEmail(email)"
              class="hover:bg-pink-100 cursor-pointer"
            >
              <td class="px-6 py-4 border border-pink-200 whitespace-nowrap">{{ email.subject }}</td>
              <td class="px-6 py-4 border border-pink-200 whitespace-nowrap">{{ email.from }}</td>
              <td class="px-6 py-4 border border-pink-200 whitespace-nowrap">{{ email.to }}</td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Selected Email Details -->
      <div class="p-4 border-t border-pink-200 bg-pink-50 text-sm">
        <h2 class="font-bold text-lg mb-2 text-pink-600">{{ selectedEmail.subject }}</h2>
        <p>{{ selectedEmail.body }}</p>
      </div>
    </div>
  </div>






  <h1>{{ msg }}</h1>

  <div class="card">
    <button type="button" @click="count++">count is {{ count }}</button>
    <p>
      Edit
      <code>components/HelloWorld.vue</code> to test HMR
    </p>
  </div>

  <p>
    Check out
    <a href="https://vuejs.org/guide/quick-start.html#local" target="_blank"
      >create-vue</a
    >, the official Vue + Vite starter
  </p>
  <p>
    Install
    <a href="https://github.com/vuejs/language-tools" target="_blank">Volar</a>
    in your IDE for a better DX
  </p>
  <p class="read-the-docs">Click on the Vite and Vue logos to learn more</p>
</template>

<style scoped>
.read-the-docs {
  color: #888;
}
</style>
