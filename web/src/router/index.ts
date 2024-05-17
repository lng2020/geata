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
import { useAuthStore } from '@/store'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: Login,
    },
    {
      path: '/',
      name: 'dashboard',
      component: Dashboard,
      meta: {
        requireAuth: true
      }
    },
    {
      path: '/station/:id',
      name: 'station',
      component: Station,
      meta: {
        requireAuth: true
      }
    },
    {
      path: '/station/create',
      component: StationCreate,
      meta: {
        requireAuth: true
      }
    },
    {
      path: '/mapping/:id',
      name: 'mapping',
      component: Mapping,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/setting/:id',
      name: 'setting',
      component: Setting,
      meta: {
        requireAuth: true
      }
    },
    {
      path: '/management',
      name: 'management',
      component: Management,
      meta: {
        requireAuth: true
      }
    },
    {
      path: '/user',
      name: 'user',
      component: User,
      meta: {
        requireAuth: true
      }
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
  console.log(to.name)
  if (to.name === 'login') {
    if (authStore.token) {
      next({ name: 'dashboard' })
    } else {
      next()
    }
  } else {
    const requiresAuth = to.matched.some((record) => record.meta.requireAuth)
    if (requiresAuth && !authStore.token) {
      next({ name: 'login' })
    } else {
      next()
    }
  }
})

export default router
