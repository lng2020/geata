<template>
  <div class="fixed inset-0 flex items-center justify-center z-50">
    <div class="absolute inset-0 bg-gray-900 opacity-50"></div>
    <div class="bg-white rounded-lg p-6 m-4 max-w-md z-10">
      <h2 class="text-xl font-bold mb-4">{{ $t('modbusConfig') }}</h2>
      <form @submit.prevent="saveConfig">
        <div class="mb-4">
          <label for="address" class="block text-gray-700 font-bold mb-2">{{ $t('addressLabel') }}</label>
          <input
            type="number"
            id="address"
            v-model="address"
            class="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:border-blue-500"
            required
          />
        </div>
        <div class="mb-4">
          <label for="dataType" class="block text-gray-700 font-bold mb-2">{{ $t('dataTypeLabel') }}</label>
          <select
            id="dataType"
            v-model="dataType"
            class="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:border-blue-500"
            required
          >
            <option value="int16">{{ $t('int16') }}</option>
            <option value="int32">{{ $t('int32') }}</option>
            <option value="float">{{ $t('float') }}</option>
          </select>
        </div>
        <div class="flex justify-end">
          <button
            type="submit"
            class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded mr-2"
          >
            {{ $t('save') }}
          </button>
          <button
            type="button"
            @click="cancel"
            class="bg-gray-300 hover:bg-gray-400 text-gray-800 font-bold py-2 px-4 rounded"
          >
            {{ $t('cancel') }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, defineEmits } from 'vue'

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