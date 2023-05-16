<script setup lang="ts">
import { ref } from 'vue';

const sharecode = ref<string>("");

function getShareCode() {
  console.log(sharecode.value);
}

</script>
<template>
  <div class="sharecode">
    <div class="lable">ShareCode: </div>
    <input v-model="sharecode" />
    <button @click="getShareCode">获取</button>
  </div>
</template>

<style scoped lang="scss">
.sharecode {
  display: flex;
  align-items: center;
  justify-content: center;
  margin-top: 20px;

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
}
</style>
