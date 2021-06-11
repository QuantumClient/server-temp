<template>
  <div class="container_custom">
    <a class="block max-w-sm gap-3 mx-auto sm:max-w-full group hover:no-underline focus:no-underline lg:grid lg:grid-cols-12 bg-coolGray-50">
      <img src="~/assets/images/Qubit-1.png" alt="Qubit Minecraft Client" class="m-4 rounded sm:h-96 lg:col-span-7 shadow-md rounded-2xl ">
      <div class="p-6 space-y-2 lg:col-span-5 text-contrast">
        <h3 class="text-2xl font-semibold sm:text-4xl ">Qubit</h3>
        <span class="text-s text-colors-coolGray-600">{{ project.version }}</span>
        <p class="text-lg">Qubit is a modern minecraft utility mod made with fabric. It has many advanced features to aid your game play.</p>
        <a :href="this.project.link" target="_blank">
          <button class="py-2 w-full px-8 mx-6 my-8 font-semibold rounded-md md:w-96 lg:w-80 xl:w-64 bg-contrast text-colors-white shadow-md cursor-pointer transform duration-500 hover:-translate-y-1">
            Download
          </button>
        </a>
      </div>
    </a>
  </div>
</template>

<script lang="ts">
import Vue from "vue";

interface Project {
  name: string;
  version: string;
  github: string;
  link: string;
}

export default Vue.extend({
  name: "qubit",
  data() {
    return {
      project: {} as Project
    }
  },
  methods: {
    async getQubit() {
      try {
        let response = await this.$axios.get('/api/projects/qubit')
        this.project = response.data
      } catch (err) {
        console.log(err)
      }
    }
  },
  created() {
    this.getQubit();
  }
})
</script>

<style scoped>

</style>
