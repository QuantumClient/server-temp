<template>
  <div class="flex flex-col justify-center w-full px-8 mx-6 my-12 text-center rounded-md md:w-96 lg:w-80 xl:w-64 bg-bluh text-coolGray-100 shadow-md transform duration-500 hover:-translate-y-1">
    <a>Account Key</a>
    <button @click="downloadConfig" class=" py-1 bg-indigo-600 text-white rounded-md text-m focus:outline-none">
      Download
    </button>
  </div>
</template>

<script lang="ts">
import Vue from "vue";

interface UserConfig {
  uuid: string;
  username: string;
  password: string;
}

export default Vue.extend({
  name: "download",
  middleware: 'auth',
  methods: {
    async downloadConfig() {
      try {
        let data = (await this.$axios.$get(`/api/v1/users/${this.$auth.user.uuid}/key`)).data
        this.$download(data, "key.qt")
      } catch (err) {
        console.log(err)
      }
    }
  }
})
</script>

<style scoped>

</style>
