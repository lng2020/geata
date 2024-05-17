<template>
  <div class="container mx-auto p-6">
    <h1 class="text-3xl font-bold mb-6">System Settings</h1>

    <div class="grid grid-cols-2 gap-8">
      <section>
        <!-- Language -->
        <div class="mb-8">
          <h2 class="text-xl font-bold mb-2">Language</h2>
          <select v-model="language" class="border border-gray-300 px-2 py-1 rounded-md">
            <option value="en">English</option>
            <option value="zh">中文</option>
          </select>
        </div>

        <!-- Interface Theme -->
        <div class="mb-8">
          <h2 class="text-xl font-bold mb-2">Theme</h2>
          <select v-model="theme" class="border border-gray-300 px-2 py-1 rounded-md">
            <option value="light">Light</option>
            <option value="dark">Dark</option>
            <option value="auto">Auto</option>
          </select>
        </div>

        <!-- Data Retention Policy -->
        <div class="mb-8">
          <h2 class="text-xl font-bold mb-4">Data Retention Policy</h2>
          <div class="mb-2">
            <label class="mr-2">Real-time Data:</label>
            <input
              type="number"
              v-model="realTimeDataRetention"
              class="border border-gray-300 px-2 py-1 rounded-md"
              min="1"
              max="365"
            />
            <span class="ml-2">days</span>
          </div>
          <div>
            <label class="mr-2">Historical Data:</label>
            <input
              type="number"
              v-model="historicalDataRetention"
              class="border border-gray-300 px-2 py-1 rounded-md"
              min="1"
              max="365"
            />
            <span class="ml-2">days</span>
          </div>
        </div>

        <!-- MQTT Connection Parameters -->
        <div>
          <h2 class="text-xl font-bold mb-4">MQTT Connection Parameters</h2>
          <div class="mb-4">
            <label class="block mb-1">Server Address:</label>
            <input
              type="text"
              v-model="mqttServerAddress"
              class="w-full border border-gray-300 px-2 py-1 rounded-md"
            />
          </div>
          <div class="mb-4">
            <label class="block mb-1">Port:</label>
            <input
              type="number"
              v-model="mqttServerPort"
              class="w-full border border-gray-300 px-2 py-1 rounded-md"
            />
          </div>
          <div class="mb-4">
            <label class="block mb-1">Username:</label>
            <input
              type="text"
              v-model="mqttUsername"
              class="w-full border border-gray-300 px-2 py-1 rounded-md"
            />
          </div>
          <div class="mb-4">
            <label class="block mb-1">Password:</label>
            <input
              type="password"
              v-model="mqttPassword"
              class="w-full border border-gray-300 px-2 py-1 rounded-md"
            />
          </div>
          <button @click="testMqttConnection" class="px-4 py-2 bg-blue-500 text-white rounded-md">
            Test Connection
          </button>
        </div>
      </section>

      <section>
        <!-- Alarm Rules -->
        <div class="mb-8">
          <h2 class="text-xl font-bold mb-4">Alarm Rules</h2>
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
                  class="px-2 py-1 bg-blue-500 text-white rounded-md"
                >
                  Edit
                </button>
                <button
                  @click="deleteRule(index)"
                  class="ml-2 px-2 py-1 bg-red-500 text-white rounded-md"
                >
                  Delete
                </button>
              </div>
            </li>
          </ul>
          <button
            @click="showAddRuleDialog = true"
            class="mt-4 px-4 py-2 bg-green-500 text-white rounded-md"
          >
            Add Rule
          </button>
        </div>

        <!-- System Logs -->
        <div>
          <h2 class="text-xl font-bold mb-4">System Logs</h2>
          <div class="mb-4 flex items-center">
            <label class="mr-2 w-20">Log Level:</label>
            <select v-model="logLevel" class="border border-gray-300 px-2 py-1 rounded-md">
              <option value="debug">Debug</option>
              <option value="info">Info</option>
              <option value="warning">Warning</option>
              <option value="error">Error</option>
            </select>
          </div>
          <div class="mb-4 flex items-center">
            <label class="mr-2 w-20">Time Range:</label>
            <input
              type="date"
              v-model="logStartDate"
              class="border border-gray-300 px-2 py-1 rounded-md"
            />
            <span class="mx-2">to</span>
            <input
              type="date"
              v-model="logEndDate"
              class="border border-gray-300 px-2 py-1 rounded-md"
            />
          </div>
          <div class="flex">
            <button @click="refreshLogs" class="px-4 py-2 bg-blue-500 text-white rounded-md">
              Refresh
            </button>
            <button @click="exportLogs" class="ml-2 px-4 py-2 bg-green-500 text-white rounded-md">
              Export
            </button>
          </div>
          <table class="mt-4 w-full border-collapse">
            <thead>
              <tr class="bg-gray-200">
                <th class="px-4 py-2 border border-gray-300">Timestamp</th>
                <th class="px-4 py-2 border border-gray-300">Level</th>
                <th class="px-4 py-2 border border-gray-300">Message</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="log in filteredLogs" :key="log.id" class="hover:bg-gray-100">
                <td class="border border-gray-300 px-4 py-2">{{ log.timestamp }}</td>
                <td class="border border-gray-300 px-4 py-2">{{ log.level }}</td>
                <td class="border border-gray-300 px-4 py-2">{{ log.message }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </section>
    </div>

    <!-- Add Rule Dialog -->
    <div
      v-if="showAddRuleDialog"
      class="fixed inset-0 flex items-center justify-center bg-gray-900 bg-opacity-50"
    >
      <div class="bg-white p-6 rounded-md">
        <h3 class="text-xl font-bold mb-4">Add Alarm Rule</h3>
        <div class="mb-4">
          <label class="block mb-1">Condition:</label>
          <input
            type="text"
            v-model="newRule.condition"
            class="w-full border border-gray-300 px-2 py-1 rounded-md"
          />
        </div>
        <div class="mb-4">
          <label class="block mb-1">Notification:</label>
          <input
            type="text"
            v-model="newRule.notification"
            class="w-full border border-gray-300 px-2 py-1 rounded-md"
          />
        </div>
        <div class="flex justify-end">
          <button @click="addRule" class="px-4 py-2 bg-blue-500 text-white rounded-md">Add</button>
          <button
            @click="showAddRuleDialog = false"
            class="ml-2 px-4 py-2 bg-gray-500 text-white rounded-md"
          >
            Cancel
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

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

const language = ref<string>('en')
const theme = ref<string>('light')
const realTimeDataRetention = ref<number>(7)
const historicalDataRetention = ref<number>(30)
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
</script>
