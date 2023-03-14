<script setup lang="ts">
import { onMounted, ref, watch } from 'vue';
import qrcode from 'qrcode-generator';
import { BaseUrl } from '../const';
const QRCodeIMG = ref<HTMLImageElement>();

const props = defineProps<{
  path: string
}>();

const emits = defineEmits(["close"]);
function generateQRCode() {
  const qr = qrcode(0, 'L');
  qr.addData(BaseUrl + "get/" + props.path);
  qr.make();
  QRCodeIMG.value!.src = qr.createDataURL(4);
}

function close(this: any) {
  emits.call(this, "close");
}
watch(() => props.path, () => {
  generateQRCode();
})
</script>
<template>
  <div class="popup">
    <img class="qrcode" ref="QRCodeIMG" @click="close" />
  </div>
</template>
<style lang="scss">
.popup {
  z-index: 10;
  height: 100vh;
  width: 100vw;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  position: fixed;
  top: 0;
  left: 0;

  .qrcode {
    z-index: 11;
    width: 400px;
    margin: auto;
  }
}
</style>
