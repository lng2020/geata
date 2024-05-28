<template>
  <div class="h-screen w-64 bg-gray-800">
    <div class="flex items-center px-4 py-4 bg-gray-900">
      <img src="../assets/gopher.png" alt="logo" class="w-8 h-8 mr-3" />
      <span class="text-white text-xl font-bold">Geata</span>
    </div>
    <nav class="mt-6">
      <router-link to="/" class="flex items-center px-5 py-3 text-white hover:bg-gray-700">
        <font-awesome-icon icon="fa-solid fa-home" class="text-lg mr-3"></font-awesome-icon>
        <span>{{ $t('dashboard') }}</span>
      </router-link>
      <div v-for="(station, index) in stations" :key="index" class="mt-4">
        <button
          @click="toggleOptions(index)"
          class="flex items-center justify-between w-full px-5 py-3 text-white hover:bg-gray-700"
        >
          <div class="flex items-center">
            <font-awesome-icon
              icon="fa-solid fa-charging-station"
              class="text-lg mr-3"
            ></font-awesome-icon>
            <span>{{ station.name }}</span>
          </div>
          <font-awesome-icon
            :icon="station.showOptions ? 'fa-solid fa-chevron-up' : 'fa-solid fa-chevron-down'"
            class="text-sm"
          ></font-awesome-icon>
        </button>
        <div
          v-show="station.showOptions"
          class="mt-2 pt-2 pb-2 bg-gray-700 flex flex-col items-center justify-center space-y-2"
        >
          <router-link
            :to="{ path: '/station/' + station.id }"
            class="block text-sm text-gray-400 hover:text-white"
          >
            {{ $t('overview') }}
          </router-link>
          <router-link
            :to="{ path: '/mapping/' + station.id }"
            class="block text-sm text-gray-400 hover:text-white"
          >
            {{ $t('mapping') }}
          </router-link>
          <router-link
            :to="{ path: '/setting/' + station.id }"
            class="block text-sm text-gray-400 hover:text-white"
          >
            {{ $t('settings') }}
          </router-link>
        </div>
      </div>
    </nav>
  </div>
</template>

<script lang="ts" setup>
import { userGlobalStore } from '@/store'
import { reactive } from 'vue'
const store = userGlobalStore()
const stations = reactive(store.stations)
const toggleOptions = (index: number) => {
  stations[index].showOptions = !stations[index].showOptions
}
</script>
