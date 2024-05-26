export interface Station {
  id: number
  name: string
  host: string
  port: number
  isOnline: boolean
  lastOnlineTime: string
  createdAt: string
  updatedAt: string
  showOptions: boolean
}

export type MappingType = 'MQTT' | 'Modbus' | 'IEC61850'
export type RoleType = 'admin' | 'user'

export interface Mapping {
  id: number
  iec61850Reference: string
  type: MappingType
}

export interface DataAttribute {
  name: string
  ref: string
  value: string
  dataSource: string
}

export interface DataObject {
  name: string
  description: string
  dataAttributes: DataAttribute[]
}

export interface LogicalNode {
  name: string
  description: string
}

export interface LogicalDevice {
  name: string
  description: string
  logicalNode: LogicalNode[]
}

export interface IEC61850Model {
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
