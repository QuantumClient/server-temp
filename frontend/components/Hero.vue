<template>
  <div class="container_custom">
    <div>
      <h1 class="text-7xl sm:text-9xl text-accent pb-4">
        Quantum
      </h1>
      <h2 class="text-4xl sm:text-5xl text-accent pb-3 sm:pt-1">
        Modern utility mods
      </h2>
      <div class="links">
        <NuxtLink to="/discord" class="button">
          Discord
        </NuxtLink>
        <NuxtLink to="/github" class="button">
          Github
        </NuxtLink>
        <NuxtLink to="/projects/qubit" class="button">
          Download
        </NuxtLink>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Hero"
}
</script>

<style scoped>

</style>
