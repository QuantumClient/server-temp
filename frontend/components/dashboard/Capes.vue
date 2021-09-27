<template>
  <div class="bg-white p-4 bg-accent shadow-md rounded-2xl">
    <div class="px-4 py-2">
      <a class=" text-2xl">Capes</a>
      <button type="button" @click="showModal">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="w-5 h-5 "
          viewBox="0 0 24 24"
          stroke-width="2.3"
          stroke="#ffffff"
          fill="none"
          stroke-linecap="round"
          stroke-linejoin="round"
        >
          <line x1="11.5" y1="4" x2="11.5" y2="20" />
          <line x1="4" y1="11.5" x2="20" y2="11.5" />
        </svg>
      </button>
    </div>

    <div class="bg-gray-300 border-b"/>

    <CapeModal :state="isModalVisible" @close="isModalVisible = false" @add="(cape) => capes.push(cape)" />
    <table class="min-w-full table-auto overflow-x-scroll ">
      <thead class="bg-gray-300 border-b">
      <tr>
        <th class="px-4 py-2">Username</th>
        <th class="px-4 py-2">Enabled</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(cape, index) in capes" :key="capes" class="border-b ">
        <td class=" px-4 py-2">{{ cape.username }}</td>
        <td class=" px-4 py-2">
          <button @click="setEnabled(cape)">
            <div v-if="cape.enabled" >
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
      </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import CapeModal from "~/components/CapeModal.vue";
interface Cape {
  uuid: string;
  username: string;
  enabled: boolean;
  type: number;
}

export default Vue.extend({
  name: "capes",
  components: {CapeModal},
  data: () => {
    return {
      isModalVisible: false,
      capes: [] as Cape[]
    }
  },
  methods: {
    showModal() {
      this.isModalVisible = true;
    },
    async setEnabled(cape: Cape) {
      try {
        await this.$axios.post(`/api/v1/capes/${cape.uuid}/enabled`)
          .then(() => cape.enabled = !cape.enabled);
      } catch (e) {
        this.$toast.error('Error', {
          position: "top-center",
          duration : 500
        });
      }
    },
    async GetAllCapes() {
      this.capes = (await this.$axios.get(`/api/v1/users/${this.$auth.user.uuid}/capes`)).data
    }
  },
  mounted() {
    this.GetAllCapes();
  }
})
</script>

<style scoped>

</style>
