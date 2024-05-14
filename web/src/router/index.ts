import Dashboard from '@/views/Dashboard.vue'
import Station from '@/views/Station.vue'
import Mapping from '@/views/Mapping.vue'
import Setting from '@/views/Setting.vue'
import StationCreate from '@/views/StationCreate.vue'
import Login from '@/views/Login.vue'
import Unauthorized from '@/views/Unauthorized.vue'
import Management from '@/views/Management.vue'
import User from '@/views/User.vue'
import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/store'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: Login,
      meta: {
        hideSidebar: true
      }
    },
    {
      path: '/',
      name: 'dashboard',
      component: Dashboard,
    },
    {
      path: '/station/:id',
      name: 'station',
      component: Station,
    },
    {
      path: '/station/create',
      component: StationCreate,
    },
    {
      path: '/mapping/:id',
      name: 'mapping',
      component: Mapping,
    },
    {
      path: '/setting/:id',
      name: 'setting',
      component: Setting,
    },
    {
      path: '/management',
      name: 'management',
      component: Management,
    },
    {
      path: '/user',
      name: 'user',
      component: User,
    },
    {
      path: '/unauthorized',
      name: 'unauthorized',
      component: Unauthorized
    }
  ]
})

router.beforeEach(async (to, from, next) => {
  const authStore = useAuthStore()

  if (to.name === 'Login') {
    if (authStore.user) {
      next({ name: '/' })
    } else {
      next()
    }
  } else {
    const requiresAuth = to.matched.some(record => record.meta.requiresAuth)

    if (requiresAuth && !authStore.user) {
      next({ name: 'Login' })
    } else {
      next()
    }
  }
})

export default router
