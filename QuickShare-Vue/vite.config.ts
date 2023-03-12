import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import {AntDesignVueResolver} from 'unplugin-vue-components/resolvers'
import Components from 'unplugin-vue-components/vite'
Components({
  resolvers: [AntDesignVueResolver()],
})
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()]
})
