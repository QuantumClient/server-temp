<template>
  <div>
    <div class="justify-center p-2 rounded-t-2xl  bg-accent">
      <h2 class="text-xl text-center">Key</h2>
    </div>

    <div class="p-5 bg-contrast">
      <label class="text-s my-2.5">Password</label>
      <input
        v-model="user.password"
        type="password"
        class="w-full h-full appearance-none rounded-md focus:border-bluh text-accent"
      />
    </div>

    <div @click="downloadConfig" class="p-2 rounded-b-2xl bg-accent hover:bg-bluh shadow-md">
      <button class=" py-1 bg-indigo-600 text-white rounded-md text-m focus:outline-none">
        Download
      </button>
    </div>
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
  data() {
    return {
      user: {username: "", uuid: ""} as UserConfig
    }
  },
  methods: {
    async downloadConfig() {
      try {

        this.user.username = this.$auth.user.username.toString()
        this.user.uuid = this.$auth.user.uuid.toString()
        this.$download(this.user, "key.qt")
      } catch (err) {
        console.log(err)
      }
    }
  }
})
</script>

<style scoped>

</style>
