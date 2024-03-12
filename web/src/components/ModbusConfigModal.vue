<template>
  <div class="fixed inset-0 flex items-center justify-center z-50">
    <div class="absolute inset-0 bg-gray-900 opacity-50"></div>
    <div class="bg-white rounded-lg p-6 m-4 max-w-md z-10">
      <h2 class="text-xl font-bold mb-4">Modbus TCP Configuration</h2>
      <form @submit.prevent="saveConfig">
        <div class="mb-4">
          <label for="address" class="block text-gray-700 font-bold mb-2">Address:</label>
          <input
            type="number"
            id="address"
            v-model="address"
            class="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:border-blue-500"
            required
          />
        </div>
        <div class="mb-4">
          <label for="dataType" class="block text-gray-700 font-bold mb-2">Data Type:</label>
          <select
            id="dataType"
            v-model="dataType"
            class="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:border-blue-500"
            required
          >
            <option value="int16">INT16</option>
            <option value="int32">INT32</option>
            <option value="float">FLOAT</option>
            <!-- Add more data type options as needed -->
          </select>
        </div>
        <div class="flex justify-end">
          <button
            type="submit"
            class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded mr-2"
          >
            Save
          </button>
          <button
            type="button"
            @click="cancel"
            class="bg-gray-300 hover:bg-gray-400 text-gray-800 font-bold py-2 px-4 rounded"
          >
            Cancel
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'

const address = ref(0)
const dataType = ref('int16')

const emit = defineEmits<{
  (event: 'save', config: { address: number; dataType: string }): void
  (event: 'cancel'): void
}>()

const saveConfig = () => {
  emit('save', { address: address.value, dataType: dataType.value })
}

const cancel = () => {
  emit('cancel')
}
</script>
