<template>
  <table class="min-w-full border-collapse border border-gray-300 mt-4">
    <thead>
      <tr>
        <th class="px-4 py-2 border-b border-gray-300 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">IEC61850 Reference</th>
        <th class="px-4 py-2 border-b border-gray-300 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Mapping Type</th>
        <th class="px-4 py-2 border-b border-gray-300 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">Details</th>
      </tr>
    </thead>
    <tbody class="bg-white">
      <tr v-for="mapping in mappings" :key="mapping.id" class="hover:bg-gray-100">
        <td class="px-4 py-2 border-b border-gray-300 text-sm text-gray-700">{{ mapping.iec61850Reference }}</td>
        <td class="px-4 py-2 border-b border-gray-300 text-sm text-gray-700">{{ mapping.type }}</td>
        <td class="px-4 py-2 border-b border-gray-300 text-sm text-gray-700">
          <template v-if="mapping.type !== 'IEC61850'">
            <button  @click="editMapping(mapping)" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-1 px-2 rounded focus:outline-none focus:shadow-outline">
              Detail
            </button>
          </template>
          <template v-else>
            <span class="text-gray-500">\</span>
          </template>
        </td>
      </tr>
    </tbody>
  </table>
</template>
  
<script lang="ts" setup>

type MappingType = "MQTT" | "ModbusTCP" | "IEC61850";

interface Mapping {
  id: number;
  iec61850Reference: string;
  type: MappingType;
}

const mappings: Mapping[] = [
  {
    id: 1,
    iec61850Reference: "LD0/LLN0$ST$stVal",
    type: 'MQTT'
  },
  {
    id: 2,
    iec61850Reference: "LD0/LLN0$ST$stVal",
    type: 'ModbusTCP'
  },
  {
    id: 3,
    iec61850Reference: "LD0/LLN0$ST$stVal",
    type: 'IEC61850'
  }
];

const editMapping = (mapping: Mapping) => {
  console.log(mapping);
};

</script>

<style scoped >

</style>
