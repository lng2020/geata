
<template>
  <div class="container mx-auto px-2 py-2">
    <div class="bg-white border border-gray-200 rounded-md p-6 mb-6">
    <h2 class="text-2xl font-bold mb-6">{{ $t('userManagement') }}</h2>
      <div class="flex justify-between items-center mb-6 bg-white">
      <div class="flex space-x-2">
        <input
          type="text"
          v-model="searchQuery"
          :placeholder="$t('searchUsers')"
          class="border border-gray-300 rounded-l-md px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <button
          @click="searchUsers"
          class="bg-blue-500 text-white rounded-r-md px-4 py-2 hover:bg-blue-600"
        >
          {{ $t('search') }}
        </button>
      </div>
      <div class="flex space-x-2">
        <button
          @click="openCreateUserDialog"
          class="bg-green-500 text-white rounded-md px-4 py-2 hover:bg-green-600"
        >
          {{ $t('newUser') }}
        </button>
        <button
          @click="toggleSSO"
          :class="[
            'rounded-md px-4 py-2',
            ssoEnabled
              ? 'bg-blue-500 text-white hover:bg-blue-600'
              : 'bg-gray-200 text-gray-700 hover:bg-gray-300'
          ]"
        >
          {{ ssoEnabled ? $t('ssoEnabled') : $t('ssoDisabled') }}
        </button>
      </div>
    </div>
    <div class="overflow-x-auto">
      <table class="w-full">
        <thead>
          <tr class="bg-gray-100 text-gray-600 uppercase text-sm leading-normal">
            <th class="py-3 px-6 text-left">{{ $t('username') }}</th>
            <th class="py-3 px-6 text-left">{{ $t('role') }}</th>
            <th class="py-3 px-6 text-left">{{ $t('status') }}</th>
            <th class="py-3 px-6 text-left">{{ $t('Actions') }}</th>
          </tr>
        </thead>
        <tbody class="text-gray-600 text-sm">
          <tr v-for="user in users" :key="user.id" class="border-b border-gray-200 hover:bg-gray-100">
            <td class="py-4 px-6">{{ user.username }}</td>
            <td class="py-4 px-6">{{ user.role }}</td>
            <td class="py-4 px-6">
              <span
                class="inline-block rounded-full px-3 py-1 text-sm font-semibold"
                :class="user.status === 'Active' ? 'bg-green-500 text-white' : 'bg-red-500 text-white'"
              >
                {{ user.status }}
              </span>
            </td>
            <td class="py-4 px-6 space-x-1">
              <button
                @click="openEditUserDialog(user)"
                class="bg-blue-500 text-white rounded px-2 py-1 hover:bg-blue-600"
              >
                Edit
              </button>
              <button
                @click="deleteUser(user)"
                class="bg-red-500 text-white rounded px-2 py-1 hover:bg-red-600"
              >
                Delete
              </button>
              <button
                @click="toggleUserStatus(user)"
                :class="[
                  'px-2 py-1 rounded',
                  user.status === 'Active'
                    ? 'bg-green-500 text-white hover:bg-green-600'
                    : 'bg-yellow-500 text-white hover:bg-yellow-600'
                ]"
              >
                {{ user.status === 'Active' ? 'Disable' : 'Enable' }}
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    </div>

    <div class="mt-6 bg-white border border-gray-200 rounded-md p-6">
      <h2 class="text-2xl font-bold mb-4">{{ $t('loginLogs') }}</h2>
      <div class="flex justify-between items-center mb-4">
        <div class="flex space-x-4">
          <label for="startDate" class="flex items-center">
            <span class="mr-2">{{ $t('startDate') }}:</span>
            <input
              type="date"
              id="startDate"
              v-model="startDate"
              class="border border-gray-300 rounded-md px-2 py-1"
            />
          </label>
          <label for="endDate" class="flex items-center">
            <span class="mr-2">{{ $t('endDate') }}:</span>
            <input
              type="date"
              id="endDate"
              v-model="endDate"
              class="border border-gray-300 rounded-md px-2 py-1"
            />
          </label>
        </div>
        <button
          @click="filterLogs"
          class="bg-blue-500 text-white rounded-md px-4 py-2 hover:bg-blue-600"
        >
          {{ $t('filter') }}
        </button>
      </div>
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead>
            <tr class="bg-gray-100 text-gray-600 uppercase text-sm leading-normal">
              <th class="py-3 px-6 text-left">{{ $t('date') }}</th>
              <th class="py-3 px-6 text-left">{{ $t('user') }}</th>
              <th class="py-3 px-6 text-left">{{ $t('action') }}</th>
            </tr>
          </thead>
          <tbody class="text-gray-600 text-sm">
            <tr v-for="log in logs" :key="log.id" class="border-b border-gray-200 hover:bg-gray-100">
              <td class="py-4 px-6">{{ log.date }}</td>
              <td class="py-4 px-6">{{ log.user }}</td>
              <td class="py-4 px-6">{{ log.action }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import type { User } from 'src/types/types'
import { onBeforeRouteUpdate } from 'vue-router';

onMounted(() => {
  fetchUsers()
}),

onBeforeRouteUpdate(() => {
  fetchUsers()
})
interface Log {
  id: number
  date: string
  user: string
  action: string
}

const users = ref<User[]>([])
const searchQuery = ref('')
const ssoEnabled = ref(false)
const logs = ref<Log[]>([])
const startDate = ref('')
const endDate = ref('')
logs.value = [
  { id: 1, date: '2024-6-7', user: 'admin', action: 'Login' },
  { id: 2, date: '2024-6-7', user: 'admin', action: 'Logout' },
  { id: 3, date: '2024-6-7', user: 'admin', action: 'Login' },
  { id: 4, date: '2024-6-7', user: 'test_user_1', action: 'Login'},
  { id: 5, date: '2024-6-7', user: 'test_user_1', action: 'Logout'},
  { id: 6, date: '2024-6-7', user: 'test_user_2', action: 'Login'},
  { id: 7, date: '2024-6-7', user: 'test_user_2', action: 'Logout'},
]

const fetchUsers = async () => {
  try {
    const token = localStorage.getItem('token')
    const response = await fetch('/api/management/user', {
      headers: {
        Authorization: `Bearer ${token}`
      }
    })
    users.value = await response.json()
  } catch (error) {
    console.error('Error fetching users:', error)
  }
}

const searchUsers = () => {}

const openCreateUserDialog = () => {}

const toggleSSO = () => {
  ssoEnabled.value = !ssoEnabled.value
}

const openEditUserDialog = (user: User) => {
  console.log('Editing user:', user)
}

const deleteUser = (user: User) => {
  console.log('Deleting user:', user)
}

const toggleUserStatus = (user: User) => {
  user.status = user.status === 'Active' ? 'Inactive' : 'Active'
}

const filterLogs = () => {}
</script>
