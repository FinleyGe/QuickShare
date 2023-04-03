<script setup lang="ts">
import { ref } from "vue";
import FileInfo from "./FileInfo.vue";
const file = ref<File | undefined>();
const filePath = ref<string>("");

function handleFileChange(e: Event) {
  file.value = (e.target as HTMLInputElement)?.files?.[0];
}

function handleClickUpload() {
  const input = document.createElement("input");
  input.type = "file";
  input.onchange = handleFileChange;
  input.click();
}

function handleDrop(e: DragEvent) {
  e.preventDefault();
  file.value = e.dataTransfer?.files?.[0];
}
</script>

<template>
  <div class="upload">
    <div class="drag" @click="handleClickUpload" @drop="handleDrop" @dragover="(e: DragEvent) => { e.preventDefault() }">
      点击或拖放文件到此处
    </div>
    <FileInfo :file="file" />
    <div class="button-set">
      <button :disabled="file == undefined">上传</button>
      <button>取消</button>
    </div>
  </div>
</template>

<style scoped lang="scss">
@import "../styles/colors";

.upload {
  width: 80%;
  margin: auto;
  display: grid;

  .drag {
    height: 100px;
    background-color: #AED6F1;
    width: 80%;
    margin: auto;
    margin-block: 20px;
    border-radius: 20px;
    border: 2px solid #0077CC;
    text-align: center;
    justify-content: center;
    align-items: center;
    display: grid;
    cursor: pointer;
  }

  .button-set {
    margin: auto;
  }

}
</style>
