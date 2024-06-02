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
        <tr v-for="m in paginatedData" :key="m.id" class="hover:bg-gray-100">
          <td class="px-4 py-2 border-b border-gray-300 text-sm text-gray-700">
            {{ m.iec61850Ref }}
          </td>
          <td class="px-4 py-2 border-b border-gray-300 text-sm text-gray-700">
            <select
              v-model="m.type"
              @change="handleTypeChange(m)"
              class="bg-white border border-gray-300 text-gray-700 py-1 px-2 rounded leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
            >
              <option value="IEC61850">{{ $t('iec61850') }}</option>
              <option value="Modbus">{{ $t('modbus') }}</option>
              <option value="MQTT">{{ $t('mqtt') }}</option>
            </select>
          </td>
          <td class="px-4 py-2 border-b border-gray-300 text-sm text-gray-700">
            <template v-if="m.type !== 'IEC61850'">
              <button
                @click="editMapping(m)"
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
    <div v-if="totalItems > pageSize">
      <Pagination
        :current-page="currentPage"
        :total-items="totalItems"
        :page-size="pageSize"
        @update:current-page="currentPage = $event"
      />
    </div>
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
import { ref } from 'vue'
import { onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useGlobalStore } from '@/store'
import type { Mapping } from '@/types/types'
import ConfirmationModal from '@/components/ConfirmationModal.vue'
import MQTTConfigModal from '@/components/MQTTConfigModal.vue'
import ModbusConfigModal from '@/components/ModbusConfigModal.vue'
import Pagination from '@/components/Pagination.vue'
import UtilsModal from '@/components/UtilsModal.vue'

const router = useRouter()
const store = useGlobalStore()
let mapping = ref<Mapping[]>([])
const showModal = ref(false)
const selectedMapping = ref<Mapping | null>(null)
const showMQTTConfig = ref(false)
const showModbusConfig = ref(false)
const showUtils = ref(false)
const currentPage = ref(1)
const pageSize = ref(12)
const totalItems = computed(() => mapping.value.length)

const paginatedData = computed(() => {
  const start = (currentPage.value - 1) * pageSize.value
  const end = start + pageSize.value
  return mapping.value.slice(start, end)
})

onMounted(async () => {
  const stationId = Number(router.currentRoute.value.params.id)
  if (!stationId) {
    return
  }
  const station = store.stations.find((s) => s.id === stationId)
  if (!station) {
    return
  }
  const modelId = station.modelId
  mapping.value = await fetchMapping(modelId)
})

const fetchMapping = async (modelId: number) => {
  try {
    const response = await fetch(`/api/iec61850/model/${modelId}/mapping_rule`)
    return await response.json()
  } catch (error) {
    console.error('Error:', error)
  }
}
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
