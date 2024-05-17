import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { fas } from '@fortawesome/free-solid-svg-icons'
import './index.css'

import App from './App.vue'
import router from './router'
import i18n from './i18n'
import { useAuthStore } from './store'

library.add(fas)

const pinia = createPinia()
const app = createApp(App)

app.use(pinia)

const authStore = useAuthStore()
if (authStore.user?.lang) {
  i18n.global.locale.value = authStore.user.lang
}

app.use(router).use(i18n).component('font-awesome-icon', FontAwesomeIcon).mount('#app')
