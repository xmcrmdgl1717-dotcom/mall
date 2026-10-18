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
      { path: 'dashboard', name: 'Dashboard', component: () => import('../views/Dashboard.vue') },
      { path: 'products', name: 'Products', component: () => import('../views/ProductList.vue') },
      { path: 'categories', name: 'Categories', component: () => import('../views/CategoryList.vue') },
      { path: 'orders', name: 'Orders', component: () => import('../views/OrderList.vue') },
      { path: 'users', name: 'Users', component: () => import('../views/UserList.vue') },
      { path: 'merchants', name: 'Merchants', component: () => import('../views/MerchantList.vue') },
      { path: 'banners', name: 'Banners', component: () => import('../views/BannerList.vue') },
      { path: 'roles', name: 'Roles', component: () => import('../views/RoleList.vue') },
      // 新增优惠券路由
      { path: 'coupons', name: 'Coupons', component: () => import('../views/CouponList.vue') }
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
