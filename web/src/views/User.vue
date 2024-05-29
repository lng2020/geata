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
          <div class="flex justify-between">
            <h2 class="text-2xl font-bold mb-4">{{ $t('MQTTConnectionParameters') }}</h2>
            <button
              @click="testMqttConnection"
              class="mb-4 px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 focus:outline-none focus:shadow-outline"
            >
              {{ $t('testConnection') }}
            </button>
          </div>

          <div>
            <div class="grid grid-cols-2 gap-4">
              <div class="mb-4">
                <label class="block text-gray-700 font-bold mb-2">{{ $t('address') }}</label>
                <input
                  type="text"
                  v-model="mqttServerAddress"
                  class="w-full border border-gray-400 hover:border-gray-500 px-2 py-1 rounded shadow focus:outline-none focus:shadow-outline"
                />
              </div>
              <div>
                <label class="block text-gray-700 font-bold mb-2">{{ $t('port') }}</label>
                <input
                  type="number"
                  v-model="mqttServerPort"
                  class="w-full border border-gray-400 hover:border-gray-500 px-2 py-1 rounded shadow focus:outline-none focus:shadow-outline"
                />
              </div>
              <div>
                <label class="block text-gray-700 font-bold mb-2">{{ $t('username') }}</label>
                <input
                  type="text"
                  v-model="mqttUsername"
                  class="w-full border border-gray-400 hover:border-gray-500 px-2 py-1 rounded shadow focus:outline-none focus:shadow-outline"
                />
              </div>
              <div>
                <label class="block text-gray-700 font-bold mb-2">{{ $t('password') }}</label>
                <input
                  type="password"
                  v-model="mqttPassword"
                  class="w-full border border-gray-400 hover:border-gray-500 px-2 py-1 rounded shadow focus:outline-none focus:shadow-outline"
                />
              </div>
            </div>
          </div>
          <button
            class="mt-6 px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 focus:outline-none focus:shadow-outline"
          >
            {{ $t('save') }}
          </button>
        </section>
      </div>

      <div>
        <section class="bg-white shadow rounded-lg p-6">
          <div class="flex justify-between">
            <h2 class="text-2xl font-bold mb-4">{{ $t('alarmRules') }}</h2>
            <button
              @click="showAddRuleDialog = true"
              class="mb-4 px-4 py-2 bg-green-500 text-white rounded-md hover:bg-green-600 focus:outline-none focus:shadow-outline"
            >
              {{ $t('addRule') }}
            </button>
          </div>
          <ul class="space-y-2">
            <li
              v-for="(rule, index) in alarmRules"
              :key="index"
              class="flex items-center justify-between bg-gray-100 p-2 rounded-md"
            >
              <div class="flex items-center">
                <span>{{ rule.condition }}</span>
                <span class="mx-2">|</span>
                <span>{{ rule.notification }}</span>
              </div>
              <div>
                <button
                  @click="editRule(index)"
                  class="px-4 py-1 bg-blue-500 text-white rounded-md hover:bg-blue-600 focus:outline-none focus:shadow-outline"
                >
                  {{ $t('edit') }}
                </button>
                <button
                  @click="deleteRule(index)"
                  class="ml-2 px-4 py-1 bg-red-500 text-white rounded-md hover:bg-red-600 focus:outline-none focus:shadow-outline"
                >
                  {{ $t('delete') }}
                </button>
              </div>
            </li>
          </ul>
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

    <div
      v-if="showAddRuleDialog"
      class="fixed inset-0 flex items-center justify-center bg-gray-900 bg-opacity-50"
    >
      <div class="bg-white p-6 rounded-lg shadow-xl">
        <h3 class="text-xl font-bold mb-4">{{ $t('addAlarmRule') }}</h3>
        <div class="mb-4">
          <label class="block text-gray-700 font-bold mb-2">{{ $t('condition') }}</label>
          <input
            type="text"
            v-model="newRule.condition"
            class="w-full border border-gray-400 hover:border-gray-500 px-2 py-1 rounded shadow focus:outline-none focus:shadow-outline"
          />
        </div>
        <div class="mb-4">
          <label class="block text-gray-700 font-bold mb-2">{{ $t('notification') }}</label>
          <input
            type="text"
            v-model="newRule.notification"
            class="w-full border border-gray-400 hover:border-gray-500 px-2 py-1 rounded shadow focus:outline-none focus:shadow-outline"
          />
        </div>
        <div class="flex justify-end">
          <button
            @click="addRule"
            class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 focus:outline-none focus:shadow-outline"
          >
            {{ $t('add') }}
          </button>
          <button
            @click="showAddRuleDialog = false"
            class="ml-2 px-4 py-2 bg-gray-500 text-white rounded-md hover:bg-gray-600 focus:outline-none focus:shadow-outline"
          >
            {{ $t('cancel') }}
          </button>
        </div>
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

interface Rule {
  condition: string
  notification: string
}

const authStore = useAuthStore()
const language = ref<string>('en-US')
const theme = ref<string>('light')
const mqttServerAddress = ref<string>('mqtt://localhost')
const mqttServerPort = ref<number>(1883)
const mqttUsername = ref<string>('')
const mqttPassword = ref<string>('')
const alarmRules = ref<Rule[]>([
  { condition: 'Temperature > 30', notification: 'Send email' },
  { condition: 'Humidity < 40%', notification: 'Send SMS' }
])
const showAddRuleDialog = ref<boolean>(false)
const newRule = ref<Rule>({ condition: '', notification: '' })
const logLevel = ref<string>('info')
const logStartDate = ref<string>('')
const logEndDate = ref<string>('')
const logs = ref<Log[]>([
  { id: 1, timestamp: '2023-05-15 10:00:00', level: 'info', message: 'System started' },
  { id: 2, timestamp: '2023-05-15 11:30:00', level: 'warning', message: 'Low disk space' },
  { id: 3, timestamp: '2023-05-15 13:45:00', level: 'error', message: 'Connection lost' }
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

function testMqttConnection(): void {
  console.log('Testing MQTT connection')
}

function editRule(index: number): void {
  console.log('Editing rule:', index)
}

function deleteRule(index: number): void {
  alarmRules.value.splice(index, 1)
}

function addRule(): void {
  alarmRules.value.push({ ...newRule.value })
  newRule.value.condition = ''
  newRule.value.notification = ''
  showAddRuleDialog.value = false
}

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
