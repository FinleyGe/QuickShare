<script setup lang="ts">
import { AxiosProgressEvent } from "axios";
import { ref } from "vue";
import { UploadFile } from "../api";
import FileInfo from "./FileInfo.vue";
import ProcessBar from "./ProcessBar.vue";
const file = ref<File | undefined>();
const progress = ref<number>(0);

const hash = ref<string>("");

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

async function upload() {
  if (file.value == undefined) {
    return;
  }
  var res = await UploadFile(file.value,
    function (e: AxiosProgressEvent): void {
      progress.value = e.progress == undefined ? 0 : e!.progress * 100;
      console.log(e);
    });
  if (res.data.message == "OK") {
    hash.value = res.data.data.hash;
    alert("上传成功");
  } else {
    alert("上传失败");
  }
}

function cancel() {
  file.value = undefined;
  progress.value = 0;
}

</script>

<template>
  <div class="upload">
    <div class="drag" @click="handleClickUpload" @drop="handleDrop" @dragover="(e: DragEvent) => { e.preventDefault() }">
      点击或拖放文件到此处
    </div>
    <FileInfo :file="file" :hash="hash" />
    <div class="button-set">
      <button :disabled="file == undefined" @click="upload">上传</button>
      <button @click="cancel">清除</button>
    </div>
    <ProcessBar :percent="progress" v-show="file != undefined" />
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
    display: grid;
    grid-template-columns: 1fr 1fr;
    grid-column-gap: 20px;
    width: 30%;

    button {
      width: 100%;
      border-radius: 10px;
      border: none;
      background-color: white;
      color: black;
      cursor: pointer;
      animation: 0.5s ease-in-out 0s 1 normal none running pulse;
    }

    button:hover {
      background-color: #0077CC;
      color: white;
    }

    button:disabled {
      background-color: gray;
      color: white;
      cursor: not-allowed;
    }

  }

}
</style>
