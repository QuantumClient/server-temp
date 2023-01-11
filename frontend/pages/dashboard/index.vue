<template>
  <div class="container_custom  flex-wrap">
    <Download />
    <Project :name="this.name" />
    <Capes />
    <div class="overflow-hidden shadow-lg bg-accent transition duration-500 ease-in-out transform hover:-translate-y-5 hover:shadow-2xl rounded-lg h-90 w-60 md:w-80 m-auto">
      <h1 class="text-2xl font-medium border-b py-2"> {{ $auth.user.username }} </h1>
      <div class=" bg-white w-full p-4">
        <p class="text-gray-800 text-sm font-medium mb-2">
          A comprehensive guide about online education.
        </p>
        <button class="bg-bluh rounded-md">
          Change username
        </button>
        <button class="bg-bluh rounded-md">
          Change Password
        </button>
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

export default Vue.extend({
  name: "index",
  components: {Capes, Project, Download},
  middleware: authMiddleware,
  data() {
    return {
      name: "Quantum"
    }
  },
})
</script>

<style scoped>

</style>
