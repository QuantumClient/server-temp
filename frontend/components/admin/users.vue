<template>
  <div class="bg-white p-4 bg-accent shadow-md rounded-2xl " >
    <table class="min-w-full table-auto overflow-x-scroll ">
      <thead class="bg-gray-300 border-b">
      <tr>
        <th class="px-4 py-2">Username</th>
        <th class="px-4 py-2">Access</th>
        <th class="px-4 py-2">Admin</th>
        <th class="px-4 py-2">Date Joined</th>
        <th class="px-4 py-2">HWID</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(user, index) in users" :key="index" class="border-b ">
        <td class=" px-4 py-2">{{ user.username }}</td>
        <td class=" px-4 py-2">
          <button @click="setAccess(user)">
          <div v-if="user.access" >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="w-5 h5 "
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="#ffffff"
              fill="none"
              stroke-linecap="round"
              stroke-linejoin="round"
            >
              <path stroke="none" d="M0 0h24v24H0z" />
              <path d="M5 12l5 5l10 -10" />
            </svg>
          </div>
          <div v-else>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="w-5 h-5 "
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="#ffffff"
              fill="none"
              stroke-linecap="round"
              stroke-linejoin="round"
            >
              <path stroke="none" d="M0 0h24v24H0z" />
              <line x1="18" y1="6" x2="6" y2="18" />
              <line x1="6" y1="6" x2="18" y2="18" />
            </svg>
          </div>
          </button>
        </td>
        <td class=" px-4 py-2">
          <button @click="setAdmin(user)">
          <div v-if="user.admin" >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="w-5 h5 "
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="#ffffff"
              fill="none"
              stroke-linecap="round"
              stroke-linejoin="round"
            >
              <path stroke="none" d="M0 0h24v24H0z" />
              <path d="M5 12l5 5l10 -10" />
            </svg>
          </div>
          <div v-else>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="w-5 h-5 "
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="#ffffff"
              fill="none"
              stroke-linecap="round"
              stroke-linejoin="round"
            >
              <path stroke="none" d="M0 0h24v24H0z" />
              <line x1="18" y1="6" x2="6" y2="18" />
              <line x1="6" y1="6" x2="18" y2="18" />
            </svg>
          </div>
          </button>
        </td>
        <td class=" px-4 py-2">{{ user.created_at | moment }}</td>
        <td class=" px-4 py-2">
          <button @click="resetHwid(user)">
            <svg
              class="w-5 h-5 "
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg">
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M9 3v2m6-2v2M9 19v2m6-2v2M5 9H3m2 6H3m18-6h-2m2 6h-2M7 19h10a2 2 0 002-2V7a2 2 0 00-2-2H7a2 2 0 00-2 2v10a2 2 0 002 2zM9 9h6v6H9V9z"
              ></path></svg>
          </button>
        </td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
import moment from 'moment'
import Vue from 'vue'

interface User {
  username: string;
  uuid: string;
  access: boolean;
  admin: boolean;
}

export default Vue.extend({
  name: "users",
  data() {
    return {
      users: [] as User[]
    }
  },
  filters: {
    moment: function(date: string) {
      return moment(date).format('MMMM Do YYYY')
    }
  },
  methods: {
    async setAdmin(user: User) {
      try {
        await this.$axios.post(`/api/v1/auth/users/${user.uuid}/admin`)
          .then(() => user.admin = !user.admin).then(() =>
            this.$toast.info(`admin for ${user.username} has been set to ${user.admin}`, {
            duration : 500
        }));
      } catch (e) {
        this.$toast.error('Error', {
          position: "top-center",
          duration : 500
        });
      }
    },
    async resetHwid(user: User) {
      try {
        await this.$axios.post(`/api/v1/auth/users/${user.uuid}/hwid`)
          .then(() => this.$toast.info(`Hwid for ${user.username} has been reset`, {
            duration : 500
          }));
      } catch (e) {
        this.$toast.error('Error', {
          position: "top-center",
          duration : 500
        });
      }
    },
    async setAccess(user: User) {
      try {
        await this.$axios.post(`/api/v1/auth/users/${user.uuid}/access`)
          .then(() => user.access = !user.access).then(() =>
            this.$toast.info(`Access for ${user.username} has been set to ${user.access}`, {
              duration : 500
            }));
      } catch (e) {
        this.$toast.error('Error', {
          position: "top-center",
          duration : 500
        });
      }
    }
  },
  async fetch() {
    this.users = (await this.$axios.get('/api/v1/auth/users')).data
  }
});
</script>

<style scoped>

</style>
