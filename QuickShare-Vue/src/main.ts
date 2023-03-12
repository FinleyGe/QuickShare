import { createApp } from 'vue'
import App from './App.vue'
import router from './router/router'
import { pinia } from './store/init'

createApp(App)
  .use(router)
  .use(pinia)
  .mount('#app')
