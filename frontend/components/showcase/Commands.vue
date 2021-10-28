<template>
  <div >
    <section class="py-6">
      <div class="justify-center rounded-md p-4 sm:p-10 " >
        <h1 class="text-6xl sm:text-8xl pb-4 leading-none text-center">Commands</h1>
        <div class="mx-auto min-h-full flex flex-wrap flex-row h-auto">
          <li v-for="f in this.items" class="overflow-hidden mx-auto justify-between">
            <div class="bg-accent shadow-md rounded-md p-2.5 my-2 text-center transform duration-500 hover:-translate-y-1">
              <h1 class="text-lg border-b border-colors-coolGray-100"> {{f.Name}}</h1>
              <span class="text-sm text-colors-coolGray-100"> {{ f.Description }} </span>
            </div>
          </li>
        </div>
      </div>
    </section>
  </div>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  name: "Commands",
  props: {
    items: {} as any
  }
})
</script>

<style scoped>

</style>
