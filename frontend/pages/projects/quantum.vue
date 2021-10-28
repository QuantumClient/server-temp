<template>
  <div >
    <div class="mb-5 container flex flex-wrap justify-center mx-auto pt-20 space-y-2 ">
      <img class=" sm:h-96  shadow-md rounded-2xl " src="/Q-p1.png"  alt="Quantum Minecraft 1.17 Hack Client "/>
      <p class="text-contrast text-lg mx-20">Quantum is a Minecraft utility mod for fabric 1.17; it has almost all the features you would want in a client. Highway digging is one the client's strongest feature, but frankly it will do what ever you want.</p>
    </div>
    <modules class="" :modules="ms"/>
    <HUD :items="HUDItems"/>
    <Commands :items="cs"/>
    <p class="text-center text-base text-colors-gray-50 p-4">Data isn't very accurate</p>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Modules from "~/components/showcase/Modules.vue";
import HUD from "~/components/showcase/HUD.vue";
import Commands from "~/components/showcase/Commands.vue";

export default Vue.extend({
  name: "quantum",
  components: {Commands, HUD, Modules},
  data: () => {
    return {
      ms: {},
      HUDItems: [],
      cs: []
    }
  },
  methods: {
    async getData() {
      const res = (await this.$axios.get('/Quantum-Manifest.json')).data
      this.ms = res.Modules
      this.HUDItems = res.HUDItems
      this.cs = res.Commands;
    }
  },
  mounted() {
    this.getData()
  }
})
</script>

<style scoped>

</style>
