<script setup lang="ts">
import {message} from "ant-design-vue";
import {ref} from "vue";
import {UploadFile} from "../api";
import {BaseUrl} from "../const";
import QRCode from 'qrcode-vue3'
const file = ref<File>();

const filePath = ref<string>("");

function handleFileChange(e: Event) {
  file.value = (e.target as HTMLInputElement)?.files?.[0];
  UploadFile(file.value!)
    .then((_res) => {
      message.success("upload success");
      console.log(_res);
      filePath.value = BaseUrl + "get/" + _res.data.data.hash;
    })
    .catch((_err) => {
      message.error("upload failed");
    });
}
</script>

<template>
  {{ filePath}}
  <QRCode :value="filePath" />
  <input type="file" v-on:change="handleFileChange" />
</template>

<style scoped>

</style>
