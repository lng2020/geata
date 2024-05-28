<template>
  <div class="container mx-auto px-4 py-8">
    <h1 class="text-3xl font-bold mb-6">{{ $t('settings') }}</h1>
    <div v-if="editedStation" class="flex">
      <div class="w-1/2 pr-4">
        <div class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
          <h2 class="text-xl font-bold mb-4">{{ $t('basicInformation') }}</h2>
          <div class="mb-4">
            <label class="block text-gray-700 text-sm font-bold mb-2" for="name">{{
              $t('name')
            }}</label>
            <input
              v-model="editedStation.name"
              class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              id="name"
              type="text"
              :placeholder="$t('powerStationName')"
            />
          </div>
          <div class="mb-4">
            <label class="block text-gray-700 text-sm font-bold mb-2" for="address">{{
              $t('address')
            }}</label>
            <input
              v-model="editedStation.host"
              class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              id="address"
              type="text"
              :placeholder="$t('powerStationAddress')"
            />
          </div>
          <div class="mb-4">
            <label class="block text-gray-700 text-sm font-bold mb-2" for="port">{{
              $t('port')
            }}</label>
            <input
              v-model.number="editedStation.port"
              class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              id="port"
              type="number"
              :placeholder="$t('port')"
            />
          </div>
          <div class="flex items-center justify-between">
            <button
              @click="saveChanges"
              class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
              type="button"
            >
              {{ $t('save') }}
            </button>
            <button
              @click="cancel"
              class="bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
              type="button"
            >
              {{ $t('cancel') }}
            </button>
          </div>
        </div>
      </div>
      <div class="w-1/2 pl-4">
        <div class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
          <h2 class="text-xl font-bold mb-4">{{ $t('dataSourceConfiguration') }}</h2>
          <ul class="space-y-4">
            <li
              v-for="(dataSource, index) in dataSources"
              :key="index"
              class="flex items-center justify-between"
            >
              <div class="flex items-center">
                <span class="text-gray-700 font-semibold">{{ dataSource.name }}</span>
                <span class="ml-2">
                  <div v-if="dataSource.status === 'connected'" class="flex items-center">
                    <svg
                      class="w-5 h-5 text-green-500 mr-1"
                      fill="none"
                      stroke="currentColor"
                      viewBox="0 0 24 24"
                      xmlns="http://www.w3.org/2000/svg"
                    >
                      <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M5 13l4 4L19 7"
                      ></path>
                    </svg>
                    <span class="text-green-500 font-semibold">{{ $t('connected') }}</span>
                  </div>
                  <div v-else-if="dataSource.status === 'disconnected'" class="flex items-center">
                    <svg
                      class="w-5 h-5 text-red-500 mr-1"
                      fill="none"
                      stroke="currentColor"
                      viewBox="0 0 24 24"
                      xmlns="http://www.w3.org/2000/svg"
                    >
                      <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
                      ></path>
                    </svg>
                    <span class="text-red-500 font-semibold">{{ $t('disconnected') }}</span>
                  </div>
                </span>
              </div>
              <button
                @click="testDataSource(dataSource)"
                class="bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-4 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                type="button"
              >
                {{ $t('testConnection') }}
              </button>
            </li>
          </ul>
        </div>

        <div class="bg-white shadow-md rounded px-8 pt-6 pb-8">
          <h2 class="text-xl font-bold mb-4">{{ $t('icdFileUpload') }}</h2>
          <div class="mb-4">
            <label for="icdFile" class="block text-gray-700 font-semibold mb-2">
              {{ $t('selectIcdFile') }}
            </label>
            <div class="relative">
              <input
                @change="handleFileUpload"
                class="appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:ring-2 focus:ring-blue-500"
                id="icdFile"
                type="file"
                accept=".icd"
              />
            </div>
          </div>
          <div v-if="uploadProgress > 0" class="mb-4">
            <label class="block text-gray-700 font-semibold mb-2">{{ $t('uploadProgress') }}</label>
            <div class="relative pt-1">
              <div class="overflow-hidden h-2 text-xs flex rounded bg-blue-200">
                <div
                  :style="{ width: uploadProgress + '%' }"
                  class="shadow-none flex flex-col text-center whitespace-nowrap text-white justify-center bg-blue-500 transition-all duration-500"
                ></div>
              </div>
            </div>
          </div>
          <div v-if="parsedIcdData" class="mt-4">
            <label class="block text-gray-700 font-semibold mb-2">{{ $t('parsedIcdData') }}</label>
            <pre class="bg-gray-100 rounded p-4 text-sm overflow-auto">{{ parsedIcdData }}</pre>
          </div>
        </div>
      </div>
    </div>
    <div v-else>
      <p>{{ $t('loadingPowerStation') }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { userGlobalStore } from '@/store'
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import type { Station, DataSource } from '@/types/types'

const store = userGlobalStore()
const route = useRoute()
const router = useRouter()
const ID = ref<number | null>(null)
let station = ref<Station | undefined>()
let editedStation = ref<Station | undefined>()
const dataSources = ref<DataSource[]>([
  { name: 'IEC61850', status: 'disconnected' },
  { name: 'Modbus', status: 'disconnected' },
  { name: 'MQTT', status: 'disconnected' }
])
const uploadProgress = ref(0)
const parsedIcdData = ref<string | null>(null)

onMounted(() => {
  const id = route.params.id
  if (typeof id === 'string') {
    ID.value = parseInt(id, 10)
    station.value = store.getStationByID(ID.value)
    editedStation.value = station.value
  } else {
    console.error('Invalid power station ID')
    router.push('/error')
  }
})

const saveChanges = () => {
  if (editedStation.value) {
    station.value = editedStation.value
    // TODO: update store method
  }
}

const cancel = () => {
  if (station.value) {
    editedStation.value = station.value
  }
}

const testDataSource = (dataSource: DataSource) => {
  // TODO: Implement data source connectivity test
  // Update dataSource.status based on the test result
}
const handleFileUpload = (event: Event) => {
  const file = (event.target as HTMLInputElement).files?.[0]
  if (file) {
    const formData = new FormData()
    formData.append('icdFile', file)

    // TODO: Implement file upload logic
    // You can use axios or fetch to send the formData to the backend API
    // Update uploadProgress during the upload process
    // Once the upload is complete, trigger the parsing of the ICD file
    // Update parsedIcdData with the parsed result
  }
}
</script>
