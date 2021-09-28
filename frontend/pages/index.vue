<template>
  <div>
   <hero/>
    <projects id="projects" class="anchor"/>
    <devs />
    <button class="button" @click="$auth.refreshTokens()"> click </button>

  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import Hero from "~/components/Hero.vue";
import Devs from "~/components/Devs.vue";
import Projects from "~/components/projects.vue";

export default Vue.extend({
  components: {Projects, Hero, Devs}
})
</script>

<style>
</style>
