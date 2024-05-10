import Dashboard from '@/views/Dashboard.vue'
import Station from '@/views/Station.vue'
import Mapping from '@/views/Mapping.vue'
import Setting from '@/views/Setting.vue'
import StationCreate from '@/views/StationCreate.vue'
import Login from '@/views/Login.vue'
import Unauthorized from '@/views/Unauthorized.vue'
import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/store'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/',
      name: 'dashboard',
      component: Dashboard,
      meta: { requiresAuth: true }
    },
    {
      path: '/station/:id',
      name: 'station',
      component: Station,
      meta: { requiresAuth: true }
    },
    {
      path: '/station/create',
      component: StationCreate,
      meta: { requiresAuth: true }
    },
    {
      path: '/mapping/:id',
      name: 'mapping',
      component: Mapping,
      meta: { requiresAuth: true }
    },
    {
      path: '/setting/:id',
      name: 'setting',
      component: Setting,
      meta: { requiresAuth: true }
    },
    {
      path: '/unauthorized',
      name: 'unauthorized',
      component: Unauthorized
    }
  ]
})

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore();
  const requiresAuth = to.matched.some(record => record.meta.requiresAuth);

  if (requiresAuth && !authStore.user) {
    next({ name: 'login' });
  } else if (to.name === 'login' && authStore.user) {
    next({ name: 'dashboard' });
  } else {
    next();
  }
});

export default router