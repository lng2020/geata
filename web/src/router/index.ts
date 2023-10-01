import Dashboard from '@/views/Dashboard.vue'
import Station from '@/views/Station.vue'
import Mapping from '@/views/Mapping.vue'
import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'dashborad',
      component: Dashboard
    },
    {
      path: '/station/:id',
      name: 'station',
      component: Station
    },
    {
      path: '/mapping/:id',
      name: 'mapping',
      component: Mapping
    }
  ]
})

export default router
