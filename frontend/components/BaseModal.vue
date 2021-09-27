<template>
  <div class="modal-backdrop" @click.prevent="close">
    <div @click.stop >
      <slot />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  name: "BaseModal",
  methods: {
    close() {
      this.$emit('close');
    },
    handleEscape (e: KeyboardEvent) {
      if (e.key === 'Esc' || e.key === 'Escape') {
        this.close();
      }
    }
  },
  beforeMount () {
    document.addEventListener('keydown', this.handleEscape)
  },
  beforeDestroy () {
    document.removeEventListener('keydown', this.handleEscape)
  },
})
</script>

<style scoped>
.modal-backdrop {
  position: fixed;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  background-color: rgba(0, 0, 0, 0.4);
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
