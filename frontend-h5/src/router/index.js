import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home
      
    }
 
    {
  path: '/product/:id',
  name: 'ProductDetail',
  component: () => import('../views/ProductDetail.vue')
    } 
  ]
})

// 🌟 必须有下面这一句，才能被 main.js 正确引入
export default router

