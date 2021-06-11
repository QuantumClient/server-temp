<template>
  <div class="container_custom flex w-full mx-auto pt-20 ">
    <div class=" px-4 md:px-0 md:mt-8 mb-16 text-gray-800 flex flex-wrap justify-center">
      <AdminUsers />
      <AdminProjectsTable />
      <AdminCapeTable />
    </div>
  </div>
</template>

<script lang="ts">
import adminMiddleware from "~/middleware/admin";
import Vue from "vue";

export default Vue.extend({
  name: "index",
  middleware: adminMiddleware
})
</script>

<style scoped>

</style>
