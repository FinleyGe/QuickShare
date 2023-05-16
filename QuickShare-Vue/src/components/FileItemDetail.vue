<script setup lang="ts">
import { computed } from 'vue';
import { IFileInfoDetail } from '../api/type';
import { BaseUrl } from '../const'
const props = defineProps<{
  fileInfo: IFileInfoDetail | undefined,
}>();

const emits = defineEmits(["close"]);


function close(this: any) {
  emits.call(this, "close");
}

const fileSize = computed(() => {
  if (props.fileInfo) {
    let size = props.fileInfo.size;
    let unit = "B";
    if (size > 1024) {
      size /= 1024;
      unit = "KB";
    }
    if (size > 1024) {
      size /= 1024;
      unit = "MB";
    }
    if (size > 1024) {
      size /= 1024;
      unit = "GB";
    }
    return size.toFixed(2) + unit;
  }
  return "";
})


</script>
<template>
  <div class="popup">
    <div class="info">
      <div class="name">文件名： {{ props.fileInfo?.name }}</div>
      <div class="type">类型：{{ props.fileInfo?.type }}</div>
      <div classs="size">大小: {{ fileSize }} </div>
      <img :src="BaseUrl + 'get/' + props.fileInfo?.hash" v-if="props.fileInfo?.type == 'image/png'">
    </div>
    <button @click="">Share</button>
    <button @click="close">Close</button>
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

  .info {
    background-color: white;
    padding: 10px;
    border-radius: 10px;
    max-height: 60vh;
  }

  button {
    margin-top: 10px;
    padding: 10px;
    border-radius: 10px;
  }
}
</style>
