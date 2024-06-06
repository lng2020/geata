export interface Station {
  id: number
  name: string
  host: string
  port: number
  modelId: number
  isOnline: boolean
  lastOnlineTime: string
  showOptions: boolean
}

export type MappingType = 'MQTT' | 'Modbus' | 'IEC61850'
export type RoleType = 'admin' | 'user'

export interface Mapping {
  id: number
  iec61850Ref: string
  type: MappingType
}

export interface DataAttribute {
  name: string
  iec61850Ref: string
  value: string
  dataSource: string
}

export interface DataObject {
  name: string
  description: string
  dataAttribute: DataAttribute[]
}

export interface LogicalNode {
  id: number
  name: string
  description: string
}

export interface LogicalDevice {
  name: string
  description: string
  logicalNode: LogicalNode[]
}

export interface IEC61850Model {
  id: number
  name: string
  description: string
  logicalDevice: LogicalDevice[]
}

export interface AuditLog {
  id: number
  message: string
  createAt: string
}

export interface User {
  id: number
  username: string
  role: RoleType
  lang: string
  avatar: string
}

export interface LoginResponse {
  token: string
  user: User
}

export interface DataSource {
  name: string
  status: string
}

export interface IEC61850Config {
  host: string
  port: number
}

export interface ModbusConfig {
  url: string
}

export interface MQTTConfig {
  broker: string
  clientId: string
  username: string
  password: string
  topic: string
}

export interface ModbusDetail {
  startAddress: number
  dataType: string 
}

export interface MQTTDetail {
  topic: string
  qos: string
}