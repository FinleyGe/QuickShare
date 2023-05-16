<script setup lang="ts">
import { computed } from 'vue';
import { BaseUrl } from '../const';

const props = defineProps<{
  file: File | undefined;
  hash: string | undefined;
}>();

const link = computed(() => {
  if (props.hash) {
    return BaseUrl + "get/" + props.hash;
  }
  return "未上传";
});

</script>
<template>
  <div class="fileInfo" v-show="props.file != undefined">
    <div class="filename">文件名: {{ props.file?.name }}</div>
    <div class="filesize">文件大小: {{ props.file?.size + 'b' }}</div>
    <div class="filetype">文件类型: {{ props.file?.type }}</div>
    <div class="link">链接: {{ link }}</div>
  </div>
</template>
<style scoped lang="scss">
.fileInfo {
  display: grid;
  grid-template-columns: 1fr 1fr 1fr;
  grid-template-rows: 1fr 1fr;
  grid-template-areas:
    "filename link link"
    "filesize filetype filetype";

  .filename {
    grid-area: filename;
  }

  .filesize {
    grid-area: filesize;
  }

  .filetype {
    grid-area: filetype;
  }

  .link {
    grid-area: link;
  }
}
</style>
