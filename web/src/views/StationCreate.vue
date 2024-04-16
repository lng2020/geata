<template> <div class="container mx-auto py-8"> <h1 class="text-3xl font-bold mb-6">Create Station</h1>
<form @submit.prevent="createStation" class="max-w-md">
  <div class="mb-4">
    <label for="name" class="block text-gray-700 font-bold mb-2">Station Name</label>
    <input
      type="text"
      id="name"
      v-model="stationName"
      class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500"
      required
    />
  </div>

  <div class="mb-4">
    <label for="host" class="block text-gray-700 font-bold mb-2">Host</label>
    <input
      type="text"
      id="host"
      v-model="host"
      class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500"
      required
    />
  </div>

  <div class="mb-4">
    <label for="port" class="block text-gray-700 font-bold mb-2">Port</label>
    <input
      type="number"
      id="port"
      v-model="port"
      class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500"
      required
    />
  </div>

  <div class="mb-4">
    <label for="icdFile" class="block text-gray-700 font-bold mb-2">ICD File</label>
    <input
      type="file"
      id="icdFile"
      @change="handleFileUpload"
      class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500"
      accept=".icd"
      required
    />
  </div>

  <button
    type="submit"
    class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded"
  >
    Create
  </button>
</form>
</div>
</template> 
<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
const router = useRouter()
const stationName = ref('')
const host = ref('')
const port = ref(0)
const icdFile = ref<File | null>(null) 
function handleFileUpload(event: Event) {
    const file = (event.target as HTMLInputElement).files?.[0] || null
    icdFile.value = file 
}
async function createStation() {
    try { 
        const formData = new FormData()
        formData.append('name', stationName.value)
        formData.append('host', host.value)
        formData.append('port', port.value.toString())
        if (icdFile.value) { 
            formData.append('icdFile', icdFile.value)
        } 
        // Send a POST request to the server to create the station 
        const response = await fetch('/api/stations', { method: 'POST', body: formData }) 
        if (response.ok) { 
            // Station created successfully 
            const data = await response.json()
            console.log('Station created:', data) 
            // Reset the form fields 
            stationName.value = '' 
            host.value = '' 
            port.value = 0 
            icdFile.value = null 
            // Navigate back to the dashboard 
            router.push('/') 
        } else { 
            // Handle error case 
            const errData = await response.json()
            console.error('Failed to create station:', errData) 
        } 
    } catch (error) { 
        console.error('Error creating station:', error)
    }
} 
</script>