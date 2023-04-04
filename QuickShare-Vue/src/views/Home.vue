<script setup lang="ts">
import Upload from '../components/Upload.vue'
import FileItemDetail from '../components/FileItemDetail.vue'
import { useStore } from '../store/init'
import type { DeviceSize } from '../store/init'
import { ref } from 'vue';
import Login from '../components/Login.vue'
import { GetFileList } from '../api/index'
import FileItem from '../components/FileItem.vue'
import { IFileInfo, IFileInfoDetail } from '../api/type';
import QRCode from '../components/QRCode.vue'

const store = useStore()
const pageSize = ref<DeviceSize>(store.deviceSize);
const showLogin = ref<boolean>(false);
const fileList = ref<IFileInfo[]>([]);
const QRCodePath = ref<string>("");
const isShowQRCode = ref<boolean>(false);
const isShowDetail = ref<boolean>(false);
const fileDetail = ref<IFileInfoDetail | undefined>();
function handleLogin() {
  if (!store.isLogin) {
    showLogin.value = !showLogin.value;
  }
}
function Logined() {
  console.log("logined");
  showLogin.value = false;
}
async function refresh() {
  fileList.value = await GetFileList();
  console.log(fileList.value)
}
function showQRCode(path: string) {
  QRCodePath.value = path;
  isShowQRCode.value = true;
}

function showDetail(detail: IFileInfoDetail) {
  fileDetail.value = detail
  isShowDetail.value = true;
}
</script>

<template>
  <div class="base" :class="pageSize">
    <header>
      <div class="title">
        QuickShare
      </div>
      <div class="button-set">
        <button @click="handleLogin">{{ showLogin ? '返回' : '登陆' }}</button>
        <button @click="refresh">刷新</button>
      </div>
    </header>
    <main>
      <Upload v-show="!showLogin" />
      <Login v-show="showLogin" @logined="Logined" />
      <QRCode @close="isShowQRCode = false" :path="QRCodePath" v-show="isShowQRCode" />
      <FileItemDetail @close="isShowDetail = false" :file-info="fileDetail" v-show="isShowDetail" />
      <div class="files">
        <FileItem :fileInfo="file" v-for="file in fileList" @showQRCode="showQRCode" @showDetail="showDetail" />
      </div>
    </main>
    <footer>
      FinelyGe © 2022 ~ {{ new Date().getFullYear() }}
    </footer>
  </div>
</template>

<style scoped lang="scss">
@import "../styles/colors.scss";

.base {
  min-height: 100vh;
  background-color: $background-light;

  &.large {
    display: grid;
    grid-template-rows: 80px 1fr 30px;
  }

  header {
    display: grid;
    grid-template-columns: 1fr 100px;
    padding-inline: 10px;
    align-items: center;
    background-color: $background-dark;

    .title {
      background-color: $background-light;
      color: $text-dark;
      font-size: 2rem;
      text-align: center;
      padding: 10px;
      width: fit-content;
      border-radius: 1rem;
    }
  }

  main {
    height: 200px;
  }

  footer {
    text-align: center;
  }
}
</style>
