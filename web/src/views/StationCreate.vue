<template>
  <div class="container mx-auto py-8">
    <h1 class="text-3xl font-bold mb-6">Create Station</h1>
    <form @submit.prevent="createStation" class="max-w-md">
      <div v-show="step === 1">
        <h2 class="text-2xl font-bold mb-4">Step 1: Basic Information</h2>
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
      </div>

      <div v-show="step === 2">
        <h2 class="text-2xl font-bold mb-4">Step 2: Upload ICD File</h2>
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
      </div>

      <div v-show="step === 3">
        <h2 class="text-2xl font-bold mb-4">Step 3: ICD Analysis</h2>
        <div class="mb-4">
          <label class="block text-gray-700 font-bold mb-2">Analysis Progress</label>
          <div class="bg-gray-200 rounded-full h-4">
            <div
              class="bg-blue-500 h-4 rounded-full transition-all duration-500 ease-in-out"
              :style="{ width: `${progress}%` }"
            ></div>
          </div>
          <div class="text-right mt-2">{{ progress }}%</div>
        </div>
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
      </div>

      <div class="flex justify-between mt-8">
        <button
          type="button"
          @click="prevStep"
          class="bg-gray-300 hover:bg-gray-400 text-gray-800 font-bold py-2 px-4 rounded"
          :disabled="step === 1"
        >
          Previous
        </button>
        <button
          type="button"
          @click="nextStep"
          class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded"
          :disabled="step === 3"
        >
          Next
        </button>
        <button
          type="submit"
          class="bg-green-500 hover:bg-green-600 text-white font-bold py-2 px-4 rounded"
          :disabled="step !== 3"
        >
          Create
        </button>
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
const icdData = ref(null)
const icdStats = ref(null)

function handleFileUpload(event: Event) {
  const file = (event.target as HTMLInputElement).files?.[0] || null
  icdFile.value = file
}

function handleFileDrop(event: DragEvent) {
  const file = event.dataTransfer?.files[0] || null
  icdFile.value = file
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

function prevStep() {
  if (step.value > 1) {
    step.value--
  }
}

async function nextStep() {
  if (step.value === 2) {
    await analyzeICD()
  }
  if (step.value < 3) {
    step.value++
  }
}

async function analyzeICD() {
  // 模拟 ICD 解析进度
  progress.value = 0
  const interval = setInterval(() => {
    progress.value += 10
    if (progress.value >= 100) {
      clearInterval(interval)
      // 模拟解析完成后的数据和统计信息
      icdData.value = {
        /* ... */
      }
      icdStats.value = {
        /* ... */
      }
    }
  }, 500)
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
    const response = await fetch('/api/stations', { method: 'POST', body: formData })
    if (response.ok) {
      const data = await response.json()
      console.log('Station created:', data)
      router.push('/')
    } else {
      const errData = await response.json()
      console.error('Failed to create station:', errData)
    }
  } catch (error) {
    console.error('Error creating station:', error)
  }
}
</script>
