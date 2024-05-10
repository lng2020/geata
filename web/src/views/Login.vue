<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="max-w-md w-full px-6 py-8 bg-white shadow-md rounded-lg">
      <h2 class="text-2xl font-semibold mb-6 text-center">Login/Register</h2>

      <!-- Login Form -->
      <form @submit.prevent="login" v-if="!isRegistering">
        <div class="mb-4">
          <label for="username" class="block text-gray-700 font-bold mb-2">Username</label>
          <input
            type="text"
            id="username"
            v-model="username"
            class="w-full px-3 py-2 text-gray-700 border rounded-lg focus:outline-none"
            required
          />
        </div>
        <div class="mb-6">
          <label for="password" class="block text-gray-700 font-bold mb-2">Password</label>
          <input
            type="password"
            id="password"
            v-model="password"
            class="w-full px-3 py-2 text-gray-700 border rounded-lg focus:outline-none"
            required
          />
        </div>
        <div class="flex items-center justify-between">
          <button
            class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
            type="submit"
          >
            Login
          </button>
          <a
            class="inline-block align-baseline font-bold text-sm text-blue-500 hover:text-blue-800"
            href="#"
            @click.prevent="isRegistering = true"
          >
            Register
          </a>
        </div>
      </form>

      <!-- Registration Form -->
      <form @submit.prevent="register" v-else>
        <div class="mb-4">
          <label for="username" class="block text-gray-700 font-bold mb-2">Username</label>
          <input
            type="text"
            id="username"
            v-model="username"
            class="w-full px-3 py-2 text-gray-700 border rounded-lg focus:outline-none"
            required
          />
        </div>

        <div class="mb-6">
          <label for="password" class="block text-gray-700 font-bold mb-2">Password</label>
          <input
            type="password"
            id="password"
            v-model="password"
            class="w-full px-3 py-2 text-gray-700 border rounded-lg focus:outline-none"
            required
          />
        </div>
        <div class="flex items-center justify-between">
          <button
            class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
            type="submit"
          >
            Register
          </button>
          <a
            class="inline-block align-baseline font-bold text-sm text-blue-500 hover:text-blue-800"
            href="#"
            @click.prevent="isRegistering = false"
          >
            Login
          </a>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useAuthStore } from '@/stores/store'
import { useRouter } from 'vue-router'

const authStore = useAuthStore()
const router = useRouter()

const isRegistering = ref(false)
const username = ref('')
const password = ref('')

const login = async () => {
  await authStore.login(username.value, password.value)
  router.push('/')
}

const register = async () => {
  await authStore.register(username.value, password.value)
  router.push('/')
}
</script>
