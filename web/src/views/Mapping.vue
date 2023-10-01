<template>
  <div class="p-4">
    <h1 class="text-2xl font-semibold mb-4">Mapping Web Page</h1>
    <MappingTable />
  </div>
</template>
  
<script lang="ts">
export default {
  name: "Mapping",
}
</script>

<script lang="ts" setup>
import MappingTable from "@/components/MappingTable.vue";
</script>
  
<style scoped></style>
  