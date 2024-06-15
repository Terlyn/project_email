<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const emails = ref([])
const searchQuery = ref('')

const fetchEmails = async () => {
  try {
    const response = await axios.get('http://localhost:8080', {
      params: { q: searchQuery.value }
    })
    emails.value = response.data
  } catch (error) {
    console.error('Error fetching emails:', error)
  }
}

onMounted(() => {
  fetchEmails()
})

const searchEmails = () => {
  fetchEmails()
}

defineProps({
  Emails: Array,
})
</script>

<style scoped>
/* Estilos opcionales según tus necesidades */
</style>


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
        <button @click="searchEmails" class="bg-pink-500 text-white p-2 ml-2 rounded-lg hover:bg-pink-600 focus:outline-none focus:ring-2 focus:ring-pink-500 text-sm">
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
              v-for="(email, index) in Emails"
              :key="index"
              class="hover:bg-pink-100 cursor-pointer"
            >
              <td class="px-6 py-4 border border-pink-200 whitespace-nowrap">{{ email.subject }}</td>
              <td class="px-6 py-4 border border-pink-200 whitespace-nowrap">{{ email.from }}</td>
              <td class="px-6 py-4 border border-pink-200 whitespace-nowrap">{{ email.to }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>
