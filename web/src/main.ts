import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { fas } from '@fortawesome/free-solid-svg-icons'
import './index.css'

import App from './App.vue'
import router from './router'

library.add(fas)

createApp(App)
.use(createPinia())
.use(router)
.component('font-awesome-icon', FontAwesomeIcon)
.mount('#app')
