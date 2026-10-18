<template>
  <div class="login-page">
    <van-nav-bar title="Login / Register" left-arrow @click-left="$router.back()" fixed placeholder />
    
    <div class="header-logo">FRESHCO</div>

    <van-form @submit="onSubmit">
      <van-cell-group inset>
        <van-field v-model="username" name="username" label="Username" placeholder="请输入用户名" :rules="[{ required: true, message: '请填写用户名' }]" />
        <van-field v-model="password" type="password" name="password" label="Password" placeholder="请输入密码" :rules="[{ required: true, message: '请填写密码' }]" />
      </van-cell-group>

      <div style="margin: 16px;">
        <van-button round block type="primary" native-type="submit" :loading="loading">
          {{ isRegister ? '注册新账号' : '登录' }}
        </van-button>
      </div>
    </van-form>

    <div class="switch-mode" @click="isRegister = !isRegister">
      {{ isRegister ? '已有账号？去登录' : '没有账号？去注册' }}
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'
import { showSuccessToast, showFailToast } from 'vant'

const router = useRouter()
const API_BASE = import.meta.env.VITE_API_BASE || 'https://mall-api-l3mx.onrender.com/api/v1'

const username = ref('')
const password = ref('')
const isRegister = ref(false)
const loading = ref(false)

const onSubmit = async () => {
  loading.value = true
  try {
    if (isRegister.value) {
      await axios.post(`${API_BASE}/buyer/register`, { username: username.value, password: password.value })
      showSuccessToast('注册成功，请登录')
      isRegister.value = false
    } else {
      const res = await axios.post(`${API_BASE}/buyer/login`, { username: username.value, password: password.value })
      localStorage.setItem('buyer_token', res.data.data.token)
      localStorage.setItem('buyer_name', res.data.data.username)
      showSuccessToast('登录成功')
      router.push('/user')
    }
  } catch (error) {
    showFailToast(error.response?.data?.msg || '操作失败')
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-page { min-height: 100vh; background: #f5f5f5; }
.header-logo { text-align: center; font-size: 32px; font-weight: 900; color: #333; padding: 40px 0; letter-spacing: 2px; }
.switch-mode { text-align: center; font-size: 14px; color: #1989fa; margin-top: 20px; }
</style>
