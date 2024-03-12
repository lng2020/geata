<template>
  <div class="p-4">
    <IEC61850Tree />
  </div>
</template>

<script lang="ts">
export default {
  name: 'Station'
}
</script>

<script lang="ts" setup>
import IEC61850Tree from '@/components/IEC61850Tree.vue'
</script>

<style scoped></style>
