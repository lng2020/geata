import { createI18n } from 'vue-i18n'

const enUS = {
  dashboar: 'Dashboard',
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
  views: {
    dashboard: {
      title: '仪表盘',
      overview: '概览',
      settings: '设置'
    }
  }
}

type MessageSchema = typeof enUS

const i18n = createI18n<[MessageSchema], 'en-US', 'zh-CN'>({
  locale: 'en-US',
  messages: {
    'en-US': enUS,
    'zh-CN': zhCN
  }
})

export default i18n
