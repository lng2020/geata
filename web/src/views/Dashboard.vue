<template>
  <div class="container mx-auto py-8">
    <div class="flex items-center justify-between mb-6">
      <h1 class="text-3xl font-bold">{{ $t('dashboard') }}</h1>
      <div class="flex items-center">
        <input
          type="text"
          v-model="searchQuery"
          placeholder="search stations by name..."
          class="w-64 px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <button
          @click="searchStations"
          class="ml-4 px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500"
        >
          Search Stations
        </button>
        <button
          @click="createStation"
          class="ml-4 px-4 py-2 bg-green-500 text-white rounded-md hover:bg-green-600 focus:outline-none focus:ring-2 focus:ring-green-500"
        >
          Create Station
        </button>
      </div>
    </div>

    <div class="grid grid-cols-3 gap-4 mb-6">
      <div class="col-span-2 bg-white shadow-md rounded-lg p-6">
        <table class="w-full">
          <thead>
            <tr class="bg-gray-100 text-gray-600 uppercase text-sm leading-normal">
              <th class="py-3 px-6 text-left">{{ $t('name') }}</th>
              <th class="py-3 px-6 text-left">{{ $t('isOnline') }}</th>
              <th class="py-3 px-6 text-left">{{ $t('lastOnlineTime') }}</th>
              <th class="py-3 px-6 text-left">ADDRESS</th>
            </tr>
          </thead>
          <tbody class="text-gray-600 text-sm">
            <tr
              v-for="station in filteredStations"
              :key="station.id"
              class="border-b border-gray-200 hover:bg-gray-100"
            >
              <td class="py-4 px-6">{{ station.name }}</td>
              <td class="py-4 px-6">
                <span
                  class="inline-block rounded-full px-3 py-1 text-sm font-semibold"
                  :class="station.isOnline ? 'bg-green-500 text-white' : 'bg-red-500 text-white'"
                >
                  {{ station.isOnline ? 'online' : 'offline' }}
                </span>
              </td>
              <td class="py-4 px-6">{{ station.lastOnlineTime }}</td>
              <td class="py-4 px-6">{{ station.host }}</td>
            </tr>
          </tbody>
        </table>
      </div>

      <div class="bg-white shadow-md rounded-lg p-4">
        <div class="flex flex-col w-full h-full">
          <h2 class="text-xl font-semibold mb-4 text-left">ONLINE STATUS</h2>
          <div class="flex justify-center items-center w-full h-full">
            <Pie :data="stationStatusData" :options="chartOptions" />
          </div>
        </div>
      </div>
    </div>
    <div class="bg-white shadow-md rounded-lg p-4">
      <div class="flex flex-col w-full h-full">
        <h2 class="text-xl font-semibold mb-4 text-left">ONLINE STATUS</h2>
        <div class="flex justify-center items-center w-full h-full">
          <Line :data="onlineStationsData" :options="chartOptions" />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { userGlobalStore } from '@/store'
import { computed, reactive, ref } from 'vue'
import { Pie, Line } from 'vue-chartjs'
import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  ArcElement,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement
} from 'chart.js'
import router from '@/router'
ChartJS.register(
  Title,
  Tooltip,
  Legend,
  ArcElement,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement
)
const store = userGlobalStore()
const stations = reactive(store.stations)
const searchQuery = ref('')
const filteredStations = computed(() => {
  if (searchQuery.value.trim() === '') {
    return stations
  } else {
    const query = searchQuery.value.toLowerCase()
    return stations.filter(
      (station) =>
        station.name.toLowerCase().includes(query) || station.host.toLowerCase().includes(query)
    )
  }
})
const stationStatusData = computed(() => ({
  labels: ['online', 'offline'],
  datasets: [
    {
      data: [stations.filter((s) => s.isOnline).length, stations.filter((s) => !s.isOnline).length],
      backgroundColor: ['rgb(75, 192, 192)', 'rgb(255, 99, 132)']
    }
  ]
}))
const onlineStationsData = computed(() => ({
  labels: ['Day 1', 'Day 2', 'Day 3', 'Day 4', 'Day 5', 'Day 6', 'Day 7'],
  datasets: [
    {
      label: 'online station number',
      data: [0, 5, 10, 15, 12, 18, 20, 22, 25],
      fill: false,
      borderColor: 'rgb(75, 192, 192)',
      tension: 0.1
    }
  ]
}))
const chartOptions = { responsive: true, maintainAspectRatio: false }
function searchStations() {
  console.log('Search stations with query:', searchQuery.value)
}
const createStation = () => {
  router.push('/station/create')
}
</script>
