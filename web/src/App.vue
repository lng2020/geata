<template>
  <div class="flex h-screen">
    <Sidebar v-if="$route.meta.requireAuth" />
    <main class="flex-1 p-4 bg-gray-100">
      <div class="flex justify-between items-center">
        <Breadcrumbs v-if="$route.meta.requireAuth" />
        <UserInfo v-if="$route.meta.requireAuth" />
      </div>
      <div class="mt-4">
        <router-view></router-view>
      </div>
    </main>
  </div>
</template>

<script lang="ts" setup>
import Sidebar from '@/components/Sidebar.vue'
import UserInfo from '@/components/UserInfo.vue'
import Breadcrumbs from '@/components/Breadcrumbs.vue'
</script>
