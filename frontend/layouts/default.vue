<template>
  <div>
    <NavBar />
    <Nuxt />
    <Footer />
  </div>
</template>

<style>

*,
*::before,
*::after {
  box-sizing: border-box;
  margin: 0;
}

</style>
<script>
import NavBar from "../components/NavBar";
import Footer from "../components/Footer";

export default {
  components: {Footer, NavBar}
}
</script>
