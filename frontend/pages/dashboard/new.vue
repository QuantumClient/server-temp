<template>
  <div class="container min-h-screen pt-20 flex flex-row mx-auto justify-center space-x-3 items-center">
    <div class="flex flex-col space-y-2 p-4">
      <capes class=""/>
      <capes class=" "/>
    </div>
    <div class="flex flex-col bg-accent rounded-2xl p-4">
      <h1 class="text-center text-2xl">Account </h1>
      <div class="border-b" />
      <div class="flex flex-col space-y-2 pb-2">
        <div class="">
          <span v-if=""></span>
          <label for="username">Username:</label>
          <input id="username" class="bg-transparent rounded-md m-1 focus:border-b focus:border-bluh " :rules="[this.$validator.required,this.$validator.username]" v-model="newUsername" type="text" name="username" required>
        </div>
        <button class="bg-bluh p-2 rounded-xl shadow-md" id="update" @click="updateUsername"><span>Update username</span></button>
        <button class="bg-bluh p-2 rounded-xl shadow-md" id="passwordChange" type="submit"><span>Change password</span></button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Download from "~/components/dashboard/download.vue";
import Project from "~/components/dashboard/project.vue";
import authMiddleware from "~/middleware/auth";
import Capes from "~/components/dashboard/Capes.vue";
import validator from "~/utils/validator";

export default Vue.extend({
  name: "new",
  components: {Capes},
  middleware: authMiddleware,
  data() {
    return {
      name: "Quantum",
      newUsername: this.$auth.user.username as string
    }
  },
  methods: {
    updateUsername() {
      if (validator.username(this.newUsername)) {

      }
      console.log(this.newUsername)
    },


  }
})
</script>

<style scoped>

</style>
