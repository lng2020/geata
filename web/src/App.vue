<template>
  <div class="flex h-screen">
    <Sidebar v-if="$route.meta.requireAuth" />
    <main class="flex-1 p-4">
      <div class="flex justify-end mb-4">
        <UserInfo v-if="$route.meta.requireAuth" />
      </div>
      <router-view></router-view>
    </main>
  </div>
</template>

<script lang="ts" setup>
import Sidebar from '@/components/Sidebar.vue'
import UserInfo from '@/components/UserInfo.vue'
</script>
