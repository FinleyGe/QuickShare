<script setup lang="ts">
import { DeleteFile, GetFileDetail } from '../api';
import { IFileInfo } from '../api/type';
const props = defineProps<{
  fileInfo: IFileInfo
}>();

const emits = defineEmits(["showQRCode", "showDetail"]);

function qrcode() {
  emits("showQRCode", props.fileInfo.hash);
}

async function handleDelete() {
  var r = window.confirm("确定删除吗？");
  if (!r) {
    return;
  }

  var res = await DeleteFile(props.fileInfo.hash);
  if (res.message == "OK") {
    alert("删除成功");
  } else {
    alert("删除失败");
  }
}
async function getDetail() {
  var res = await GetFileDetail(props.fileInfo.hash);
  emits("showDetail", res)
}

</script>

<template>
  <div class="file-item">
    <lable class="name" @click="getDetail">{{ props.fileInfo.name }}</lable>
    <lable class="type">{{ props.fileInfo.type.split("/")[0] }}</lable>
    <button class="showQR" @click="qrcode"> Share </button>
    <button class="delete" @click="handleDelete"> X </button>
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
    max-width: 300px;
    overflow: hidden;
  }
}
</style>
