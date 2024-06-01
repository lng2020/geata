<template>
  <div class="flex items-center justify-center space-x-4 p-2">
    <button
      :disabled="currentPage === 1"
      @click="emit('update:currentPage', 1)"
      class="px-2 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
    >
      &laquo;
    </button>
    <button
      :disabled="currentPage === 1"
      @click="emit('update:currentPage', currentPage - 1)"
      class="px-2 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
    >
      &lsaquo;
    </button>
    <span class="text-sm text-gray-700">
      {{ currentPage }} / {{ totalPages }}
    </span>
    <button
      :disabled="currentPage === totalPages"
      @click="emit('update:currentPage', currentPage + 1)"
      class="px-2 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
    >
      &rsaquo;
    </button>
    <button
      :disabled="currentPage === totalPages"
      @click="emit('update:currentPage', totalPages)"
      class="px-2 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
    >
      &raquo;
    </button>
  </div>
</template>

<script lang="ts" setup>
import { computed } from 'vue'
const props = defineProps({
  currentPage: {
    type: Number,
    required: true
  },
  totalItems: {
    type: Number,
    required: true
  },
  pageSize: {
    type: Number,
    default: 10
  }
})

const emit = defineEmits(['update:currentPage'])

const totalPages = computed(() => Math.ceil(props.totalItems / props.pageSize))
</script>
