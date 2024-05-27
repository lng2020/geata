import { createI18n } from 'vue-i18n'

const enUS = {
  dashboard: 'Dashboard',
  overview: 'Overview',
  settings: 'Settings',
  onlineStatus: 'ONLINE STATUS',
  searchStationPlaceholder: 'Search stations by name...',
  searchStation: 'Search Station',
  createStation: 'Create Station',
  name: 'NAME',
  isOnline: 'IS ONLINE',
  lastOnlineTime: 'LAST ONLINE TIME',
  address: 'ADDRESS',
  online: 'online',
  offline: 'offline',
  host: 'host',
  port: 'port',
  day1: 'day1',
  day2: 'day2',
  day3: 'day3',
  day4: 'day4',
  day5: 'day5',
  day6: 'day6',
  day7: 'day7',
  onlineStationNumber: 'Online Station Number',
  login: 'Login',
  register: 'Register',
  username: 'Username',
  password: 'Password',
  usernameRequired: 'Username is required',
  passwordRequired: 'Password is required',
  registrationSuccess: 'Registration successful. Please login.',
  loginError: 'An error occurred. Please try again.',
  unauthorizedTitle: 'Unauthorized',
  unauthorizedMessage: 'You do not have permission to access this page.',
  goBackToDashboard: 'Go Back to Dashboard',
  home: 'Home',
  profile: 'Profile',
  settings: 'Settings',
  mapping: 'Mapping',
  userAvatar: "User Avatar",
  settings: "Settings",
  admin: "Admin",
  logout: "Logout",
  title: "Confirmation",
  message: "Are you sure you want to proceed?",
  confirm: "Confirm",
  save: "Save",
  cancel: "Cancel",
  modbusConfigTitle: "Modbus Configuration",
  addressLabel: "Address:",
  dataTypeLabel: "Data Type:",
  int16: "INT16",
  int32: "INT32",
  float: "FLOAT",
  dataSourceMappingTitle: "Data Source Mapping",
  batchEdit: "Batch Edit",
  importExport: "Import/Export",
  utils: "Utils",
  logicalNode: "Logical Node",
  dataSourceType: "Data Source Type",
  configuration: "Configuration",
  edit: "Edit",
  iec61850: "IEC61850",
  modbus: "Modbus",
  mqtt: "MQTT",
  confirmation: "Confirmation",
  areYouSure: "Are you sure you want to select IEC61850?",
  mqttConfigTitle: "MQTT Configuration",
  topicLabel: "Topic:",
  qosLabel: "QoS:",
  iec61850ExplorerTitle: "IEC61850 Explorer",
  iec61850Model: "IEC61850 Model",
  searchPlaceholder: "Search...",
  logicalNodeDetails: "Logical Node Details",
  export: "Export",
  refresh: "Refresh",
  selectNode: "Select a Logical Node to view its details",
  doName: "DO Name",
  daName: "DA Name",
  reference: "Reference",
  value: "Value",
  dataSource: "Data Source"
}

const zhCN = {
  dashboard: '仪表盘',
  overview: '总览',
  settings: '设置',
  onlineStatus: '在线状态',
  searchStationPlaceholder: '按名称搜索站点...',
  searchStation: '搜索站点',
  createStation: '创建站点',
  name: '名称',
  isOnline: '在线状态',
  lastOnlineTime: '最后在线时间',
  address: '地址',
  online: '在线',
  offline: '离线',
  host: '主机',
  port: '端口',
  day1: '第一天',
  day2: '第二天',
  day3: '第三天',
  day4: '第四天',
  day5: '第五天',
  day6: '第六天',
  day7: '第七天',
  onlineStationNumber: '在线站点数量',
  login: '登录',
  register: '注册',
  username: '用户名',
  password: '密码',
  usernameRequired: '用户名是必须的',
  passwordRequired: '密码是必须的',
  registrationSuccess: '注册成功，请登录。',
  loginError: '发生错误，请重试。',
  unauthorizedTitle: '未授权',
  unauthorizedMessage: '您没有权限访问此页面。',
  goBackToDashboard: '返回仪表板',
  home: '首页',
  profile: '个人资料',
  settings: '设置',
  mapping: '映射',
  userAvatar: "用户头像",
  settings: "设置",
  admin: "管理员",
  logout: "登出",
  title: "确认",
  message: "您确定要继续吗？",
  confirm: "确认",
  cancel: "取消",
  save: "保存",
  modbusConfig: "Modbus配置",
  addressLabel: "地址:",
  dataTypeLabel: "数据类型:",
  int16: "整型16",
  int32: "整型32",
  float: "浮点型",
  dataSourceMappingTitle: "数据源映射",
  batchEdit: "批量编辑",
  importExport: "导入/导出",
  utils: "工具",
  logicalNode: "逻辑节点",
  dataSourceType: "数据源类型",
  configuration: "配置",
  edit: "编辑",
  iec61850: "IEC61850",
  modbus: "Modbus",
  mqtt: "MQTT",
  confirmation: "确认",
  areYouSure: "您确定要选择IEC61850吗？",
  mqttConfigTitle: "MQTT配置",
  topicLabel: "主题：",
  qosLabel: "服务质量：",
  iec61850ExplorerTitle: "IEC61850 浏览器",
  iec61850Model: "IEC61850 模型",
  searchPlaceholder: "搜索...",
  logicalNodeDetails: "逻辑节点详情",
  export: "导出",
  refresh: "刷新",
  selectNode: "选择逻辑节点以查看其详细信息",
  doName: "DO名称",
  daName: "DA名称",
  reference: "IEC61850引用",
  value: "值",
  dataSource: "数据源"
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
