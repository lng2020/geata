<template>
  <table class="min-w-full">
    <thead>
      <tr>
        <th class="px-4 py-2">Mapping Type</th>
        <th class="px-4 py-2">Details</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="mapping in mappings" :key="mapping.id">
        <td class="px-4 py-2">{{ mapping.type }}</td>
        <td class="px-4 py-2">
          <button @click="editMapping(mapping)" class="text-blue-500 underline">
            Detail/Edit
          </button>
        </td>
      </tr>
    </tbody>
  </table>
</template>
  
<script lang="ts" setup>

enum MappingType {
  MQTT,
  ModbusTCP
}

interface Mapping {
  id: number;
  type: MappingType;
}

const mappings: Mapping[] = [
  {
    id: 1,
    type: MappingType.MQTT
  },
  {
    id: 2,
    type: MappingType.ModbusTCP
  }
];

const editMapping = (mapping: Mapping) => {
  console.log(mapping);
};

</script>

<style scoped >

</style>
