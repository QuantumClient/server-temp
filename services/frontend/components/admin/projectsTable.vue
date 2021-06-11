<template>
  <div class="bg-white p-4 bg-accent shadow-md rounded-2xl min-w-min h-auto max-w-sm w-72 m-8" >
    <table class="table-auto overflow-x-scroll w-full">
      <thead class="bg-gray-300 border-b">
      <tr>
        <th class="px-4 py-2">Name</th>
        <th class="px-4 py-2">Version</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(project, index) in projects" :key="index" class="border-b ">
        <td class=" px-4 py-2">{{ project.name }}</td>
        <td class=" px-4 py-2">
          <input v-model="project.version" type="text" name="version" class="w-20 focus:outline-none rounded-md text-colors-black" autocomplete="off">
          <button @click="updateVerison(project)">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>
          </button>
        </td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
import moment from 'moment'
import Vue from 'vue'
import projects from "~/components/projects.vue";

interface Project {
  name: string;
  version: string;
  github: string;
}

export default Vue.extend({
  name: "users",
  data() {
    return {
      projects: [] as Project[]
    }
  },
  methods: {
    async getAllUsers() {
      try {
        let response = await this.$axios.get('/api/projects')
        this.projects = response.data
      } catch (err) {
        console.log(err)
      }
    },
    async updateVerison(project: Project) {
      try {
        await this.$axios.put(`/api/projects/${project.name}`,  {
          name: project.name,
          version: project.version
        }).then(() =>
            this.$toast.info(`Verison for ${project.name} has been set to ${project.version}`, {
            duration : 500
        }));
      } catch (e) {
        this.$toast.error('Error', {
          position: "top-center",
          duration : 500
        });
      }
    }
  },
  created() {
    this.getAllUsers();
  }
});
</script>

