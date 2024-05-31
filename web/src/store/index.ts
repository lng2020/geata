import { defineStore } from 'pinia'
import { useLocalStorage } from '@vueuse/core'
import type {
  Station,
  Mapping,
  IEC61850Model,
  LogicalNode,
  User,
  LoginResponse
} from '@/types/types'
import i18n from '@/i18n/index'

export const useAuthStore = defineStore({
  id: 'auth',
  state: () => ({
    token: useLocalStorage('token', ''),
    user: useLocalStorage<User | null>('user', null, {
      serializer: {
        read: (v: any) => (v ? JSON.parse(v) : null),
        write: (v: any) => JSON.stringify(v)
      }
    })
  }),
  actions: {
    async login(username: string, password: string) {
      const credentials = { username, password }
      const response = await fetch('/api/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(credentials)
      })
      if (!response.ok) {
        return
      }
      const data: LoginResponse = await response.json()
      this.token = data.token
      this.user = data.user
      localStorage.setItem('token', this.token)
      localStorage.setItem('user', JSON.stringify(this.user))
    },
    logout() {
      this.token = ''
      this.user = null
      localStorage.removeItem('token')
      localStorage.removeItem('user')
    },
    async register(username: string, password: string) {
      const credentials = { username, password }
      const response = await fetch('/api/register', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(credentials)
      })
      if (!response.ok) {
        return
      }
      const data: LoginResponse = await response.json()
      this.token = data.token
      this.user = data.user
      localStorage.setItem('token', this.token)
      localStorage.setItem('user', JSON.stringify(this.user))
    },
    async changeLang(lang: string) {
      const response = await fetch('/api/user/lang', {
        method: 'PATCH',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${this.token}`
        },
        body: JSON.stringify({
          lang: lang
        })
      })
      if (response.ok) {
        i18n.global.locale.value = lang
        this.user = { ...this.user, lang }
      }
    },
    async changeTheme(theme: string) {
      console.log('theme', theme)
    }
  }
})
export const useGlobalStore = defineStore({
  id: 'global',
  state: () => ({
    stations: [] as Station[],
    mappings: [] as Mapping[],
    models: [] as IEC61850Model[],
    logicalNodes: [] as LogicalNode[],
    auditLogs: [] as AuditLog[]
  }),
  actions: {
    async fetchStations() {
      const response = await fetch('/api/station')
      const data = await response.json()
      const respStation: Station[] = data.map((item: any) => ({
        id: item.id,
        name: item.name,
        host: item.host,
        port: item.port,
        modelHash: item.modelHash,
        isOnline: item.isOnline,
        lastOnlineTime: new Date(item.lastOnlineTime)
      }))
      this.stations = respStation
    },
    getStationById(id: number): Station {
      return this.stations.find((station) => station.id === id)
    },
    async fetchMappings() {
      const response = await fetch('/api/mappings')
      const data = await response.json()
      this.mappings = data
    }
  }
})
