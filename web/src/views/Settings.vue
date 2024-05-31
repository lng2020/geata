<template>
  <div class="container mx-auto px-2 py-2">
    <div class="flex justify-between mb-6">
      <h1 class="text-3xl font-bold">{{ $t('settings') }}</h1>
      <button
        @click="deleteStation"
        class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
        type="button"
      >
        {{ $t('deleteStation') }}
      </button>
    </div>
    <div class="flex">
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
              class="text-sm bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
              type="button"
            >
              {{ $t('save') }}
            </button>
            <button
              @click="cancel"
              class="text-sm bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
              type="button"
            >
              {{ $t('cancel') }}
            </button>
          </div>
        </div>
        <div class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
          <div class="flex justify-between items-center">
            <h2 class="text-xl font-bold">{{ $t('alarmRules') }}</h2>
            <button
              @click="showAddRuleDialog = true"
              class="text-sm px-4 py-2 bg-green-500 text-white rounded-md hover:bg-green-600 focus:outline-none focus:shadow-outline"
            >
              {{ $t('addRule') }}
            </button>
          </div>
          <ul class="mt-4 space-y-4">
            <li
              v-for="(rule, index) in alarmRules"
              :key="index"
              class="flex items-center justify-between bg-gray-100 p-4 rounded-md"
            >
              <div class="flex items-center space-x-2">
                <span>{{ rule.condition }}</span>
                <span class="text-gray-500">|</span>
                <span>{{ rule.notification }}</span>
              </div>
              <div class="flex space-x-2">
                <button
                  @click="editRule(index)"
                  class="text-sm px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 focus:outline-none focus:shadow-outline"
                >
                  {{ $t('edit') }}
                </button>
                <button
                  @click="deleteRule(index)"
                  class="text-sm px-4 py-2 bg-red-500 text-white rounded-md hover:bg-red-600 focus:outline-none focus:shadow-outline"
                >
                  {{ $t('delete') }}
                </button>
              </div>
            </li>
          </ul>
        </div>
      </div>
      <div class="w-1/2 pl-4">
        <div class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
          <h2 class="text-xl font-bold mb-4">{{ $t('dataSourceConfiguration') }}</h2>
          <ul class="space-y-4">
            <li>
              <section class="bg-white shadow rounded-lg p-6">
                <div class="flex justify-between">
                  <h2 class="font-bold mb-4">{{ $t('MQTT') }}</h2>
                  <button
                    @click="testMqttConnection"
                    class="text-sm mb-4 px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 focus:outline-none focus:shadow-outline"
                  >
                    {{ $t('testConnection') }}
                  </button>
                </div>
                <div>
                  <div class="grid grid-cols-2 gap-4">
                    <div class="mb-4">
                      <label class="block text-gray-700 font-bold text-sm mb-2">{{
                        $t('address')
                      }}</label>
                      <input
                        type="text"
                        v-model="mqttServerAddress"
                        class="w-full border border-gray-400 hover:border-gray-500 px-2 py-1 rounded shadow focus:outline-none focus:shadow-outline"
                      />
                    </div>
                    <div>
                      <label class="block text-gray-700 font-bold text-sm mb-2">{{
                        $t('port')
                      }}</label>
                      <input
                        type="number"
                        v-model="mqttServerPort"
                        class="w-full border border-gray-400 hover:border-gray-500 px-2 py-1 rounded shadow focus:outline-none focus:shadow-outline"
                      />
                    </div>
                    <div>
                      <label class="block text-gray-700 font-bold text-sm mb-2">{{
                        $t('username')
                      }}</label>
                      <input
                        type="text"
                        v-model="mqttUsername"
                        class="w-full border border-gray-400 hover:border-gray-500 px-2 py-1 rounded shadow focus:outline-none focus:shadow-outline"
                      />
                    </div>
                    <div>
                      <label class="block text-gray-700 font-bold text-sm mb-2">{{
                        $t('password')
                      }}</label>
                      <input
                        type="password"
                        v-model="mqttPassword"
                        class="w-full border border-gray-400 hover:border-gray-500 px-2 py-1 rounded shadow focus:outline-none focus:shadow-outline"
                      />
                    </div>
                  </div>
                </div>
              </section>
            </li>
            <!-- IEC61850 Configuration -->
            <li>
              <section class="bg-white shadow rounded-lg p-6">
                <div class="flex justify-between">
                  <h2 class="font-bold mb-4">{{ $t('IEC61850') }}</h2>
                  <button
                    @click="testIec61850Connection"
                    class="text-sm mb-4 px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 focus:outline-none focus:shadow-outline"
                  >
                    {{ $t('testConnection') }}
                  </button>
                </div>
                <div class="grid grid-cols-2 gap-4">
                  <div class="mb-4">
                    <label class="block text-gray-700 font-bold text-sm mb-2">{{
                      $t('host')
                    }}</label>
                    <input
                      type="text"
                      v-model="iec61850Host"
                      class="w-full border border-gray-400 hover:border-gray-500 px-2 py-1 rounded shadow focus:outline-none focus:shadow-outline"
                    />
                  </div>
                  <div class="mb-4">
                    <label class="block text-gray-700 font-bold text-sm mb-2">{{
                      $t('port')
                    }}</label>
                    <input
                      type="number"
                      v-model="iec61850Port"
                      class="w-full border border-gray-400 hover:border-gray-500 px-2 py-1 rounded shadow focus:outline-none focus:shadow-outline"
                    />
                  </div>
                </div>
              </section>
            </li>
            <!-- Modbus Configuration -->
            <li>
              <section class="bg-white shadow rounded-lg p-6">
                <div class="flex justify-between">
                  <h2 class="font-bold mb-4">{{ $t('Modbus') }}</h2>
                  <button
                    @click="testModbusConnection"
                    class="text-sm mb-4 px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 focus:outline-none focus:shadow-outline"
                  >
                    {{ $t('testConnection') }}
                  </button>
                </div>
                <div>
                  <div class="mb-4">
                    <label class="block text-gray-700 font-bold mb-2">{{ $t('url') }}</label>
                    <input
                      type="text"
                      v-model="modbusUrl"
                      class="w-full border border-gray-400 hover:border-gray-500 px-2 py-1 rounded shadow focus:outline-none focus:shadow-outline"
                    />
                  </div>
                </div>
              </section>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
  <!-- Add Rule Dialog -->
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
</template>

<script setup lang="ts">
import { useGlobalStore } from '@/store'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import type { Station } from '@/types/types'

interface Rule {
  condition: string
  notification: string
}
const store = useGlobalStore()
const router = useRouter()
const stationId = Number(router.currentRoute.value.params.id)
const station = store.stations.find((s) => s.id === stationId)
if (!station) {
  router.push('/error')
}
let rawStation = ref<Station>(station as Station)
let editedStation = ref<Station>(station as Station)
const mqttServerAddress = ref('')
const mqttServerPort = ref(0)
const mqttUsername = ref('')
const mqttPassword = ref('')

const iec61850Host = ref('')
const iec61850Port = ref(0)

const modbusUrl = ref('')

const testMqttConnection = () => {
  // TODO: Implement MQTT connection test
}

const testIec61850Connection = () => {
  // TODO: Implement IEC61850 connection test
}

const testModbusConnection = () => {
  // TODO: Implement Modbus connection test
}

const saveChanges = () => {
  if (editedStation.value) {
    rawStation.value = editedStation.value
  }
}

const cancel = () => {
  if (rawStation.value) {
    editedStation.value = rawStation.value
  }
}

const alarmRules = ref<Rule[]>([
  { condition: 'Temperature > 30', notification: 'Send email' },
  { condition: 'Humidity < 40%', notification: 'Send SMS' }
])
const showAddRuleDialog = ref<boolean>(false)
const newRule = ref<Rule>({ condition: '', notification: '' })

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

function deleteStation() {
  console.log('Deleting station')
}
</script>
