<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="max-w-md w-full px-6 py-8 bg-white shadow-md rounded-lg text-center">
      <h2 class="text-2xl font-semibold mb-4">{{ $t('unauthorizedTitle') }}</h2>
      <p class="mb-6">{{ $t('unauthorizedMessage') }}</p>
      <router-link
        to="/"
        class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
      >
        {{ $t('goBackToDashboard') }}
      </router-link>
    </div>
  </div>
</template>
