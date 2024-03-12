<template>
  <div class="flex h-screen">
    <Sidebar />
    <main class="flex-1 p-4">
      <router-view></router-view>
    </main>
  </div>
</template>

<script lang="ts">
import Sidebar from '@/components/Sidebar.vue'

export default {
  name: 'App',
  components: {
    Sidebar
  }
}
</script>

<style scoped></style>
