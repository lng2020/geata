<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="max-w-md w-full px-6 py-8 bg-white shadow-md rounded-lg">
      <h2 class="text-2xl font-semibold mb-6 text-center">
        {{ isRegistering ? 'Register' : 'Login' }}
      </h2>
      <form @submit.prevent="handleSubmit">
        <div class="mb-4">
          <label for="username" class="block text-gray-700 font-bold mb-2">Username</label>
          <input
            type="text"
            id="username"
            v-model="username"
            class="w-full px-3 py-2 text-gray-700 border rounded-lg focus:outline-none"
            :class="{ 'border-red-500': usernameError }"
            required
          />
          <p v-if="usernameError" class="text-red-500 text-sm mt-1">{{ usernameError }}</p>
        </div>
        <div class="mb-6">
          <label for="password" class="block text-gray-700 font-bold mb-2">Password</label>
          <input
            type="password"
            id="password"
            v-model="password"
            class="w-full px-3 py-2 text-gray-700 border rounded-lg focus:outline-none"
            :class="{ 'border-red-500': passwordError }"
            required
          />
          <p v-if="passwordError" class="text-red-500 text-sm mt-1">{{ passwordError }}</p>
        </div>
        <div class="flex items-center justify-between">
          <button
            class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
            type="submit"
          >
            {{ isRegistering ? 'Register' : 'Login' }}
          </button>
          <a
            class="inline-block align-baseline font-bold text-sm text-blue-500 hover:text-blue-800"
            href="#"
            @click.prevent="toggleForm"
          >
            {{ isRegistering ? 'Login' : 'Register' }}
          </a>
        </div>
      </form>
      <p v-if="successMessage" class="text-green-500 text-center mt-4">{{ successMessage }}</p>
      <p v-if="errorMessage" class="text-red-500 text-center mt-4">{{ errorMessage }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useAuthStore } from '@/store';
import { useRouter } from 'vue-router';

const authStore = useAuthStore();
const router = useRouter();

const isRegistering = ref(false);
const username = ref('');
const password = ref('');
const usernameError = ref('');
const passwordError = ref('');
const successMessage = ref('');
const errorMessage = ref('');

const toggleForm = () => {
  isRegistering.value = !isRegistering.value;
  clearForm();
};

const clearForm = () => {
  username.value = '';
  password.value = '';
  usernameError.value = '';
  passwordError.value = '';
  successMessage.value = '';
  errorMessage.value = '';
};

const validateForm = () => {
  usernameError.value = username.value.trim() === '' ? 'Username is required' : '';
  passwordError.value = password.value.trim() === '' ? 'Password is required' : '';
  return usernameError.value === '' && passwordError.value === '';
};

const handleSubmit = async () => {
  if (!validateForm()) return;

  try {
    if (isRegistering.value) {
      await authStore.register(username.value, password.value);
      successMessage.value = 'Registration successful. Please login.';
      toggleForm();
    } else {
      await authStore.login(username.value, password.value);
      router.push('/');
    }
  } catch (error) {
    errorMessage.value = 'An error occurred. Please try again.';
  }
};
</script>
