<template>
  <div class="container mx-auto py-8 flex justify-center">
    <form @submit.prevent="createStation" class="max-w-md w-full">
      <h1 class="text-3xl font-bold mb-6 text-center">Create Station</h1>

      <div class="mb-8">
        <div v-show="step === 1" class="mb-8">
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
          <div class="flex justify-end">
            <button
              type="button"
              class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded"
              @click="nextStep"
            >
              Confirm
            </button>
          </div>
        </div>

        <div v-show="step === 2" class="mb-8">
          <div class="mb-4">
            <label for="icdFile" class="block text-gray-700 font-bold mb-2">ICD File</label>
            <div
              class="border-2 border-dashed border-gray-300 rounded-md p-4 text-center"
              @dragover.prevent
              @drop.prevent="handleFileDrop"
            >
              <input
                type="file"
                id="icdFile"
                @change="handleFileUpload"
                class="hidden"
                accept=".icd"
                required
              />
              <label for="icdFile" class="cursor-pointer">
                <span v-if="!icdFile">Click to select file or drag and drop</span>
                <span v-else>{{ icdFile.name }} ({{ formatFileSize(icdFile.size) }})</span>
              </label>
            </div>
          </div>
          <div class="mb-4">
            <label class="block text-gray-700 font-bold mb-2">Analysis Progress</label>
            <div class="bg-gray-200 rounded-full h-6">
              <div
                class="bg-blue-500 h-6 rounded-full transition-all duration-500 ease-in-out"
                :style="{ width: `${progress}%` }"
              ></div>
            </div>
            <div class="text-right mt-2">{{ progress }}%</div>
          </div>
          <div class="flex justify-end">
            <button
              type="button"
              class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded"
              @click="nextStep"
            >
              Confirm
            </button>
          </div>
        </div>

        <div v-show="step === 3" class="mb-8">
          <div v-if="icdData" class="mb-4">
            <h3 class="text-xl font-bold mb-2">ICD Structure Preview</h3>
            <pre>{{ JSON.stringify(icdData, null, 2) }}</pre>
          </div>
          <div v-if="icdStats" class="mb-4">
            <h3 class="text-xl font-bold mb-2">ICD Statistics</h3>
            <table class="table-auto w-full">
              <thead>
                <tr>
                  <th class="px-4 py-2">Item</th>
                  <th class="px-4 py-2">Count</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="(value, key) in icdStats" :key="key">
                  <td class="border px-4 py-2">{{ key }}</td>
                  <td class="border px-4 py-2">{{ value }}</td>
                </tr>
              </tbody>
            </table>
          </div>
          <div class="flex justify-end">
            <button
              type="submit"
              class="bg-green-500 hover:bg-green-600 text-white font-bold py-2 px-4 rounded"
            >
              Create
            </button>
          </div>
        </div>
      </div>

      <div class="flex justify-center items-center mt-8">
        <div class="flex items-center">
          <span
            :class="[
              'w-4 h-4 rounded-full cursor-pointer',
              step === 1 ? 'bg-blue-500' : 'bg-gray-300'
            ]"
            @click="step = 1"
          ></span>
          <div class="w-20 h-1 bg-gray-300"></div>
          <span
            :class="[
              'w-4 h-4 rounded-full cursor-pointer',
              step === 2 ? 'bg-blue-500' : 'bg-gray-300'
            ]"
            @click="step = 2"
          ></span>
          <div class="w-20 h-1 bg-gray-300"></div>
          <span
            :class="[
              'w-4 h-4 rounded-full cursor-pointer',
              step === 3 ? 'bg-blue-500' : 'bg-gray-300'
            ]"
            @click="step = 3"
          ></span>
        </div>
      </div>
    </form>
  </div>
</template>
<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const step = ref(1)
const stationName = ref('')
const host = ref('')
const port = ref(0)
const icdFile = ref<File | null>(null)
const progress = ref(0)
const fileHashName = ref('')
const IEC61850Model = ref('')

function handleFileUpload(event: Event) {
  const target = event.target as HTMLInputElement
  const file = target.files?.[0]
  if (file) {
    icdFile.value = file
    uploadICDFile(file)
  }
}

function handleFileDrop(event: DragEvent) {
  event.preventDefault()
  const file = event.dataTransfer?.files[0]
  if (file) {
    icdFile.value = file
    uploadICDFile(file)
  }
}

function formatFileSize(size: number) {
  const units = ['B', 'KB', 'MB', 'GB']
  let index = 0
  while (size >= 1024 && index < units.length - 1) {
    size /= 1024
    index++
  }
  return `${size.toFixed(2)} ${units[index]}`
}

async function nextStep() {
  if (step.value < 3) {
    step.value++
  }
}

async function uploadICDFile(file: File) {
  const formData = new FormData()
  formData.append('file', file)

  try {
    const response = await fetch('/api/upload', {
      method: 'POST',
      body: formData
    })
    if (response.ok) {
      const data = await response.json()
      console.log('File uploaded:', data)
      fileHashName.value = data.fileHashName
      IEC61850Model.value = data.IEC61850Model
      await displayIEC61850Model(data.IEC61850Model)
    } else {
      const errData = await response.json()
      console.error('Failed to upload file:', errData)
    }
  } catch (error) {
    console.error('Error uploading file:', error)
  }
}

async function displayIEC61850Model(model = '') {
}

async function createStation() {
  try {
    const formData = new FormData()
    formData.append('name', stationName.value)
    formData.append('host', host.value)
    formData.append('port', port.value.toString())
    formData.append('icdFile', fileHashName.value)
    const response = await fetch('/api/stations', {
      method: 'POST',
      body: formData
    })
    if (response.ok) {
      const resp = await response.json()
      const stationId = resp.id
      router.push(`/stations/${stationId}`)
    } else {
      const data = await response.json()
      console.error('Failed to create station:', data)
    }
  } catch (error) {
    console.error('Error creating station:', error)
  }
}
</script>
