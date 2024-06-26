<template>
  <div class="container mx-auto px-2 py-2 flex justify-center">
    <form @submit.prevent="createStation" class="max-w-md w-full">
      <h1 class="text-3xl font-bold mb-6 text-center">{{ $t('createStation')}}</h1>

      <div class="mb-8">
        <div v-show="step === 1" class="mb-8">
          <div class="mb-4">
            <label for="name" class="block text-gray-700 font-bold mb-2">{{ $t('stationName') }}</label>
            <input
              type="text"
              id="name"
              v-model="stationName"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              required
            />
          </div>
          <div class="mb-4">
            <label for="host" class="block text-gray-700 font-bold mb-2">{{ $t('host')}}</label>
            <input
              type="text"
              id="host"
              v-model="host"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              required
            />
          </div>
          <div class="mb-4">
            <label for="port" class="block text-gray-700 font-bold mb-2">{{ $t('port')}}</label>
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
              @click="validateStationConfig"
            >
              {{  $t('confirm') }}
            </button>
          </div>
        </div>

        <div v-show="step === 2" class="mb-8">
          <div class="mb-4">
            <label for="icdFile" class="block text-gray-700 font-bold mb-2">{{ $t('icdFile')}}</label>
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
                <span v-if="!icdFile">{{ $t("clickOrDropFile") }}</span>
                <span v-else>{{ icdFile.name }} ({{ formatFileSize(icdFile.size) }})</span>
              </label>
            </div>
          </div>
          <div class="flex justify-end">
            <button
              type="button"
              class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded"
              @click="validateICDFile"
            >
              Confirm
            </button>
          </div>
        </div>

        <div v-show="step === 3" class="mb-8">
          <div v-if="model" class="mb-4">
            <h2 class="text-xl font-bold mb-2">{{ $t('iec61850Model')}}</h2>
            <div
              class="border border-gray-300 rounded-md p-4 bg-white shadow-sm overflow-y-auto max-h-[500px]"
            >
              <table class="w-full border-collapse">
                <thead>
                  <tr class="bg-gray-200 text-gray-700">
                    <th class="px-4 py-2 font-medium text-left">LD</th>
                    <th class="px-4 py-2 font-medium text-left">LN</th>
                  </tr>
                </thead>
                <tbody>
                  <template v-for="(device, deviceIndex) in model.logicalDevice" :key="deviceIndex">
                    <tr
                      v-for="(node, nodeIndex) in device.logicalNode"
                      :key="nodeIndex"
                      class="border-t"
                    >
                      <td
                        v-if="nodeIndex === 0"
                        :rowspan="device.logicalNode.length"
                        class="px-4 py-2 align-top"
                      >
                        {{ device.name }}
                      </td>
                      <td class="px-4 py-2">{{ node.name }}</td>
                    </tr>
                  </template>
                </tbody>
              </table>
            </div>
          </div>
          <div class="flex justify-end">
            <button
              type="submit"
              class="bg-green-500 hover:bg-green-600 text-white font-bold py-2 px-4 rounded"
            >
            {{ $t('create') }}
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
import { type IEC61850Model } from '@/types/types'
import { useGlobalStore } from '@/store'

const router = useRouter()
const step = ref(1)
const stationName = ref('')
const host = ref('')
const port = ref(0)
const icdFile = ref<File | null>(null)
const model = ref<IEC61850Model | null>(null)

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

const nextStep = () => {
  if (step.value < 3) {
    step.value++
  }
}

const validateStationConfig = () => {
  if (stationName.value && host.value && port.value > 0) {
    nextStep()
  } else {
    alert('Please fill in all fields')
  }
}

const validateICDFile = () => {
  if (icdFile.value) {
    nextStep()
  } else {
    alert('Please upload an ICD file')
  }
}

async function uploadICDFile(file: File) {
  const formData = new FormData()
  formData.append('file', file)

  try {
    const response = await fetch('/api/iec61850/upload', {
      method: 'POST',
      body: formData
    })
    if (response.ok) {
      const data = await response.json()
      model.value = data
      console.log('Uploaded ICD file:', data)
    } else {
      const errData = await response.json()
      console.error('Failed to upload file:', errData)
    }
  } catch (error) {
    console.error('Error uploading file:', error)
  }
}

async function createStation() {
  try {
    const response = await fetch('/api/station', {
      method: 'POST',
      body: JSON.stringify({
        name: stationName.value,
        host: host.value,
        port: port.value,
        modelId: model.value?.id
      })
    })
    if (response.ok) {
      const resp = await response.json()
      const stationId = resp.id
      useGlobalStore().fetchStations()
      router.push(`/station/${stationId}`)
    } else {
      const data = await response.json()
      console.error('Failed to create station:', data)
    }
  } catch (error) {
    console.error('Error creating station:', error)
  }
}
</script>
