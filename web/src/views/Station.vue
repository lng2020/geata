<template>
  <div class="container mx-auto px-2 py-2">
    <h2 class="text-3xl font-bold mb-6">{{ $t('iec61850ExplorerTitle') }}</h2>
    <div class="flex border rounded-lg shadow-md bg-white">
      <div class="w-1/4 p-6 border-r overflow-y-auto max-h-[750px]">
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
              <template v-for="(dObj, dObjIndex) in paginatedDataObjects" :key="dObjIndex">
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
        <div v-if="dataObjects.length > pageSize">
          <Pagination
            :currentPage="currentPage"
            :totalItems="dataObjects.length"
            :pageSize="pageSize"
            @update:currentPage="currentPage = $event"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted, onUnmounted, computed } from 'vue'
import type { IEC61850Model, DataObject } from '@/types/types'
import { useRouter, onBeforeRouteUpdate } from 'vue-router'
import { useI18n } from 'vue-i18n'
import Pagination from '@/components/Pagination.vue'

const router = useRouter()
const searchQuery = useI18n().t('searchPlaceholder')
let model = ref<IEC61850Model>({} as IEC61850Model)
const lnId = ref<number>()
let modelId = ref<number>()
const currentPage = ref(1)
const pageSize = ref(8)
let refreshInterval: number | undefined

onMounted(async () => {
  const id = Number(router.currentRoute.value.params.id)
  modelId.value = id
  model.value = await fetchModelById(id)
  setlnId(model.value.logicalDevice[0].logicalNode[0].id)
  startDataRefresh()
})

onBeforeRouteUpdate(async (to, from) => {
  if (to.params.id !== from.params.id) {
    const id = Number(to.params.id)
    modelId.value = id
    model.value = await fetchModelById(id)
    if (lnId.value) {
      setlnId(lnId.value)
    } else {
      setlnId(model.value.logicalDevice[0].logicalNode[0].id)
    }
  }
})

onUnmounted(() => {
  stopDataRefresh()
})


function startDataRefresh() {
  refreshInterval = setInterval(async () => {
    if (lnId.value) {
      dataObjects.value = await fetchDataObjectsByLogicalNodeId(lnId.value)
    }
  }, 1000)
}

function stopDataRefresh() {
  if (refreshInterval) {
    clearInterval(refreshInterval)
  }
}

onBeforeRouteUpdate(async (to, from) => {
  if (to.params.id !== from.params.id) {
    const id = Number(to.params.id)
    modelId.value = id
    model.value = await fetchModelById(id)
    if(lnId.value){
      setlnId(lnId.value)
    }else {
      setlnId(model.value.logicalDevice[0].logicalNode[0].id)
    }
  }
})

let dataObjects = ref<DataObject[]>([])
const paginatedDataObjects = computed(() => {
  const start = (currentPage.value - 1) * pageSize.value
  const end = start + pageSize.value
  return dataObjects.value.slice(start, end)
})

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
