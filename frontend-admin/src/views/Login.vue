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
import { ElMessage } from 'element-plus'

const router = useRouter()
const form = reactive({ username: '', password: '' })

const login = () => {
  if (!form.username || !form.password) return ElMessage.warning('请输入账号密码')
  // 模拟登录成功，写入 token
  localStorage.setItem('admin_token', 'fake-token-123')
  ElMessage.success('登录成功')
  router.push('/dashboard')
}
</script>

<style scoped>
.login-container { height: 100vh; display: flex; justify-content: center; align-items: center; background-color: #2d3a4b; }
.login-card { width: 400px; }
</style>
