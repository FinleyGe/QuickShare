<script setup lang="ts">
import { computed, ref } from 'vue';
import { GetHashByShareCode } from '../api';
import { BaseUrl } from '../const';
const sharecode = ref<string>("");
const hash = ref<string>("");

async function getShareCode() {
  if (sharecode.value == "") {
    alert("请输入ShareCode");
    return;
  }

  let res = await GetHashByShareCode(sharecode.value);
  if (res) {
    hash.value = res.hash;
  } else {
    alert("获取Hash失败");
  }
}

const link = computed(() => {
  if (hash.value == "") { return ""; }
  return BaseUrl + 'get/' + hash.value;
})

</script>
<template>
  <div class="sharecode">
    <div class="lable">ShareCode: </div>
    <input v-model="sharecode" />
    <button @click="getShareCode"> 获取 </button>
    <a :href="link" target="_blank" v-if="hash != ''">{{ link }}</a>
  </div>
</template>

<style scoped lang="scss">
.sharecode {
  display: grid;
  align-items: center;
  justify-content: center;
  margin-top: 20px;
  width: 80%;
  margin-inline: auto;
  grid-template-areas: "lable input button"
    "link link link";
  grid-template-columns: min-content min-content min-content;

  .lable {
    margin-right: 10px;
  }

  input {
    width: 200px;
    height: 30px;
    border: 1px solid #ccc;
    border-radius: 5px;
    padding: 0 10px;
  }

  button {
    width: 60px;
    height: 30px;
    border: 1px solid #ccc;
    border-radius: 5px;
    margin-left: 10px;
  }

  a {
    grid-area: link;
  }
}
</style>
