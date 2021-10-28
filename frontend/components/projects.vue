<template>
  <div class="bg-accent rounded-2xl shadow-md">
    <section class="py-6 text-white">
      <div class="container flex flex-col items-center justify-center p-4 mx-auto sm:p-10">
        <h1 class="text-6xl text-colors-white sm:text-8xl text-accent pb-4 leading-none text-center">Projects</h1>
        <div class="flex flex-row flex-wrap-reverse justify-center mt-8 ">
          <div class="flex flex-col justify-center w-full px-8 mx-6 my-12 text-center rounded-md md:w-96 lg:w-80 xl:w-64 bg-contrast shadow-md cursor-pointer transform duration-500 hover:-translate-y-1">
            <nuxt-link to="/projects/qubit">
              <div class="my-4">
                <p class="text-xl font-semibold leading-snug">Qubit</p>
                <p>Fast, safe, and powerful public client.</p>
              </div>
            </nuxt-link>
          </div>
          <div class="flex flex-col justify-center w-full px-8 mx-6 my-12 text-center rounded-md md:w-96 lg:w-80 xl:w-64 bg-contrast shadow-md transform duration-500 hover:-translate-y-1">
           <nuxt-link to="/projects/quantum">
             <div class="my-4">
               <p class="text-xl font-semibold leading-snug">Quantum</p>
               <p>Our top of the line private client.</p>
             </div>
           </nuxt-link>

          </div>
          <div class="flex flex-col justify-center w-full px-8 mx-6 my-12 text-center rounded-md md:w-96 lg:w-80 xl:w-64 bg-contrast shadow-md transform duration-500 hover:-translate-y-1">
          <div class="my-4">
            <p class="text-xl font-semibold leading-snug">Energy</p>
            <p>The speedy event bus that powers all of our projects.</p>
          </div>
        </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: "projects"
}
</script>

<style scoped>

</style>
