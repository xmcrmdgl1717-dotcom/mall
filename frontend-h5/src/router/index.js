import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home
    },
    {
      path: '/product/:id',
      name: 'ProductDetail',
      component: () => import('../views/ProductDetail.vue')
    },
    {
      path: '/products',
      name: 'ProductList',
      component: () => import('../views/ProductList.vue')
    },
    {
      path: '/cart',
      name: 'Cart',
      component: () => import('../views/Cart.vue')
    },
    {
      path: '/checkout',
      name: 'Checkout',
      component: () => import('../views/Checkout.vue')
    },
    {
      path: '/orders',
      name: 'Orders',
      component: () => import('../views/Orders.vue')
    },
    {
      path: '/store/:id',
      name: 'StoreDetail',
      component: () => import('../views/StoreDetail.vue')
    },
    // 👇 新增：个人中心与地址管理
    {
      path: '/user',
      name: 'UserCenter',
      component: () => import('../views/UserCenter.vue')
    },
    {
      path: '/addresses',
      name: 'AddressList',
      component: () => import('../views/AddressList.vue')
    }
  ]
})

export default router
