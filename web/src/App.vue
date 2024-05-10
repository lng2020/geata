<template>
  <div class="flex h-screen">
    <template v-if = "isAuthorized">
      <Sidebar />
    </template>
    <main class="flex-1 p-4">
      <router-view></router-view>
    </main>
  </div>
</template>

<script lang="ts" setup>
import Sidebar from '@/components/Sidebar.vue'
import { useAuthStore } from '@/stores/store'
const isAuthorized = useAuthStore().user
</script>
