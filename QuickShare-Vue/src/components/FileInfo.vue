<script setup lang="ts">
const props = defineProps<{
  file: File | undefined;
}>();

</script>
<template>
  <div class="fileInfo" v-show="props.file != undefined">
    <div class="filename">文件名: {{ props.file?.name }}</div>
    <div class="filesize">文件大小: {{ props.file?.size + 'b' }}</div>
    <div class="filetype">文件类型: {{ props.file?.type }}</div>
  </div>
</template>
<style scoped lang="scss">
.fileInfo {
  display: grid;
  grid-template-columns: 1fr 1fr 1fr;
  grid-template-rows: 1fr 1fr;
  grid-template-areas:
    "filename filename filename"
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
}
</style>
