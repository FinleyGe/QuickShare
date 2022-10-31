<script setup lang="ts">
import {message} from 'ant-design-vue'
import {UploadOutlined} from '@ant-design/icons-vue'
import {Button as AButton} from 'ant-design-vue'
import {ref} from 'vue'
import {api, UploadFile} from '../api'
import {BaseUrl} from '../const'
import qrcode from '@chenfengyuan/vue-qrcode'
const file = ref<File>();
const filePath = ref<string>('');
function handleUpload(){
//upload a file
UploadFile(file.value!).then(_res=>{
message.success('upload success')
console.log(_res)
filePath.value = BaseUrl + 'download/' + _res.data.data.hash
}).catch(_err=>{
message.error('upload failed')
})

}

function handleFileChange (e: Event) {
  file.value = (e.target as HTMLInputElement)?.files?.[0]
}
</script>

<template>
{{file}}
{{filePath}}

<qrcode :value="filePath"></qrcode>
<form>
<input
type="file"
v-on:change="handleFileChange"
>
<a-button
@click="handleUpload">
<upload-outlined></upload-outlined>
Click To Upload
</a-button>
</form>
</template>

<style scoped>
</style>
