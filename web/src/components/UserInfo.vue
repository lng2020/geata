<template>
  <div class="user-info relative">
    <div
      class="flex items-center cursor-pointer bg-white p-3 rounded-lg shadow-md hover:shadow-lg transition-shadow duration-200"
      @click="toggleDropdown"
    >
      <img
        :src="avatar"
        alt="User Avatar"
        class="w-8 h-8 rounded-full mr-3 border border-gray-300"
      />
      <span class="font-medium text-gray-800 hidden sm:block">{{ name }}</span>
      <svg
        class="w-4 h-4 ml-2 fill-current text-gray-600"
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 20 20"
      >
        <path
          d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z"
        />
      </svg>
    </div>
    <div
      v-if="isDropdownOpen"
      class="dropdown absolute right-0 mt-2 w-48 bg-white rounded-lg shadow-lg z-50"
    >
      <a
        href="#"
        class="block px-4 py-2 text-sm text-gray-800 hover:bg-gray-100"
        @click="navigateToUserPage"
        >{{ $t('settings') }}</a
      >
      <a href="#" class="block px-4 py-2 text-sm text-gray-800 hover:bg-gray-100">{{ name }}</a>
      <a href="#" class="block px-4 py-2 text-sm text-gray-800 hover:bg-gray-100" v-if="isAdmin">{{
        $t('admin')
      }}</a>
      <a href="#" class="block px-4 py-2 text-sm text-gray-800 hover:bg-gray-100" @click="logout">{{
        $t('logout')
      }}</a>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/store'

const router = useRouter()
const authStore = useAuthStore()
const user = authStore.user
const avatar = user?.avatar ? user.avatar : '/src/assets/gopher.png'
const name = user?.username
const isAdmin = user?.role === 'admin'
const isDropdownOpen = ref(false)

const toggleDropdown = () => {
  isDropdownOpen.value = !isDropdownOpen.value
}

const navigateToUserPage = () => {
  router.push('/user')
  isDropdownOpen.value = false
}

const logout = () => {
  authStore.logout()
  router.push('/login')
}
</script>
