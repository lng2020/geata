<template>
  <div class="user-info flex items-center cursor-pointer" @click="navigateToUserPage">
    <img :src="avatar" alt="User Avatar" class="avatar w-10 h-10 rounded-full mr-2" />
    <span class="name font-bold mr-2">{{ name }}</span>
    <span v-if="isAdmin" class="admin-badge bg-red-500 text-white px-2 py-1 rounded text-xs"
      >Admin</span
    >
  </div>
</template>

<script lang="ts" setup>
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/store'

const router = useRouter()
const authStore = useAuthStore()
const user = authStore.user
const avatar = user?.avatar
const name = user?.username
const isAdmin = user?.role === 'admin'
const navigateToUserPage = () => {
  router.push('/user')
}
</script>
