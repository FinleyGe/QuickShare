<script setup lang="ts">
import { IFileInfoDetail } from '../api/type';
import { BaseUrl } from '../const'
const props = defineProps<{
  fileInfo: IFileInfoDetail,
}>();

const emits = defineEmits(["close"]);


function close(this: any) {
  emits.call(this, "close");
}


</script>
<template>
  <div class="popup">
    <div class="info">
      <div class="name">文件名： {{ props.fileInfo.name }}</div>
      <div class="type">类型：{{ props.fileInfo.type }}</div>
      <div classs="size">大小: {{ props.fileInfo.size }} B</div>
      <img :src="BaseUrl + 'get/' + props.fileInfo.hash" v-if="props.fileInfo.type == 'image/png'">
    </div>
    <button class="btn-close" @click="close">Close</button>
  </div>
</template>
<style scoped lang="scss">
.popup {
  z-index: 10;
  height: 100vh;
  width: 100vw;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  position: fixed;
  top: 0;
  left: 0;

}
</style>
