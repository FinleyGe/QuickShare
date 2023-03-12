<script setup lang="ts">
import Upload from '../components/Upload.vue'
import { useStore } from '../store/init'
import type { DeviceSize } from '../store/init'
import { ref } from 'vue';

const store = useStore()
const pageSize = ref<DeviceSize>(store.deviceSize);
</script>

<template>
  <div class="base" :class="pageSize">
    <header>
      <div class="title">
        QuickShare
      </div>
      <div class="button-set">
        <button>登陆</button>
        <button>注册</button>
      </div>
    </header>
    <main>
      <Upload />
    </main>
    <footer>

    </footer>
  </div>
</template>

<style scoped lang="scss">
@import "../styles/colors.scss";

.base {
  min-height: 100vh;
  background-color: $background-light;

  &.large {
    display: grid;
    grid-template-rows: 80px 1fr 30px;
  }

  header {
    display: grid;
    grid-template-columns: 1fr 100px;
    padding-inline: 10px;
    align-items: center;
    background-color: $background-dark;

    .title {
      background-color: $background-light;
      color: $text-dark;
      font-size: 2rem;
      text-align: center;
      padding: 10px;
      width: fit-content;
      border-radius: 1rem;
    }
  }

  main {
    height: 200px;
  }
}
</style>
