import { defineStore } from 'pinia'

interface Station {
    id: number;
    name: string;
    showOptions: boolean;
}

export const userGlobalStore = defineStore({
    id: 'global',
    state: () => ({
        stations: [
            {
                id: 1,
                name: 'Station 1',
                showOptions: false,
            },
            {
                id: 2,
                name: 'Station 2',
                showOptions: false,
            },
        ] as Station[],
    }),
    actions: {
        async fetchStations() {
            const response = await fetch('/api/stations')
            const data = await response.json()
            this.stations = data
        },
    },
});