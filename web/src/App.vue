<template>
  <div class="flex h-screen">
    <Sidebar v-if="!$route.meta.hideSidebar" />
    <main class="flex-1 p-4">
      <router-view></router-view>
    </main>
  </div>
</template>

<script lang="ts" setup>
import Sidebar from '@/components/Sidebar.vue'
</script>
