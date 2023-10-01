<template>
    <div class=" h-screen bg-indigo-950" :class="sidebarVisible ? 'w-72' : 'w-15'">
        <div class="mb-2 relative flex items-center">
            <img src="../assets/gopher.png" alt="logo" class="w-10 h-10 m-4 duration-300">
            <span class="text-white text-2xl font-extrabold" :class="{ 'hidden': !sidebarVisible }">Geata</span>
            <button @click="sidebarVisible = !sidebarVisible">
                <font-awesome-icon icon="fa-solid fa-arrow-left"
                    class=" bg-white rounded-full text-3xl absolute -right-3 top-9 border border-indigo-700 cursor-pointer"
                    :class="{ 'rotate-180': !sidebarVisible }"></font-awesome-icon>
            </button>
        </div>
        <ul>
            <li class="mb-4 flex items-center">
                <font-awesome-icon icon="fa-solid fa-home" class="text-white text-2xl ml-6 mr-4"></font-awesome-icon>
                <router-link to="/" class="text-white text-xl ml-2"
                    :class="{ 'hidden': !sidebarVisible }">Dashboard</router-link>
            </li>
            <li v-for="(station, index) in stations" :key="index" @click="toggleOptions(index)">
                <div class="mb-4 flex items-center">
                    <font-awesome-icon icon="fa-solid fa-charging-station" class=" text-white text-2xl ml-6 mr-4">
                    </font-awesome-icon>
                    <div :class="{ 'hidden': !sidebarVisible }">
                        <span class="text-white text-xl ml-2">{{ station.name }}</span>
                        <font-awesome-icon icon="fa-solid fa-chevron-down" class="text-white text -right-4"
                            :class="{ 'rotate-180': !station.showOptions }"></font-awesome-icon>
                    </div>
                </div>
                <div :class="{ 'hidden': !sidebarVisible || !station.showOptions }">
                    <ul v-show="station.showOptions">
                        <li class="text-white text ml-32 mb-4">
                            <router-link :to="{ path: '/station/' + station.id }">Overview</router-link>
                        </li>
                        <li class="text-white text ml-32 mb-4">
                            <router-link :to="{ path: '/mapping/' + station.id }">mapping</router-link>
                        </li>
                    </ul>
                </div>
            </li>
        </ul>
    </div>
</template>
  
<script lang="ts" setup>
import { userGlobalStore } from '@/stores/store';
import { ref, reactive } from 'vue';
const store = userGlobalStore();
let stations = reactive(store.stations);
let sidebarVisible = ref(store.sidebarVisible);
const toggleOptions = (index: number) => {
    stations[index].showOptions = !stations[index].showOptions;
}
</script>

<style scoped></style>
  