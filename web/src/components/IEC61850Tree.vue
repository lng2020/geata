<template>
  <div class="container mx-auto px-4 py-8">
    <h2 class="text-3xl font-bold mb-6">IEC 61850 Explorer</h2>
    <div class="flex border rounded-lg shadow-md">
      <!-- Left Side: IED Model -->
      <div class="w-1/4 bg-gray-100 p-6 border-r">
        <h3 class="text-xl font-semibold mb-4">IED Model</h3>
        <ul class="space-y-4">
          <li v-for="(ied, iedIndex) in iedModel" :key="iedIndex">
            <div class="flex items-center space-x-2">
              <span class="text-gray-800 font-medium">{{ ied.name }}</span>
            </div>
            <ul v-if="ied.logicalDevices && ied.logicalDevices.length" class="ml-6 mt-2 space-y-2">
              <li v-for="(ld, ldIndex) in ied.logicalDevices" :key="ldIndex">
                <div class="flex items-center space-x-2">
                  <span class="text-gray-600">{{ ld.name }}</span>
                </div>
                <ul v-if="ld.logicalNodes && ld.logicalNodes.length" class="ml-10 mt-2 space-y-1">
                  <li v-for="(ln, lnIndex) in ld.logicalNodes" :key="lnIndex">
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
        <h3 class="text-xl font-semibold mb-4">Logical Node Details</h3>
        <template v-if="selectedLogicalNode">
          <table class="w-full border-collapse bg-white rounded-lg shadow-sm">
            <thead>
              <tr class="bg-gray-200 text-gray-700">
                <th class="px-4 py-2 font-medium text-left">DO Name</th>
                <th class="px-4 py-2 font-medium text-left">DA Name</th>
                <th class="px-4 py-2 font-medium text-left">Reference</th>
                <th class="px-4 py-2 font-medium text-left">Value</th>
                <th class="px-4 py-2 font-medium text-left">Data Source</th>
              </tr>
            </thead>
            <tbody>
              <template v-for="(dObj, dObjIndex) in selectedLogicalNode.dataObjects" :key="dObjIndex">
                <tr v-for="(dAttr, dAttrIndex) in dObj.dataAttributes" :key="dAttrIndex" class="border-t">
                  <td v-if="dAttrIndex === 0" :rowspan="dObj.dataAttributes.length" class="px-4 py-2 align-top">{{ dObj.name }}</td>
                  <td class="px-4 py-2">{{ dAttr.name }}</td>
                  <td class="px-4 py-2">{{ dAttr.ref }}</td>
                  <td class="px-4 py-2">{{ dAttr.value }}</td>
                  <td class="px-4 py-2">{{ dAttr.dataSource }}</td>
                </tr>
              </template>
            </tbody>
          </table>
        </template>
        <p v-else class="text-gray-600">Select a Logical Node to view its details.</p>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed } from 'vue';

const iedModel = ref([
  {
    name: 'IED1',
    logicalDevices: [
      {
        name: 'LD1',
        logicalNodes: [
          { name: 'LN1' },
          { name: 'LN2' },
        ],
      },
    ],
  },
]);

const logicalNodes = ref([
  {
    name: 'LN1',
    dataObjects: [
      {
        name: 'DO1',
        dataAttributes: [
          { name: 'DA1', ref: 'LD0/LLN0$ST$stVal', value: 'Value 1', dataSource: 'MQTT' },
          { name: 'DA2', ref: 'LD0/LLN0$ST$stVal', value: 'Value 2', dataSource: 'MQTT' },
        ],
      },
      {
        name: 'DO2',
        dataAttributes: [
          { name: 'DA3', ref: 'LD0/LLN0$ST$stVal', value: 'Value 3', dataSource: 'Modbus' },
        ],
      },
    ],
  },
  {
    name: 'LN2',
    dataObjects: [
      {
        name: 'DO3',
        dataAttributes: [
          { name: 'DA4', ref: 'LD0/LLN0$ST$stVal', value: 'Value 4', dataSource: 'IEC 61850' },
        ],
      },
    ],
  },
]);

const selectedLN = ref('');

const selectedLogicalNode = computed(() => {
  return logicalNodes.value.find(ln => ln.name === selectedLN.value);
});

function selectLogicalNode(lnName: string) {
  selectedLN.value = lnName;
}
</script>