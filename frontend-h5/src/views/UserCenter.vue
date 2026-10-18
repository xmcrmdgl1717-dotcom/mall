<template>
  <div class="user-page">
    <van-nav-bar title="My Account" fixed placeholder />

    <div class="user-header">
      <div class="avatar"></div>
      <div class="info" v-if="isLogin">
        <h2>{{ username }}</h2>
        <p>欢迎回来</p>
      </div>
      <div class="info" v-else>
        <h2>未登录</h2>
        <p @click="$router.push('/login')" style="color:#1989fa; cursor:pointer;">点击登录 / 注册 ></p>
      </div>
    </div>

    <div class="menu-list">
      <van-cell title="我的订单" is-link to="/orders" icon="orders-o" />
      <van-cell title="收货地址" is-link icon="location-o" />
      <van-cell title="联系客服" is-link icon="service-o" />
      <van-cell title="设置" is-link icon="setting-o" />
    </div>

    <div style="padding: 20px;" v-if="isLogin">
      <van-button block type="danger" plain round @click="logout">退出登录</van-button>
    </div>

    <van-tabbar v-model="active" route>
      <van-tabbar-item replace to="/" icon="home-o">Home</van-tabbar-item>
      <van-tabbar-item replace to="/products" icon="search">Search</van-tabbar-item>
      <van-tabbar-item replace to="/cart" icon="cart-o">Cart</van-tabbar-item>
      <van-tabbar-item replace to="/user" icon="user-o">Me</van-tabbar-item>
    </van-tabbar>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const isLogin = ref(false)
const username = ref('')
const active = ref(3)

onMounted(() => {
  const token = localStorage.getItem('buyer_token')
  if (token) {
    isLogin.value = true
    username.value = localStorage.getItem('buyer_name') || 'User'
  }
})

const logout = () => {
  localStorage.removeItem('buyer_token')
  localStorage.removeItem('buyer_name')
  isLogin.value = false
  router.push('/')
}
</script>

<style scoped>
.user-page { min-height: 100vh; background: #f7f8fa; padding-bottom: 50px; }
.user-header { background: #fff; padding: 30px 20px; display: flex; align-items: center; gap: 15px; }
.avatar { width: 60px; height: 60px; background: #ddd; border-radius: 50%; }
.info h2 { margin: 0 0 5px; font-size: 18px; }
.info p { margin: 0; font-size: 13px; color: #999; }
.menu-list { margin-top: 10px; }
</style>
