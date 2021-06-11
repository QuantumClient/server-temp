<template>
  <div class="w-screen h-screen flex justify-center items-center">
    <form @submit.prevent="login" class="p-10 flex justify-center items-center flex-col bg-accent rounded-2xl shadow-md">
      <p class="subtitle mb-5 text-3xl">Login</p>
      <input :rules="[this.$validator.username]" v-model="username" type="text" name="username" class="text-accent mb-5 p-3 w-80 rounded-xl border-4 focus:outline-none focus:border-bluh" autocomplete="off" placeholder="Username" required>
      <input :rules="[this.$validator.password]" v-model="password" type="password" name="password" class="text-accent mb-5 p-3  w-80 border-4 focus:outline-none focus:border-bluh rounded-xl" autocomplete="off" placeholder="Password" required>
      <button class="bg-bluh font-bold p-2 rounded-xl shadow-md w-80 mb-5" id="login" type="submit"><span>Login</span></button>
      <p class="text-sm">Don't have an account yet?
        <NuxtLink to="/auth/Signup" class="hover:underline">Sign up.</NuxtLink>
      </p>
    </form>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  name: "Signup",
  data() {
    return {
      username: '',
      password: ''
    }
  },
  methods: {
    async login() {
      try {
        await this.$auth.loginWith('local', {
          data: {
            username: this.username,
            password: this.password
          }
        }).then(() => this.$toast.success('Logged in', {
          position: "top-center",
          duration : 5000
        }));
      } catch (e) {
        this.$toast.error('Error while authenticating', {
          position: "top-center",
          duration : 5000
        });
      }
    }
  }
});
</script>

<style scoped>

</style>
