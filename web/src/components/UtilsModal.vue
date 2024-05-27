<template>
  <div class="fixed z-10 inset-0 overflow-y-auto">
    <div
      class="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0"
    >
      <div class="fixed inset-0 transition-opacity" aria-hidden="true">
        <div class="absolute inset-0 bg-gray-500 opacity-75"></div>
      </div>

      <div
        class="inline-block align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full"
        role="dialog"
        aria-modal="true"
        aria-labelledby="modal-headline"
      >
        <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
          <div class="sm:flex sm:items-start">
            <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
              <h3 class="text-lg leading-6 font-medium text-gray-900" id="modal-headline">Utils</h3>
              <div class="mt-4">
                <div class="mb-6">
                  <h4 class="text-xl font-semibold mb-2">Modbus Address Calculator</h4>
                  <div class="flex items-center">
                    <label class="mr-2">Start Address:</label>
                    <input
                      type="text"
                      v-model="modbusStartAddress"
                      class="border border-gray-300 rounded py-2 px-3 mr-2"
                    />
                    <button
                      @click="calculateModbusAddress"
                      class="bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-4 rounded"
                    >
                      Calculate
                    </button>
                  </div>
                  <div v-if="modbusCalculatedAddress" class="mt-2">
                    Calculated Address: {{ modbusCalculatedAddress }}
                  </div>
                </div>
                <div>
                  <h4 class="text-xl font-semibold mb-2">MQTT Topic Subscription Test</h4>
                  <div class="flex items-center mb-2">
                    <label class="mr-2">Topic:</label>
                    <input
                      type="text"
                      v-model="mqttTestTopic"
                      class="border border-gray-300 rounded py-2 px-3 mr-2"
                    />
                    <button
                      @click="testMqttSubscription"
                      class="bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-4 rounded mr-2"
                    >
                      Subscribe
                    </button>
                    <button
                      @click="publishMqttMessage"
                      class="bg-green-500 hover:bg-green-600 text-white font-semibold py-2 px-4 rounded"
                    >
                      Publish
                    </button>
                  </div>
                  <div v-if="mqttTestMessage" class="mt-2">
                    Received Message: {{ mqttTestMessage }}
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="bg-gray-50 px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse">
          <button
            @click="$emit('close')"
            type="button"
            class="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm"
          >
            Close
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'

const modbusStartAddress = ref('')
const modbusCalculatedAddress = ref('')
const mqttTestTopic = ref('')
const mqttTestMessage = ref('')

const calculateModbusAddress = () => {
  // ...
}

const testMqttSubscription = () => {
  // ...
}

const publishMqttMessage = () => {
  // ...
}
</script>
