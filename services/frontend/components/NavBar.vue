<template>
  <nav class="select-none fixed w-full p-4 bg-accent rounded-b-2xl shadow-md ">
    <div class="flex items-center justify-between">

      <!-- Header logo
      <div>
        <QLogo />
      </div>
      -->
      <div>
        <a class="text-xl font-semibold lg:text-2xl hover:text-main">
          <NuxtLink to="/">
            Quantum
          </NuxtLink>
        </a>
      </div>



      <!-- Mobile toggle -->
      <div class="md:hidden">
        <button @click="drawer">
          <svg
            class="h-8 w-8 fill-current text-black"
            fill="none" stroke-linecap="round"
            stroke-linejoin="round" stroke-width="2"
            viewBox="0 0 24 24" stroke="currentColor">
            <path d="M4 6h16M4 12h16M4 18h16"></path>
          </svg>
        </button>
      </div>

      <!-- Navbar -->
      <div class="hidden md:block">
        <ul class="flex space-x-6 text-sm ">
          <li class="hover:border-bluh border-accent border-b-2">
            <NuxtLink to="/">Home</NuxtLink>
          </li>
          <li class="hover:border-bluh border-accent border-b-2">
            <nuxt-link :to="{path: '/', hash: 'projects'}" v-scroll-to="{el: '#projects'}">Projects</nuxt-link>
          </li>
          <li class="hover:border-bluh border-accent border-b-2">
            <nuxt-link to="/github"> GitHub </nuxt-link>
          </li>
          <li v-if="$auth.loggedIn" class="hover:border-bluh border-accent border-b-2">
            <nuxt-link to="/dashboard/"> Profile </nuxt-link>
          </li>
          <li v-if="$auth.user.admid" class="hover:border-bluh border-accent border-b-2">
            <nuxt-link to="/dashboard/admin"> Admin </nuxt-link>
          </li>
          <li>
            <NuxtLink class="cta bg-bluh hover:bg-bluh-600 px-3 py-2 rounded text-white font-semibold" to="/auth/Login">Login</NuxtLink>
          </li>
        </ul>
      </div>

      <!-- Dark Background Transition -->
      <transition
        enter-class="opacity-0"
        enter-active-class="ease-out transition-medium"
        enter-to-class="opacity-100"
        leave-class="opacity-100"
        leave-active-class="ease-out transition-medium"
        leave-to-class="opacity-0"
      >
        <div @keydown.esc="isOpen = false" v-show="isOpen" class="z-10 fixed inset-0 transition-opacity">
          <div @click="isOpen = false" class="absolute inset-0 bg-black opacity-50" tabindex="0"></div>
        </div>
      </transition>

      <!-- Drawer Menu -->
      <aside class="p-5 transform top-0 left-0 w-64 bg-contrast fixed h-full overflow-auto ease-in-out transition-all duration-300 z-30" :class="isOpen ? 'translate-x-0' : '-translate-x-full'">

        <div class="close">
          <button class="absolute top-0 right-0 mt-4 mr-4" @click=" isOpen = false">
            <svg
              class="w-6 h-6"
              fill="none" stroke-linecap="round"
              stroke-linejoin="round" stroke-width="2"
              viewBox="0 0 24 24" stroke="currentColor">
              <path d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <span @click="isOpen = false" class="flex w-full items-center p-4 border-b">
          <img src="/favicon.png" alt="Quantum Logo">
        </span>

        <ul class="divide-y">
          <li><NuxtLink to="/" @click="isOpen = false" class="my-4 inline-block">Home</NuxtLink></li>
          <li>
            <NuxtLink :to="{path: '/', hash: 'projects'}" v-scroll-to="{el: '#projects'}" @click="isOpen = false" class="my-4 inline-block">Projects</NuxtLink>
          </li>
          <li>
            <a href="https://github.com/quantumclient" @click="isOpen = false" class="my-4 inline-block">GitHub</a>
          </li>
          <li v-if="$auth.loggedIn">
            <NuxtLink to="/dashboard" @click="isOpen = false" class="my-4 inline-block">Profile</NuxtLink>
          </li>
          <li v-if="$auth.user.admid">
            <NuxtLink to="/dashboard/admin" @click="isOpen = false" class="my-4 inline-block">Admin Panel</NuxtLink>
          </li>
          <li>
            <NuxtLink to="/auth/Login" @click="isOpen = false" class="my-8 w-full text-center font-semibold cta inline-block bg-bluh hover:bg-bluh-600  px-3 py-2 rounded text-white">Login</NuxtLink>
          </li>
        </ul>

      </aside>

    </div>
  </nav>
</template>

<script>
import Logo from "./Logo";
export default {
  components: {Logo},
  data() {
    return {
      isOpen: false
    };
  },
  methods: {
    drawer() {
      this.isOpen = !this.isOpen;
    }
  },
  watch: {
    isOpen: {
      immediate: true,
      handler(isOpen) {
        if (process.client) {
          if (isOpen) document.body.style.setProperty("overflow", "hidden");
          else document.body.style.removeProperty("overflow");
        }
      }
    }
  },
  mounted() {
    document.addEventListener("keydown", e => {
      if (e.keyCode == 27 && this.isOpen) this.isOpen = false;
    });
  }
};
</script>
