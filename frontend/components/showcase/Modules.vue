<template>
  <div class="container_custom">
    <section class="py-6">
      <div class="justify-center p-4 sm:p-10">
        <h1 class="text-6xl sm:text-8xl text-accent pb-4 text-accent leading-none text-center">Modules</h1>
        <div class="container flex flex-row flex-wrap">
          <Category :title="'Combat'" :fs="modules.Combat" />
          <Category :title="'Misc'" :fs="modules.Misc" />
          <Category :title="'Movement'" :fs="modules.Movement" />
          <Category :title="'Player'" :fs="modules.Player" />
          <Category :title="'Render'" :fs="modules.Render" />
          <Category :title="'World'" :fs="modules.World" />
          <Category :title="'Client'" :fs="modules.Client" />
        </div>
      </div>
    </section>
  </div>

</template>

<script lang="ts">
import Vue from "vue";
import Category from "~/components/showcase/Category.vue";

export default Vue.extend({
  name: "Modules",
  components: {Category},
  props: {
    modules: {} as any
  }
})
</script>

<style scoped>

</style>
