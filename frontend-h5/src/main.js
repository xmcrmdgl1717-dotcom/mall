import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'
import Vant from 'vant'
import 'vant/lib/index.css'
import axios from 'axios'

// 自动携带买家 Token
axios.interceptors.request.use(config => {
  const token = localStorage.getItem('buyer_token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

const app = createApp(App)
app.use(createPinia())
app.use(router)
app.use(Vant)
app.mount('#app')
