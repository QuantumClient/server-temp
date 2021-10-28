<template>
  <div>
    <section class="py-6">
      <div class="justify-center bg-accent rounded-md p-4 sm:p-10">
        <h1 class="text-6xl sm:text-8xl text-main pb-4 leading-none text-center">HUD Items</h1>
        <div class="flex flex-wrap">
          <li v-for="f in this.items" class="overflow-hidden">
            <Feature :fe="f" />
          </li>
        </div>
      </div>
    </section>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Feature from "~/components/showcase/Feature.vue";

export default Vue.extend({
  name: "HUD",
  components: {Feature},
  props: {
    items: {} as any
  }
})
</script>

<style scoped>

</style>
