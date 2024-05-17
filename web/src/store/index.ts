import { defineStore } from 'pinia'
import { useLocalStorage } from '@vueuse/core'
import type { Station, Mapping, IEDModel, LogicalNode, User, LoginResponse } from '@/types/types'
import i18n from '@/i18n/index'

export const useAuthStore = defineStore({
  id: 'auth',
  state: () => ({
    token: useLocalStorage('token', ''),
    user: useLocalStorage('user', null) as User | null
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
    async setLang(lang: string) {
      if (this.user) {
        this.user.lang = lang
      }

      const response = await fetch('/api/user', {
        method: 'PATCH',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${this.token}`
        },
        body: JSON.stringify({ lang })
      })
      if (response.ok) {
        i18n.global.locale.value = lang
      }
    }
  }
})

export const userGlobalStore = defineStore({
  id: 'global',
  state: () => ({
    stations: [
      {
        id: 1,
        name: 'Demo Station 1',
        host: 'localhost',
        port: 8080,
        isOnline: true,
        lastOnlineTime: '2023-06-08T10:00:00Z',
        createdAt: '2023-06-01T00:00:00Z',
        updatedAt: '2023-06-08T10:00:00Z',
        showOptions: false
      },
      {
        id: 2,
        name: 'Demo Station 2',
        host: 'localhost',
        port: 8081,
        isOnline: false,
        lastOnlineTime: '2023-06-07T15:30:00Z',
        createdAt: '2023-06-02T00:00:00Z',
        updatedAt: '2023-06-07T15:30:00Z',
        showOptions: false
      }
    ] as Station[],
    mappings: [
      {
        id: 1,
        iec61850Reference: 'LD0/LLN0$ST$stVal',
        type: 'MQTT'
      },
      {
        id: 2,
        iec61850Reference: 'LD0/LLN0$ST$stVal',
        type: 'Modbus'
      },
      {
        id: 3,
        iec61850Reference: 'LD0/LLN0$ST$stVal',
        type: 'IEC61850'
      },
      {
        id: 4,
        iec61850Reference: 'LD0/LLN0$ST$stVal',
        type: 'IEC61850'
      },
      {
        id: 5,
        iec61850Reference: 'LD0/LLN0$ST$stVal',
        type: 'Modbus'
      },
      {
        id: 6,
        iec61850Reference: 'LD0/LLN0$ST$stVal',
        type: 'IEC61850'
      },
      {
        id: 7,
        iec61850Reference: 'LD0/LLN0$ST$stVal',
        type: 'Modbus'
      },
      {
        id: 8,
        iec61850Reference: 'LD0/LLN0$ST$stVal',
        type: 'IEC61850'
      }
    ] as Mapping[],
    model: [
      {
        name: 'IED1',
        logicalDevices: [
          {
            name: 'LD1',
            logicalNodes: [{ name: 'LN1' }, { name: 'LN2' }]
          }
        ]
      }
    ] as IEDModel[],
    logicalNodes: [
      {
        name: 'LN1',
        dataObjects: [
          {
            name: 'DO1',
            dataAttributes: [
              { name: 'DA1', ref: 'LD0/LLN0$ST$stVal', value: 'Value 1', dataSource: 'MQTT' },
              { name: 'DA2', ref: 'LD0/LLN0$ST$stVal', value: 'Value 2', dataSource: 'MQTT' },
              { name: 'DA3', ref: 'LD0/LLN0$ST$stVal', value: 'Value 3', dataSource: 'IEC61850' },
              { name: 'DA4', ref: 'LD0/LLN0$ST$stVal', value: 'Value 4', dataSource: 'Modbus' },
              { name: 'DA5', ref: 'LD0/LLN0$ST$stVal', value: 'Value 5', dataSource: 'IEC61850' },
              { name: 'DA6', ref: 'LD0/LLN0$ST$stVal', value: 'Value 6', dataSource: 'Modbus' },
              { name: 'DA7', ref: 'LD0/LLN0$ST$stVal', value: 'Value 7', dataSource: 'IEC61850' },
              { name: 'DA8', ref: 'LD0/LLN0$ST$stVal', value: 'Value 8', dataSource: 'Modbus' }
            ]
          },
          {
            name: 'DO2',
            dataAttributes: [
              { name: 'DA3', ref: 'LD0/LLN0$ST$stVal', value: 'Value 3', dataSource: 'Modbus' },
              { name: 'DA4', ref: 'LD0/LLN0$ST$stVal', value: 'Value 4', dataSource: 'IEC61850' },
              { name: 'DA5', ref: 'LD0/LLN0$ST$stVal', value: 'Value 5', dataSource: 'Modbus' },
              { name: 'DA6', ref: 'LD0/LLN0$ST$stVal', value: 'Value 6', dataSource: 'IEC61850' },
              { name: 'DA7', ref: 'LD0/LLN0$ST$stVal', value: 'Value 7', dataSource: 'Modbus' },
              { name: 'DA8', ref: 'LD0/LLN0$ST$stVal', value: 'Value 8', dataSource: 'IEC61850' }
            ]
          }
        ]
      },
      {
        name: 'LN2',
        dataObjects: [
          {
            name: 'DO3',
            dataAttributes: [
              { name: 'DA4', ref: 'LD0/LLN0$ST$stVal', value: 'Value 4', dataSource: 'IEC 61850' }
            ]
          }
        ]
      }
    ] as LogicalNode[],
    auditLogs: [
      {
        id: 1,
        message: 'Station 1 is offline',
        createAt: '2023-06-08T10:00:00Z'
      },
      {
        id: 2,
        message: 'Station 2 is online',
        createAt: '2023-06-08T10:00:00Z'
      },
      {
        id: 3,
        message: 'Station 1 is online',
        createAt: '2023-06-08T10:00:00Z'
      }
    ]
  }),
  actions: {
    async fetchStations() {
      const response = await fetch('/api/stations')
      const data = await response.json()
      this.stations = data
    },
    getStationByID(ID: number): Station | undefined {
      return this.stations.find((station) => station.id === ID)
    },
    async fetchMappings() {
      const response = await fetch('/api/mappings')
      const data = await response.json()
      this.mappings = data
    }
  }
})
