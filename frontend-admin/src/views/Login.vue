<template>
  <div class="login-container">
    <el-card class="login-card">
      <h2 style="text-align: center; margin-bottom: 20px;">后台登录</h2>
      <el-form :model="form" label-width="0">
        <el-form-item>
          <el-input v-model="form.username" placeholder="用户名 (任意输入)" />
        </el-form-item>
        <el-form-item>
          <el-input v-model="form.password" type="password" placeholder="密码 (任意输入)" />
        </el-form-item>
        <el-button type="primary" style="width: 100%;" @click="login">登录</el-button>
      </el-form>
    </el-card>
  </div>
</template>

<script setup>
import { reactive } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'
import { ElMessage } from 'element-plus'

const router = useRouter()
const form = reactive({ username: 'admin', password: 'admin123' }) // 默认填好方便测试

const API_BASE = import.meta.env.VITE_API_BASE || 'http://localhost:8080/api/v1'

const login = async () => {
  if (!form.username || !form.password) return ElMessage.warning('请输入账号密码')
  try {
    const res = await axios.post(`${API_BASE}/admin/login`, form)
    localStorage.setItem('admin_token', res.data.data.token)
    ElMessage.success('登录成功')
    router.push('/dashboard')
  } catch (error) {
    ElMessage.error(error.response?.data?.msg || '登录失败')
  }
}
</script>

<style scoped>
.login-container { height: 100vh; display: flex; justify-content: center; align-items: center; background-color: #2d3a4b; }
.login-card { width: 400px; }
</style>
