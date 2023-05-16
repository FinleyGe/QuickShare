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
import ShareCode from '../components/ShareCode.vue'

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

function logout() {
  if (store.isLogin) {
    store.setLogin(false);
    // delete all cookie

    document.cookie.split(";").forEach((c) => {
      document.cookie = c
        .replace(/^ +/, "")
        .replace(/=.*/, "=;expires=" + new Date().toUTCString() + ";path=/");
    });
  }
}
</script>

<template>
  <div class="base" :class="store.deviceSize">
    <header>
      <div class="title">
        QuickShare
      </div>
      <div class="button-set">
        <button @click="handleLogin">{{ showLogin ? '返回' : '登陆' }}</button>
        <button @click="logout" v-if="store.isLogin">登出</button>
      </div>
    </header>
    <main>
      <Upload v-show="!showLogin" />
      <Login v-show="showLogin" @logined="Logined" />
      <QRCode @close="isShowQRCode = false" :path="QRCodePath" v-show="isShowQRCode" />
      <FileItemDetail @close="isShowDetail = false" :file-info="fileDetail" v-show="isShowDetail" />
      <ShareCode />
    </main>
    <aside>
      <div class="title">
        文件列表
        <button @click="refresh">最近</button>
        <button> 图片 </button>
      </div>
      <div class="files">
        <FileItem :fileInfo="file" v-for="file in fileList" @showQRCode="showQRCode" @showDetail="showDetail" />
      </div>
    </aside>
    <footer>
      QuickShare | FinelyGe © 2022 ~ {{ new Date().getFullYear() }}
    </footer>
  </div>
</template>

<style scoped lang="scss">
@import "../styles/colors.scss";

.base {
  min-height: 100vh;
  background-color: $background-light;

  display: grid;

  &.large {
    grid-template-areas: "header header header" "main main aside" "footer footer footer";
    grid-row-gap: 10px;
    grid-column-gap: 10px;
    grid-template-rows: 80px 1fr 50px;
    grid-template-columns: 1fr 1fr 20%;
  }

  &.medium,
  &.small {
    grid-template-areas: "header" "main" "aside" "footer";
    grid-row-gap: 20px;
    grid-template-rows: 80px 1fr 1fr 50px;
  }

  header {
    grid-area: header;
    display: grid;
    grid-template-columns: 1fr 100px;
    padding-inline: 10px;
    align-items: center;
    background-color: $background-dark;
    border-radius: 1rem;

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
    grid-area: main;
    background-color: $background-medium;
    border-radius: 1rem;
  }

  aside {
    grid-area: aside;

    border-radius: 1rem;
    background-color: $background-medium;

    .title {
      margin-top: 10px;
    }

    .files {
      height: 100vh;
      overflow-y: auto;
    }
  }

  footer {
    background-color: $background-medium;
    grid-area: footer;
    display: flex;
    align-items: center;
    justify-content: center;
  }
}
</style>
