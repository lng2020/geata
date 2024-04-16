<template>
  <div class="container mx-auto py-8">
    <h1 class="text-3xl font-bold mb-6">Dashboard</h1>

    <div class="grid grid-cols-2 gap-4 mb-6">
      <!-- Station Status Pie Chart -->
      <div class="bg-white shadow-md rounded-lg p-4">
        <div class="flex flex-col w-full h-ful">
          <h2 class="text-xl font-semibold mb-4 text-left">Station Status</h2>
          <div class="flex justify-center items-center w-full h-full">
            <Pie :data="stationStatusData" :options="chartOptions" />
          </div>
        </div>
      </div>

      <!-- Online Stations Line Chart -->
      <div class="bg-white shadow-md rounded-lg p-4">
        <div class="flex flex-col w-full h-full">
          <h2 class="text-xl font-semibold mb-4 text-left">Online Stations</h2>
          <div class="flex justify-center items-center w-full h-full">
            <Line :data="onlineStationsData" :options="chartOptions" />
          </div>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-3 gap-4">
      <!-- Station Table -->
      <div class="col-span-2 bg-white shadow-md rounded-lg p-6">
        <div class="flex justify-between items-center mb-6">
          <h2 class="text-2xl font-semibold text-gray-800">Stations</h2>
          <button
            @click="createStation"
            class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded"
          >
            Create Station
          </button>
        </div>
        <table class="w-full">
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

      <!-- System Audit Log -->
      <div class="bg-white shadow-md rounded-lg p-6">
        <h2 class="text-2xl font-semibold mb-6 text-gray-800">System Audit Log</h2>
        <ul class="space-y-6">
          <li
            v-for="log in auditLogs"
            :key="log.id"
            class="border-b border-gray-200 pb-6 last:border-none last:pb-0"
          >
            <div class="flex items-start">
              <div class="flex-shrink-0">
                <div class="w-10 h-10 rounded-full bg-blue-500 flex items-center justify-center">
                  <svg
                    class="w-6 h-6 text-white"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
                    ></path>
                  </svg>
                </div>
              </div>
              <div class="ml-4">
                <p class="text-base font-medium text-gray-900">{{ log.message }}</p>
                <small class="text-sm text-gray-500">{{ log.createAt }}</small>
              </div>
            </div>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { userGlobalStore } from '@/stores/store'
import { computed, reactive } from 'vue'
import { useRouter } from 'vue-router'
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

const router = useRouter()
const store = userGlobalStore()
const stations = reactive(store.stations)
const auditLogs = reactive(store.auditLogs)

const stationStatusData = computed(() => ({
  labels: ['Online', 'Offline'],
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
      label: 'Online Stations',
      data: [10, 15, 12, 18, 20, 22, 25],
      fill: false,
      borderColor: 'rgb(75, 192, 192)',
      tension: 0.1
    }
  ]
}))

const chartOptions = { responsive: true, maintainAspectRatio: false }

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
