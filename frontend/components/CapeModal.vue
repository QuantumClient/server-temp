<template>
  <BaseModal v-show="state">
    <img v-show="hover" src="~/assets/images/mc-server.png" alt="Minecraft Server" class="w-60 mx-auto bg-contrast p-1.5 shadow-md object-scale-down rounded-2xl my-2.5">
    <div class="container bg-contrast shadow-md rounded-2xl justify-center flex flex-col text-center mx-auto p-4 max-w-xl">
      <div @mouseover="mouseover">
        <h1 class="px-4 text-2xl">Link Minecraft Account</h1>
        <p  class="px-4 py-2 text-lg"> Join quantumclient.org to get a Key </p>
      </div>
      <input v-model="key" type="text" name="key" class="p-2 mx-auto mb-0.5 text-accent rounded-xl border-4 focus:outline-none focus:border-bluh" autocomplete="off" placeholder="x0k-9k8-bjf">
      <div class="mx-auto p-4">
        <button type="button" @click="addAccount() && close()" class="p-2 px-3.5 bg-accent rounded-xl shadow-md">Link</button>
        <button type="button" @click="close()" class="p-2 px-3.5 bg-colors-red-700 rounded-xl shadow-md">Close</button>
      </div>
    </div>
  </BaseModal>
</template>

<script lang="ts">
import Vue from "vue";
import BaseModal from "~/components/BaseModal.vue";
export default Vue.extend({
  name: "capeModal",
  components: {BaseModal},
  props: {
    state: {
      type: Boolean,
      default: false
    }
  },
  data: () => {
    return {
      key: "",
      hover: false,
      open: false
    }
  },
  methods: {
    close() {
      this.$props.state = false;
      this.$emit('close')
    },
    mouseover() {
      this.hover = !this.hover
    },

    async addAccount() {
      try {
        await this.$axios.post(`/api/v1/users/${this.$auth.user.uuid}/link?key=${this.key.replace(/-/g, "")}`)
          .then((res) => {
            this.$emit('add', res.data)
            this.key = ""
          });
      } catch (e) {
        this.$toast.error('Error', {
          position: "top-center",
          duration : 500
        });
      }
    }
  }
})
</script>

