<template>
  <div class=" p-4 bg-accent shadow-md rounded-2xl min-w-min h-auto max-w-sm w-72 m-8" >
    <table class="table-auto overflow-x-scroll w-full">
      <thead class="bg-gray-300 border-b">
      <tr>
        <th class="px-4 py-2">Owner</th>
        <th class="px-4 py-2">Username</th>
        <th class="px-4 py-2">Type</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(cape, index) in capes" :key="index" class="border-b ">
        <td class=" px-4 py-2">
          {{ cape.owner_username }}
        </td>
        <td class=" px-4 py-2">
          {{ cape.username }}
        </td>
        <td class=" px-4 py-2">
          <select v-model="cape.type" class="text-colors-black">
            <option v-bind:value="0">dev</option>
            <option v-bind:value="1">user</option>
            <option v-bind:value="2">default</option>
          </select>
        </td>
        <td>
          <button @click="setCapeType(cape)">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>
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
import projects from "~/components/projects.vue";

interface Cape {
  uuid: string;
  type: number;
  username: string;
  Enabled: boolean;
  owner_uuid: string;
  owner_username: string;
}

export default Vue.extend({
  name: "CapeTable",
  data() {
    return {
      capes: [] as Cape[]
    }
  },
  methods: {
    async getAllCapes() {
      try {
        let response = await this.$axios.get('/api/capes/all')
        this.capes = response.data
      } catch (err) {
        console.log(err)
      }
    },
    async setCapeType(cape: Cape) {
      try {
        await this.$axios.post(`/api/v1/capes/${cape.uuid}/type`, {
          uuid: cape.uuid,
          type: cape.type
        })
          .then(() =>
            this.$toast.info(`Type for cape ${cape.uuid} changed to ${cape.type}`, {
              duration : 5000
            }));
      } catch (e) {
        this.$toast.error('Error', {
          position: "top-center",
          duration : 800
        });
      }
    }
  },
  async fetch() {
   this.capes = (await this.$axios.get('/api/v1/capes/all')).data
  }

});
</script>

