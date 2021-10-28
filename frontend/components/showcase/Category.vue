<template>
  <ul class="w-auto h-full px-3 mx-4 bg-contrast rounded-md">
    <a class="text-center border-b"> {{this.$props.title}}</a>
    <li v-for="f in this.fs" class="overflow-hidden">
      <Feature :fe="f" />
    </li>
  </ul>
</template>

<script lang="ts">
import Vue from "vue";
import Feature from "~/components/showcase/Feature.vue";

export default Vue.extend({
  name: "Category",
  components: {Feature},
  props: {
    fs: {} as any,
    title: String
  }
})
</script>

<style scoped>

</style>
