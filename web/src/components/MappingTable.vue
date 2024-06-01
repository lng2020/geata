<template>
  <div class="container mx-auto px-2 py-2">
    <div class="flex justify-between">
      <h2 class="text-3xl font-bold mb-6">{{ $t('dataSourceMappingTitle') }}</h2>
      <div class="mb-4">
        <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mr-2">
          {{ $t('batchEdit') }}
        </button>
        <button class="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded mr-2">
          {{ $t('importExport') }}
        </button>
        <button
          @click="showUtilsModal"
          class="bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-4 rounded"
        >
          {{ $t('utils') }}
        </button>
      </div>
    </div>
    <table class="container border rounded-lg shadow-md">
      <thead>
        <tr>
          <th
            class="px-4 py-2 border-b border-gray-300 text-left text-xl font-semibold text-gray-600 uppercase tracking-wider"
          >
            {{ $t('logicalNode') }}
          </th>
          <th
            class="px-4 py-2 border-b border-gray-300 text-left text-xl font-semibold text-gray-600 uppercase tracking-wider"
          >
            {{ $t('dataSourceType') }}
          </th>
          <th
            class="px-4 py-2 border-b border-gray-300 text-left text-xl font-semibold text-gray-600 uppercase tracking-wider"
          >
            {{ $t('configuration') }}
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
              <option value="IEC61850">{{ $t('iec61850') }}</option>
              <option value="Modbus">{{ $t('modbus') }}</option>
              <option value="MQTT">{{ $t('mqtt') }}</option>
            </select>
          </td>
          <td class="px-4 py-2 border-b border-gray-300 text-sm text-gray-700">
            <template v-if="mapping.type !== 'IEC61850'">
              <button
                @click="editMapping(mapping)"
                class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-1 px-2 rounded focus:outline-none focus:shadow-outline"
              >
                {{ $t('edit') }}
              </button>
            </template>
            <template v-else>
              <span class="text-gray-500">\</span>
            </template>
          </td>
        </tr>
      </tbody>
    </table>

    <UtilsModal v-if="showUtils" @close="closeUtilsModal" />
    <ConfirmationModal
      :show="showModal"
      :title="$t('confirmation')"
      :message="$t('areYouSureToChangeType')"
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
import { useGlobalStore } from '@/store'
import ConfirmationModal from './ConfirmationModal.vue'
import MQTTConfigModal from './MQTTConfigModal.vue'
import ModbusConfigModal from './ModbusConfigModal.vue'
import UtilsModal from './UtilsModal.vue'

const store = useGlobalStore()
const mappings = reactive(store.mappings)

const showModal = ref(false)
const selectedMapping = ref<Mapping | null>(null)
const showMQTTConfig = ref(false)
const showModbusConfig = ref(false)
const showUtils = ref(false)

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

const showUtilsModal = () => {
  showUtils.value = true
}

const closeUtilsModal = () => {
  showUtils.value = false
}
</script>
