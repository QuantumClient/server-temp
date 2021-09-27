<template>
  <div class="container_custom  flex-wrap">
    <Download />
    <Project :name="this.name" />
    <Capes />
    <h2 class="subtitle text-contrast"> If you don't have private access being here is pointless.</h2>
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
