import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login.vue')
  },
  {
    path: '/',
    component: () => import('../components/Layout.vue'),
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('../views/Dashboard.vue')
      },
      {
        path: 'products',
        name: 'Products',
        component: () => import('../views/ProductList.vue')
      },
      {
        path: 'users', // 新增：买家管理
        name: 'Users',
        component: () => import('../views/UserList.vue')
      },
      {
        path: 'orders', // 新增：订单管理 (后续做)
        name: 'Orders',
        component: () => import('../views/OrderList.vue')
      }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('admin_token')
  if (to.name !== 'Login' && !token) {
    next({ name: 'Login' })
  } else {
    next()
  }
})

export default router
