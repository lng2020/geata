<template>
  <div class="container mx-auto px-2 py-2 bg-white shadow rounded-lg">
    <h1 class="text-3xl font-bold mb-6">{{ $t('userManagement') }}</h1>
    <div class="flex justify-between items-center mb-6">
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
      <table class="w-full text-left">
        <thead class="bg-gray-100">
          <tr>
            <th class="px-4 py-3">{{ $t('username') }}</th>
            <th class="px-4 py-3">{{ $t('role') }}</th>
            <th class="px-4 py-3">{{ $t('status') }}</th>
            <th class="px-4 py-3">{{ $t('Actions') }}</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="user in users" :key="user.id" class="border-b">
            <td class="px-4 py-2">{{ user.username }}</td>
            <td class="px-4 py-2">{{ user.role }}</td>
            <td class="px-4 py-2">{{ user.status }}</td>
            <td class="px-4 py-2 space-x-1">
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
    <div class="mt-6">
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
        <table class="w-full text-left">
          <thead class="bg-gray-100">
            <tr>
              <th class="px-4 py-3">{{ $t('date') }}</th>
              <th class="px-4 py-3">{{ $t('user') }}</th>
              <th class="px-4 py-3">{{ $t('action') }}</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="log in logs" :key="log.id" class="border-b">
              <td class="px-4 py-2">{{ log.date }}</td>
              <td class="px-4 py-2">{{ log.user }}</td>
              <td class="px-4 py-2">{{ log.action }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'

interface User {
  id: number
  username: string
  email: string
  role: string
  status: string
}

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

users.value = [
  { id: 1, username: 'john_doe', email: 'john@example.com', role: 'Admin', status: 'Active' },
  { id: 2, username: 'jane_smith', email: 'jane@example.com', role: 'User', status: 'Active' }
]

logs.value = [
  { id: 1, date: '2021-10-01', user: 'john_doe', action: 'Login' },
  { id: 2, date: '2021-10-02', user: 'jane_smith', action: 'Logout' }
]

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
