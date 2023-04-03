<script setup lang="ts">
import { onMounted } from 'vue';
import { RouterView } from 'vue-router';
import { useStore } from './store/init'
const store = useStore();
function handleWindowsResize() {
  let width = document.documentElement.clientWidth;

  if (width < 768) {
    store.setDeviceSize("small")
  } else if (width < 1024) {
    store.setDeviceSize("medium")
  } else {
    store.setDeviceSize("large")
  }
}


onMounted(() => {
  handleWindowsResize();
  window.onresize = handleWindowsResize;
});
</script>

<template>
  <router-view />
</template>

<style scoped lang="scss"></style>
