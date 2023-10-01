<template>
    <div class="w-64 bg-gray-200">
        <ul>
            <li class="mb-2">
                <router-link to="/">Dashboard</router-link>
            </li>
            <li v-for="(station, index) in stations" :key="index" @click="toggleOptions(index)"
                class="p-4 cursor-pointer hover:bg-gray-300">
                {{ station.name }}
                <ul v-show="station.showOptions" class="ml-4">
                    <li>
                        <router-link :to="{ path: '/station/' + station.id }">Overview</router-link>
                    </li>
                    <li>
                        <router-link :to="{ path: '/mapping/' + station.id }">mapping</router-link>
                    </li>
                </ul>
            </li>
        </ul>
    </div>
</template>
  
<script lang="ts" setup>
import { userGlobalStore } from '@/stores/store';
const store = userGlobalStore();
const stations = store.stations;
const toggleOptions = (index: number) => {
    stations[index].showOptions = !stations[index].showOptions;
}
</script>

<style scoped>

</style>
  