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
        <ul class="space-y-4">
          <li v-for="(ied, iedIndex) in filteredModel" :key="iedIndex">
            <div class="flex items-center space-x-2">
              <span class="text-gray-800 font-medium">{{ ied.name }}</span>
            </div>
            <ul v-if="ied.logicalDevice && ied.logicalDevice.length" class="ml-6 mt-2 space-y-2">
              <li v-for="(ld, ldIndex) in ied.logicalDevice" :key="ldIndex">
                <div class="flex items-center space-x-2">
                  <span class="text-gray-600">{{ ld.name }}</span>
                </div>
                <ul v-if="ld.logicalNode && ld.logicalNode.length" class="ml-10 mt-2 space-y-1">
                  <li v-for="(ln, lnIndex) in ld.logicalNode" :key="lnIndex">
                    <a
                      href="#"
                      @click.prevent="selectLogicalNode(ln.name)"
                      :class="{ 'text-blue-600': selectedLN === ln.name }"
                      class="hover:text-blue-600 focus:outline-none"
                    >
                      {{ ln.name }}
                    </a>
                  </li>
                </ul>
              </li>
            </ul>
          </li>
        </ul>
      </div>

      <!-- Right Side: Logical Node Details -->
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
        <template v-if="selectedLogicalNode">
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
              <template
                v-for="(dObj, dObjIndex) in selectedLogicalNode.dataObjects"
                :key="dObjIndex"
              >
                <tr
                  v-for="(dAttr, dAttrIndex) in dObj.dataAttributes"
                  :key="dAttrIndex"
                  class="border-t"
                >
                  <td
                    v-if="dAttrIndex === 0"
                    :rowspan="dObj.dataAttributes.length"
                    class="px-4 py-2 align-top"
                  >
                    {{ dObj.name }}
                  </td>
                  <td class="px-4 py-2">{{ dAttr.name }}</td>
                  <td class="px-4 py-2">{{ dAttr.ref }}</td>
                  <td class="px-4 py-2">{{ dAttr.value }}</td>
                  <td class="px-4 py-2">{{ dAttr.dataSource }}</td>
                </tr>
              </template>
            </tbody>
          </table>
        </template>
        <p v-else class="text-gray-600">{{ $t('selectNode') }}</p>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed, reactive } from 'vue'
import { userGlobalStore } from '@/store'

const store = userGlobalStore()
let model = reactive(store.model)
let logicalNodes = reactive(store.logicalNodes)

const selectedLN = ref('')
const searchQuery = ref('')

const filteredModel = computed(() => {
  if (searchQuery.value === '') {
    return model
  } else {
    return model.filter((ied) => {
      return ied.name.toLowerCase().includes(searchQuery.value.toLowerCase())
    })
  }
})

const selectedLogicalNode = computed(() => {
  return logicalNodes.find((ln) => ln.name === selectedLN.value)
})

function selectLogicalNode(lnName: string) {
  selectedLN.value = lnName
}

function exportData() {
  // Implement export functionality here
  console.log('Exporting data...')
}

function refreshData() {
  // Implement refresh functionality here
  console.log('Refreshing data...')
}
</script>
