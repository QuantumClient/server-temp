<template>
  <div class=" p-4 bg-accent shadow-md rounded-2xl min-w-min h-auto max-w-sm w-72 m-8" >
    <table class="table-auto overflow-x-scroll w-full">
      <thead class="bg-gray-300 border-b">
      <tr>
        <th class="px-4 py-2">UUID</th>
        <th class="px-4 py-2">Type</th>
        <th @click="toggleModal" class="px-4 py-2">Actions</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(cape, index) in capes" :key="index" class="border-b ">
        <td class=" px-4 py-2">
          {{ cape.uuid }}
        </td>
        <td class=" px-4 py-2">
          <select v-model="cape.type" class="text-colors-black">
            <option v-bind:value="0">dev</option>
            <option v-bind:value="1">user</option>
          </select>
        </td>
        <td>
          <button @click="setCapeType(cape)">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>
          </button>
          <button @click="deleteCape(cape)">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="w-5 h-5"
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
          </button>
        </td>
      </tr>
      </tbody>
    </table>

    <div v-if="showNew" class="overflow-x-hidden overflow-y-auto fixed inset-0 z-50 outline-none focus:outline-none justify-center items-center flex">
      <div class="relative w-auto my-6 mx-auto max-w-sm">
        <!--content-->
        <div class="border-2  shadow-md rounded-2xl relative flex flex-col w-full bg-contrast outline-none focus:outline-none">
          <!--header-->
          <div class="flex items-start justify-between p-5 border-b border-solid border-blueGray-200 rounded-t">
            <h3 class="text-3xl font-semibold">
              New Cape
            </h3>
          </div>
          <!--body-->
          <div class="relative p-6 flex-wrap">
            <input v-model="newCape.uuid" type="text" name="version" class="w-20 focus:outline-none rounded-md text-colors-black" autocomplete="off">
            <select v-model="newCape.type" class="text-colors-black">
              <option v-bind:value="0">dev</option>
              <option v-bind:value="1">user</option>
            </select>
          </div>
          <div class="flex items-center justify-end p-6 border-t border-solid border-blueGray-200 rounded-b">
            <button class="text-red-500 bg-transparent border border-solid border-red-500 hover:bg-red-500 hover:text-white active:bg-red-600 font-bold uppercase text-sm px-6 py-3 rounded outline-none focus:outline-none mr-1 mb-1 ease-linear transition-all duration-150" type="button" v-on:click="toggleModal()">
              Close
            </button>
            <button class="text-red-500 background-transparent font-bold uppercase px-6 py-2 text-sm outline-none focus:outline-none mr-1 mb-1 ease-linear transition-all duration-150" type="button" v-on:click="addCape(newCape)">
              Save Changes
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import moment from 'moment'
import Vue from 'vue'
import projects from "~/components/projects.vue";

interface Cape {
  uuid: string;
  type: number;
}

export default Vue.extend({
  name: "CapeTable",
  data() {
    return {
      showNew: false,
      capes: [] as Cape[],
      newCape: {} as Cape
    }
  },
  methods: {
    async getAllCapes() {
      try {
        let response = await this.$axios.get('/api/capes?form=true')
        this.capes = response.data
      } catch (err) {
        console.log(err)
      }
    },
    async deleteCape(cape: Cape) {
      try {
        await this.$axios.delete(`/api/capes/${cape.uuid}`)
          .then(() => this.capes.forEach((capeA, index) => {
            if (capeA.uuid == cape.uuid)  delete this.capes[index]
          })).then(() =>
            this.$toast.info(`Deleted cape ${cape.uuid}`, {
              duration : 5000
            }));
      } catch (e) {
        this.$toast.error('Error', {
          position: "top-center",
          duration : 800
        });
      }
    },
    async addCape(cape: Cape) {
      try {
        await this.$axios.put(`/api/capes`, {
          uuid: cape.uuid,
          type: cape.type
        })
          .then(() => this.capes.push(cape)).then(() =>
            this.$toast.info(`Added cape ${cape.uuid}`, {
              duration : 5000
            }));
      } catch (e) {
        this.$toast.error('Error', {
          position: "top-center",
          duration : 800
        });
      }
      this.toggleModal();
    },
    async setCapeType(cape: Cape) {
      try {
        await this.$axios.put(`/api/capes/${cape.uuid}`, {
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
    },
    toggleModal: function() {
      this.showNew = !this.showNew;
    }
  },
  created() {
    this.getAllCapes();
  }
});
</script>

