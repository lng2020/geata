<!-- MQTTConfig.vue -->
<template>
    <div class="fixed inset-0 flex items-center justify-center z-50">
      <div class="absolute inset-0 bg-gray-900 opacity-50"></div>
      <div class="bg-white rounded-lg p-6 m-4 max-w-md z-10">
        <h2 class="text-xl font-bold mb-4">MQTT Configuration</h2>
        <form @submit.prevent="saveConfig">
          <div class="mb-4">
            <label for="topic" class="block text-gray-700 font-bold mb-2">Topic:</label>
            <input
              type="text"
              id="topic"
              v-model="topic"
              class="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:border-blue-500"
              required
            />
          </div>
          <div class="mb-4">
            <label for="qos" class="block text-gray-700 font-bold mb-2">QoS:</label>
            <select
              id="qos"
              v-model="qos"
              class="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:border-blue-500"
              required
            >
              <option value="0">0</option>
              <option value="1">1</option>
              <option value="2">2</option>
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
  import { ref, defineEmits } from 'vue';
  
  const topic = ref('');
  const qos = ref('0');
  
  const emit = defineEmits<{
    (event: 'save', config: { topic: string; qos: string }): void;
    (event: 'cancel'): void;
  }>();
  
  const saveConfig = () => {
    emit('save', { topic: topic.value, qos: qos.value });
  };
  
  const cancel = () => {
    emit('cancel');
  };
  </script>