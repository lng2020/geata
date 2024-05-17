import { createI18n } from 'vue-i18n'

const enUS = {
  dashboard: 'Dashboard',
  overview: 'Overview',
  settings: 'Settings',
  onlineStatus: 'Online Status',
  searchPlaceholder: 'Search stations by name...',
  searchButton: 'Search',
  name: 'NAME',
  isOnline: 'IS ONLINE',
  lastOnlineTime: 'LAST ONLINE TIME',
  address: 'ADDRESS',
  online: 'online',
  offline: 'offline'
}

const zhCN = {
  dashboard: '仪表盘',
  overview: '总览',
  settings: '设置',
  onlineStatus: '在线状态',
  searchPlaceholder: '按名称搜索站点...',
  searchButton: '搜索',
  name: '名称',
  isOnline: '在线状态',
  lastOnlineTime: '最后在线时间',
  address: '地址',
  online: '在线',
  offline: '离线'
}

type MessageSchema = typeof enUS

const i18n = createI18n<[MessageSchema], 'en-US', 'zh-CN'>({
  legacy: false,
  locale: 'en-US',
  messages: {
    'en-US': enUS,
    'zh-CN': zhCN
  }
})

export default i18n
