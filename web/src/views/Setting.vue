<template>
  <div class="container mx-auto p-4">
    <h1 class="text-2xl font-bold mb-4">Settings</h1>
    <div class="flex">
      <div class="w-1/2 pr-4">
        <div class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
          <h2 class="text-xl font-bold mb-4">Configuration</h2>
          <div class="mb-4">
            <label class="block text-gray-700 text-sm font-bold mb-2" for="name"> Name </label>
            <input
              v-model="station.name"
              class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              id="name"
              type="text"
              placeholder="Station Name"
            />
          </div>
          <div class="mb-4">
            <label class="block text-gray-700 text-sm font-bold mb-2" for="host"> Host </label>
            <input
              v-model="station.host"
              class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              id="host"
              type="text"
              placeholder="Host"
            />
          </div>
          <div class="mb-4">
            <label class="block text-gray-700 text-sm font-bold mb-2" for="port"> Port </label>
            <input
              v-model.number="station.port"
              class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              id="port"
              type="number"
              placeholder="Port"
            />
          </div>
          <div class="flex items-center justify-between">
            <button
              @click="deleteStation"
              class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
              type="button"
            >
              Delete Station
            </button>
          </div>
        </div>
      </div>
      <div class="w-1/2 pl-4">
        <div class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
          <h2 class="text-xl font-bold mb-4">Status</h2>
          <div class="mb-4">
            <div class="flex items-center">
              <div
                class="w-4 h-4 rounded-full"
                :class="{ 'bg-green-500': station.isOnline, 'bg-red-500': !station.isOnline }"
              ></div>
              <span class="ml-2 text-gray-700">{{ station.isOnline ? 'Online' : 'Offline' }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

interface Station {
  id: number
  name: string
  host: string
  port: number
  isOnline: boolean
  lastOnlineTime: string
  createdAt: string
  updatedAt: string
}

const station = ref<Station>({
  id: 0,
  name: '',
  host: '',
  port: 0,
  isOnline: false,
  lastOnlineTime: '',
  createdAt: '',
  updatedAt: ''
})

const deleteStation = () => {
  if (confirm('Are you sure you want to delete this station?')) {
    // Logic to delete the station
    console.log('Deleting station:', station.value)
  }
}
</script>
