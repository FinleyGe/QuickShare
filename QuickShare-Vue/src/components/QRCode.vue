<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue';
import qrcode from 'qrcode-generator';
import { BaseUrl } from '../const';
const QRCodeIMG = ref<HTMLImageElement>();

const props = defineProps<{
  path: string
}>();

const emits = defineEmits(["close"]);

const path = computed<string>(() => BaseUrl + "get/" + props.path);

function generateQRCode() {
  const qr = qrcode(0, 'L');
  // qr.addData(BaseUrl + "get/" + props.path);
  qr.addData(path.value);
  qr.make();
  QRCodeIMG.value!.src = qr.createDataURL(4);
}

function close(this: any) {
  emits.call(this, "close");
}

function copy() {
  navigator.clipboard.writeText(path.value);
}

watch(() => props.path, () => {
  generateQRCode();
})
</script>
<template>
  <div class="popup">
    <img class="qrcode" ref="QRCodeIMG" @click="close" />
    <div class="path">
      {{ path }}
    </div>
    <button class="btn-copy" @click="copy">Copy</button>
    <button class="btn-close" @click="close">Close</button>
  </div>
</template>
<style scoped lang="scss">
.popup {
  z-index: 10;
  height: 100vh;
  width: 100vw;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  position: fixed;
  top: 0;
  left: 0;

  .qrcode {
    width: 400px;
    z-index: 11;
  }

  .path {
    margin-block: 20px;
    font-size: 20px;
    color: black;
    background-color: white;
    padding: 5px;
  }
}
</style>
