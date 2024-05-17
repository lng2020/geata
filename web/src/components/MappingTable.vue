<template>
  <div class="container mx-auto px-4 py-8">
    <h2 class="text-3xl font-bold mb-6">Mapping Table</h2>
    <table class="container border rounded-lg shadow-md">
      <thead>
        <tr>
          <th
            class="px-4 py-2 border-b border-gray-300 text-left text-xl font-semibold text-gray-600 uppercase tracking-wider"
          >
            IEC61850 Reference
          </th>
          <th
            class="px-4 py-2 border-b border-gray-300 text-left text-xl font-semibold text-gray-600 uppercase tracking-wider"
          >
            Mapping Type
          </th>
          <th
            class="px-4 py-2 border-b border-gray-300 text-left text-xl font-semibold text-gray-600 uppercase tracking-wider"
          >
            Details
          </th>
        </tr>
      </thead>
      <tbody class="bg-white">
        <tr v-for="mapping in mappings" :key="mapping.id" class="hover:bg-gray-100">
          <td class="px-4 py-2 border-b border-gray-300 text-sm text-gray-700">
            {{ mapping.iec61850Reference }}
          </td>
          <td class="px-4 py-2 border-b border-gray-300 text-sm text-gray-700">
            <select
              v-model="mapping.type"
              @change="handleTypeChange(mapping)"
              class="bg-white border border-gray-300 text-gray-700 py-1 px-2 rounded leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
            >
              <option value="MQTT">MQTT</option>
              <option value="Modbus">Modbus</option>
              <option value="IEC61850">IEC61850</option>
            </select>
          </td>
          <td class="px-4 py-2 border-b border-gray-300 text-sm text-gray-700">
            <template v-if="mapping.type !== 'IEC61850'">
              <button
                @click="editMapping(mapping)"
                class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-1 px-2 rounded focus:outline-none focus:shadow-outline"
              >
                Edit
              </button>
            </template>
            <template v-else>
              <span class="text-gray-500">\</span>
            </template>
          </td>
        </tr>
      </tbody>
    </table>

    <ConfirmationModal
      :show="showModal"
      :title="modalTitle"
      :message="modalMessage"
      @confirm="confirmAction"
      @cancel="closeModal"
    />
    <MQTTConfigModal v-if="showMQTTConfig" @save="saveMQTTConfig" @cancel="closeMQTTConfig" />
    <ModbusConfigModal
      v-if="showModbusConfig"
      @save="saveModbusConfig"
      @cancel="closeModbusConfig"
    />
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive } from 'vue'
import type { Mapping } from '@/types/types'
import { userGlobalStore } from '@/store'
import ConfirmationModal from './ConfirmationModal.vue'
import MQTTConfigModal from './MQTTConfigModal.vue'
import ModbusConfigModal from './ModbusConfigModal.vue'

const store = userGlobalStore()
const mappings = reactive(store.mappings)

const showModal = ref(false)
const modalTitle = ref('Confirmation')
const modalMessage = ref('Are you sure you want to select IEC61850?')
const selectedMapping = ref<Mapping | null>(null)
const showMQTTConfig = ref(false)
const showModbusConfig = ref(false)

const editMapping = (mapping: Mapping) => {
  showConfigModal(mapping)
}

const handleTypeChange = (mapping: Mapping) => {
  if (mapping.type === 'IEC61850') {
    selectedMapping.value = mapping
    showModal.value = true
  } else {
    showConfigModal(mapping)
  }
}

const showConfigModal = (mapping: Mapping) => {
  if (mapping.type === 'MQTT') {
    showMQTTConfig.value = true
  } else if (mapping.type === 'Modbus') {
    showModbusConfig.value = true
  }
}

const saveMQTTConfig = (config: { topic: string; qos: string }) => {
  console.log('MQTT Config', config)
  showMQTTConfig.value = false
}

const closeMQTTConfig = () => {
  showMQTTConfig.value = false
}

const saveModbusConfig = (config: { address: number; dataType: string }) => {
  console.log('Modbus Config', config)
  showModbusConfig.value = false
}

const closeModbusConfig = () => {
  showModbusConfig.value = false
}

const confirmAction = () => {
  if (selectedMapping.value) {
    closeModal()
  }
}

const closeModal = () => {
  showModal.value = false
  selectedMapping.value = null
}
</script>
