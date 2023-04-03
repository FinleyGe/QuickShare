<script setup lang="ts">
import { GetFile, DeleteFile } from '../api';
import { IFileInfo } from '../api/type';
import { BaseUrl } from '../const';
const props = defineProps<{
  fileInfo: IFileInfo
}>();

const emits = defineEmits(["showQRCode"]);
function download() {
  GetFile(props.fileInfo.hash);
}

function qrcode() {
  emits("showQRCode", props.fileInfo.hash);
}

async function handleDelete() {
  var res = await DeleteFile(props.fileInfo.hash);
  if (res) {
    alert("删除成功");
  } else {
    alert("删除失败");
  }
}
</script>

<template>
  <div class="file-item">
    <lable class="name">{{ props.fileInfo.name }}</lable>
    <button class="download" @click="download"> Download </button>
    <button class="showQR" @click="qrcode"> Share </button>
    <button class="delete" @click="handleDelete"> Delete </button>
  </div>
</template>
<style scoped lang="scss">
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
