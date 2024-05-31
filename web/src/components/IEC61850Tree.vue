<template>
  <div class="container mx-auto px-4 py-8">
    <h2 class="text-3xl font-bold mb-6">{{ $t('iec61850ExplorerTitle') }}</h2>
    <div class="flex border rounded-lg shadow-md">
      <div class="w-1/4 bg-gray-100 p-6 border-r">
        <h3 class="text-xl font-semibold mb-4">{{ $t('iec61850Model') }}</h3>
        <div class="mb-4">
          <input
            type="text"
            v-model="searchQuery"
            :placeholder="$t('searchPlaceholder')"
            class="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>
        <div class="flex items-center space-x-2">
          <span class="text-gray-800 font-medium">{{ model.name }}</span>
        </div>
        <ul v-if="model.logicalDevice && model.logicalDevice.length" class="ml-6 mt-2 space-y-2">
          <li v-for="(ld, ldIndex) in model.logicalDevice" :key="ldIndex">
            <div class="flex items-center space-x-2">
              <span class="text-gray-600">{{ ld.name }}</span>
            </div>
            <ul v-if="ld.logicalNode && ld.logicalNode.length" class="ml-10 mt-2 space-y-1">
              <li v-for="(ln, lnIndex) in ld.logicalNode" :key="lnIndex">
                <a
                  href="#"
                  @click.prevent="setlnId(ln.id)"
                  :class="{ 'text-blue-600': lnId === ln.id }"
                  class="hover:text-blue-600 focus:outline-none"
                >
                  {{ ln.name }}
                </a>
              </li>
            </ul>
          </li>
        </ul>
      </div>

      <div class="w-3/4 p-6">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-xl font-semibold">{{ $t('logicalNodeDetails') }}</h3>
          <div>
            <button
              class="px-4 py-2 bg-blue-500 text-white rounded-md mr-2 hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500"
              @click="exportData"
            >
              {{ $t('export') }}
            </button>
            <button
              class="px-4 py-2 bg-green-500 text-white rounded-md hover:bg-green-600 focus:outline-none focus:ring-2 focus:ring-green-500"
              @click="refreshData"
            >
              {{ $t('refresh') }}
            </button>
          </div>
        </div>
        <div v-if="dataObjects">
          <table class="w-full border-collapse bg-white rounded-lg shadow-sm">
            <thead>
              <tr class="bg-gray-200 text-gray-700">
                <th class="px-4 py-2 font-medium text-left">{{ $t('doName') }}</th>
                <th class="px-4 py-2 font-medium text-left">{{ $t('daName') }}</th>
                <th class="px-4 py-2 font-medium text-left">{{ $t('reference') }}</th>
                <th class="px-4 py-2 font-medium text-left">{{ $t('value') }}</th>
                <th class="px-4 py-2 font-medium text-left">{{ $t('dataSource') }}</th>
              </tr>
            </thead>
            <tbody>
              <template v-for="(dObj, dObjIndex) in dataObjects" :key="dObjIndex">
                <tr
                  v-for="(dAttr, dAttrIndex) in dObj.dataAttribute"
                  :key="dAttrIndex"
                  class="border-t"
                >
                  <td
                    v-if="dAttrIndex === 0"
                    :rowspan="dObj.dataAttribute.length"
                    class="px-4 py-2 align-top"
                  >
                    {{ dObj.name }}
                  </td>
                  <td class="px-4 py-2">{{ dAttr.name }}</td>
                  <td class="px-4 py-2">{{ dAttr.iec61850Ref }}</td>
                  <td class="px-4 py-2">{{ dAttr.value }}</td>
                  <td class="px-4 py-2">{{ dAttr.dataSource }}</td>
                </tr>
              </template>
            </tbody>
          </table>
        </div>
        <p v-else class="text-gray-600">{{ $t('selectNode') }}</p>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import type { IEC61850Model, DataObject } from '@/types/types'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'

const router = useRouter()
const searchQuery = useI18n().t('searchPlaceholder')
let model = ref<IEC61850Model>({} as IEC61850Model)
const lnId = ref<number>()
const modelId = router.currentRoute.value.params.id

onMounted(async () => {
  const id = Number(modelId)
  model.value = await fetchModelById(id)
})

let dataObjects = ref<DataObject[]>([])

async function setlnId(id: number) {
  lnId.value = id
  dataObjects.value = await fetchDataObjectsByLogicalNodeId(id)
}

async function fetchModelById(id: number): Promise<IEC61850Model> {
  const response = await fetch(`/api/iec61850/model/${id}`)
  const data = await response.json()
  return data
}

async function fetchDataObjectsByLogicalNodeId(id: number): Promise<DataObject[]> {
  const response = await fetch(`/api/iec61850/logical_node/${id}/data_object`)
  const data = await response.json()
  return data
}

function exportData() {
  console.log('Exporting data...')
}

function refreshData() {
  console.log('Refreshing data...')
}
</script>
