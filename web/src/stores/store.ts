import { defineStore } from 'pinia'
import type { Station, Mapping, IEDModel, LogicalNode } from '@/types/types'

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
              { name: 'DA2', ref: 'LD0/LLN0$ST$stVal', value: 'Value 2', dataSource: 'MQTT' }
            ]
          },
          {
            name: 'DO2',
            dataAttributes: [
              { name: 'DA3', ref: 'LD0/LLN0$ST$stVal', value: 'Value 3', dataSource: 'Modbus' }
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
    ] as LogicalNode[]
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
