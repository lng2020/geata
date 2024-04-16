<template>
  <div class="container mx-auto py-8">
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-3xl font-bold">Dashboard</h1>
      <button
        @click="createStation"
        class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded"
      >
        Create Station
      </button>
    </div>

    <table class="w-full bg-white shadow-md rounded-lg overflow-hidden">
      <thead>
        <tr class="bg-gray-100 text-gray-600 uppercase text-sm leading-normal">
          <th class="py-3 px-6 text-left">Name</th>
          <th class="py-3 px-6 text-left">Is Online</th>
          <th class="py-3 px-6 text-left">Last Online Time</th>
          <th class="py-3 px-6 text-center">Actions</th>
        </tr>
      </thead>
      <tbody class="text-gray-600 text-sm">
        <tr
          v-for="station in stations"
          :key="station.id"
          class="border-b border-gray-200 hover:bg-gray-100"
        >
          <td class="py-4 px-6">{{ station.name }}</td>
          <td class="py-4 px-6">
            <span
              class="inline-block rounded-full px-3 py-1 text-sm font-semibold"
              :class="station.isOnline ? 'bg-green-500 text-white' : 'bg-red-500 text-white'"
            >
              {{ station.isOnline ? 'Online' : 'Offline' }}
            </span>
          </td>
          <td class="py-4 px-6">{{ station.lastOnlineTime }}</td>
          <td class="py-4 px-6 text-center">
            <button
              @click="editConfig(station.id)"
              class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded"
            >
              Edit Config
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const stations = ref([
  { id: 1, name: 'Station 1', isOnline: true, lastOnlineTime: '2023-06-01 10:00:00' },
  { id: 2, name: 'Station 2', isOnline: false, lastOnlineTime: '2023-06-02 15:30:00' }
  // Add more stations as needed
])

function createStation() {
  router.push(`/station/create`)
}

function editConfig(stationId: number) {
  router.push(`/setting/${stationId}`)
}
</script>
