<template>
  <div class="container mx-auto p-6">
    <h1 class="text-3xl font-bold mb-6">{{ $t('userSettings') }}</h1>

    <div class="grid grid-cols-2 gap-8">
      <div>
        <section class="bg-white shadow rounded-lg p-6">
          <h2 class="text-2xl font-bold mb-6">{{ $t('themeAndLanguage') }}</h2>
          <div class="mb-6">
            <label class="block text-gray-700 font-bold mb-2">{{ $t('language') }}</label>
            <div class="relative">
              <select
                v-model="language"
                class="block appearance-none w-full bg-white border border-gray-400 hover:border-gray-500 px-4 py-2 pr-8 rounded shadow leading-tight focus:outline-none focus:shadow-outline"
              >
                <option value="en-US">{{ $t('english') }}</option>
                <option value="zh-CN">{{ $t('chinese') }}</option>
              </select>
              <div
                class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700"
              >
                <svg
                  class="fill-current h-4 w-4"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20"
                >
                  <path
                    d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z"
                  />
                </svg>
              </div>
            </div>
          </div>
          <div class="mb-6">
            <label class="block text-gray-700 font-bold mb-2">{{ $t('theme') }}</label>
            <div class="relative">
              <select
                v-model="theme"
                class="block appearance-none w-full bg-white border border-gray-400 hover:border-gray-500 px-4 py-2 pr-8 rounded shadow leading-tight focus:outline-none focus:shadow-outline"
              >
                <option value="light">{{ $t('light') }}</option>
                <option value="dark">{{ $t('dark') }}</option>
                <option value="auto">{{ $t('auto') }}</option>
              </select>
              <div
                class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700"
              >
                <svg
                  class="fill-current h-4 w-4"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20"
                >
                  <path
                    d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z"
                  />
                </svg>
              </div>
            </div>
          </div>
          <button
            @click="editLangAndTheme"
            class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 focus:outline-none focus:shadow-outline"
          >
            {{ $t('save') }}
          </button>
        </section>
      </div>

      <div>
        <section class="bg-white shadow rounded-lg p-6">
          <div class="flex justify-between mb-4">
            <h2 class="text-2xl font-bold">{{ $t('systemLogs') }}</h2>
            <div class="flex">
              <button
                @click="refreshLogs"
                class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 focus:outline-none focus:shadow-outline"
              >
                {{ $t('refresh') }}
              </button>
              <button
                @click="exportLogs"
                class="ml-2 px-4 py-2 bg-green-500 text-white rounded-md hover:bg-green-600 focus:outline-none focus:shadow-outline"
              >
                {{ $t('export') }}
              </button>
            </div>
          </div>
          <div class="flex items-center mb-4">
            <label class="block text-gray-700 font-bold mr-4">{{ $t('logLevel') }}</label>
            <div class="relative">
              <select
                v-model="logLevel"
                class="block appearance-none w-full bg-white border border-gray-400 hover:border-gray-500 px-4 py-2 pr-8 rounded shadow leading-tight focus:outline-none focus:shadow-outline"
              >
                <option value="debug">{{ $t('debug') }}</option>
                <option value="info">{{ $t('info') }}</option>
                <option value="warning">{{ $t('warning') }}</option>
                <option value="error">{{ $t('error') }}</option>
              </select>
              <div
                class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700"
              >
                <svg
                  class="fill-current h-4 w-4"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20"
                >
                  <path
                    d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z"
                  />
                </svg>
              </div>
            </div>
          </div>
          <div class="flex items-center mb-4">
            <label class="block text-gray-700 font-bold mr-4">{{ $t('timeRange') }}</label>
            <input
              type="date"
              v-model="logStartDate"
              class="border border-gray-400 hover:border-gray-500 px-2 py-1 rounded shadow focus:outline-none focus:shadow-outline"
            />
            <span class="mx-2">{{ $t('to') }}</span>
            <input
              type="date"
              v-model="logEndDate"
              class="border border-gray-400 hover:border-gray-500 px-2 py-1 rounded shadow focus:outline-none focus:shadow-outline"
            />
          </div>

          <table class="table-auto w-full">
            <thead>
              <tr class="bg-gray-200 text-gray-600 uppercase text-sm leading-normal">
                <th class="py-3 px-6 text-left">{{ $t('timeStamp') }}</th>
                <th class="py-3 px-6 text-left">{{ $t('level') }}</th>
                <th class="py-3 px-6 text-left">{{ $t('message') }}</th>
              </tr>
            </thead>
            <tbody class="text-gray-600 text-sm">
              <tr
                v-for="log in filteredLogs"
                :key="log.id"
                class="border-b border-gray-200 hover:bg-gray-100"
              >
                <td class="py-3 px-6">{{ log.timestamp }}</td>
                <td class="py-3 px-6">{{ log.level }}</td>
                <td class="py-3 px-6">{{ log.message }}</td>
              </tr>
            </tbody>
          </table>
        </section>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useAuthStore } from '@/store'

interface Log {
  id: number
  timestamp: string
  level: 'info' | 'warning' | 'error'
  message: string
}

const authStore = useAuthStore()
const language = ref<string>('en-US')
const theme = ref<string>('light')
const logLevel = ref<string>('info')
const logStartDate = ref<string>('')
const logEndDate = ref<string>('')
const logs = ref<Log[]>([
  { id: 1, timestamp: '2023-6-7 15:00:02', level: 'info', message: 'System started' },
  { id: 2, timestamp: '2023-6-7 15:00:22', level: 'info', message: 'admin logged in' },
  { id: 3, timestamp: '2023-6-7 15:00:31', level: 'info', message: 'test_station_1 is online' },
  { id: 4, timestamp: '2023-6-7 15:00:56', level: 'info', message: "test_station_2 is online"},
  { id: 5, timestamp: '2023-6-7 15:01:09', level: 'info', message: 'test_station_3 is created'},
])

const filteredLogs = computed(() => {
  return logs.value.filter((log) => {
    const logDate = new Date(log.timestamp)
    const startDate = logStartDate.value ? new Date(logStartDate.value) : null
    const endDate = logEndDate.value ? new Date(logEndDate.value) : null
    return (
      log.level === logLevel.value &&
      (!startDate || logDate >= startDate) &&
      (!endDate || logDate <= endDate)
    )
  })
})

function refreshLogs(): void {
  console.log('Refreshing logs')
}

function exportLogs(): void {
  console.log('Exporting logs')
}

async function editLangAndTheme() {
  try {
    await authStore.changeLang(language.value)
  } catch (error) {
    console.error(error)
  }

  try {
    await authStore.changeTheme(theme.value)
  } catch (error) {
    console.error(error)
  }
}
</script>
