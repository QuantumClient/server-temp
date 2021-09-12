<template>
  <div class="container_custom mx-auto flex flex-wrap items-center text-accent">
    <div class="w-full md:w-5/12 text-center pt-20">
      <ServerDog />
    </div>
    <div class="w-full md:w-7/12 text-center md:text-left p-4">
      <h1 class="title">
        {{ error.statusCode }} - error
      </h1>
      <h2 class="subtitle">
        {{ error.message }}
      </h2>
      <a
        rel="noopener noreferrer"
        target="_blank"
        class="button--green"
      >
        <NuxtLink to="/">Home</NuxtLink>
      </a>
    </div>
  </div>
</template>

<script>
import ServerDog from "~/components/ServerDog";
export default {
  components: {ServerDog},
  props: ['error'],
  layout: 'error' // you can set a custom layout for the error page
}
</script>
