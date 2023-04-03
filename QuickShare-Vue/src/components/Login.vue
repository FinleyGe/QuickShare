<script setup lang="ts">
import { ref } from 'vue';
import { useStore } from '../store/init'
import { LoginAPI, GetFileList } from '../api/index'
const emit = defineEmits(["logined"]);
const store = useStore();
const username = ref<string>("");
const password = ref<string>("");
async function handleLogin() {
  var res: Boolean = await LoginAPI(username.value, password.value);
  if (res) {
    alert("登陆成功");
    store.setLogin(true);
    emit("logined");
  }
}
</script>
<template>
  <div class="login">
    <div class="item">
      <label>用户名</label>
      <input type="text" v-model="username" />
    </div>

    <div class="item">
      <label>密码</label>
      <input type="password" v-model="password" />
    </div>
    <button @click="handleLogin">登陆</button>
  </div>
</template>
<style scoped lang="scss">
.login {
  width: fit-content;
  display: grid;
  grid-template-rows: 1fr 1fr 30px;

  .item {
    display: grid;
    grid-template-columns: 1fr 1fr;
    align-items: center;
  }
}
</style>
