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
          <th class="py-3 px-6 text-left">Host</th>
          <th class="py-3 px-6 text-left">Port</th>
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
          <td class="py-4 px-6">{{ station.host }}</td>
          <td class="py-4 px-6">{{ station.port }}</td>
          <td class="py-4 px-6 text-center flex justify-center items-center space-x-2">
            <button
              @click="editConfig(station.id)"
              class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-1 px-3 rounded"
            >
              Edit
            </button>
            <button
              @click="confirmDelete(station.id)"
              class="bg-red-500 hover:bg-red-700 text-white font-bold py-1 px-3 rounded focus:outline-none focus:shadow-outline"
              type="button"
            >
              Delete
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
import { userGlobalStore } from '@/stores/store'
import { reactive } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const store = userGlobalStore()
const stations = reactive(store.stations)

function createStation() {
  router.push(`/station/create`)
}

function editConfig(stationID: number) {
  router.push(`/setting/${stationID}`)
}

const confirmDelete = (stationID: number) => {
  if (confirm('Are you sure you want to delete this station?')) {
    console.log('Delete station with ID:', stationID)
  }
}
</script>
