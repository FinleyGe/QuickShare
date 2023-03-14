<script setup lang="ts">
import { GetFile } from '../api';
import { IFileInfo } from '../api/type';
const props = defineProps<{
  fileInfo: IFileInfo
}>()
const emits = defineEmits(["showQRCode"]);
function download() {
  GetFile(props.fileInfo.hash);
}

function qrcode() {
  emits("showQRCode", props.fileInfo.hash);
}
</script>

<template>
  <div class="file-item">
    <lable class="name">{{ props.fileInfo.name }}</lable>
    <label class="size"> {{ props.fileInfo.size }}</label>
    <label class="count"> {{ props.fileInfo.download_count }}</label>
    <button class="download" @click="download"> Download </button>
    <button class="showQR" @click="qrcode"> QR-Code </button>
  </div>
</template>
<style lang="scss">
.file-item {
  display: flex;
  flex-direction: row;
  align-items: center;
  height: 50px;
  column-gap: 10px;

  .name {
    width: 300px;
    overflow: hidden;
  }
}
</style>
