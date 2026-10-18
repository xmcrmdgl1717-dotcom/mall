<template>
  <div class="login-container">
    <el-card class="login-card">
      <h2 style="text-align: center; margin-bottom: 20px;">后台登录</h2>
      <el-form :model="form" label-width="0">
        <el-form-item>
          <el-input v-model="form.username" placeholder="请输入用户名" />
        </el-form-item>
        <el-form-item>
          <el-input v-model="form.password" type="password" placeholder="请输入密码" show-password />
        </el-form-item>
        <el-button type="primary" style="width: 100%;" :loading="loading" @click="login">
          {{ loading ? '登录中(唤醒服务可能需要50秒)...' : '登录' }}
        </el-button>
      </el-form>
    </el-card>
  </div>
</template>

<script setup>
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'
import { ElMessage } from 'element-plus'

const router = useRouter()
const loading = ref(false)

// 默认填好账号密码
const form = reactive({ 
  username: 'admin', 
  password: 'admin123' 
})

// 你的真实后端地址写死在这里
const API_BASE = 'https://mall-api-l3mx.onrender.com/api/v1'

const login = async () => {
  if (!form.username || !form.password) return ElMessage.warning('请输入账号密码')
  
  loading.value = true
  try {
    // 发起登录请求
    const res = await axios.post(`${API_BASE}/admin/login`, form)
    
    // 保存后端返回的 Token
    localStorage.setItem('admin_token', res.data.data.token)
    localStorage.setItem('admin_role', res.data.data.role_id)
    
    ElMessage.success('登录成功')
    router.push('/dashboard')
  } catch (error) {
    console.error('登录报错详情:', error)
    
    // 给出更明确的错误提示
    if (error.code === 'ERR_NETWORK') {
      ElMessage.error('网络连接失败，可能是后端正在唤醒，请稍后再试')
    } else {
      ElMessage.error(error.response?.data?.msg || '账号或密码错误')
    }
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-container { height: 100vh; display: flex; justify-content: center; align-items: center; background-color: #2d3a4b; }
.login-card { width: 400px; }
</style>
