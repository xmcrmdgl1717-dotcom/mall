<template>
  <el-card>
    <div style="margin-bottom: 20px; display: flex; justify-content: space-between;">
      <h3>买家用户管理</h3>
      <el-button type="primary" @click="fetchUsers">刷新列表</el-button>
    </div>
    
    <el-table :data="tableData" v-loading="loading" border style="width: 100%">
      <el-table-column prop="id" label="买家ID" width="100" />
      <el-table-column prop="username" label="用户名" />
      <el-table-column prop="phone" label="手机号" />
      <el-table-column label="状态" width="120">
        <template #default="scope">
          <el-tag :type="scope.row.status === 1 ? 'success' : 'danger'">
            {{ scope.row.status === 1 ? '正常' : '已封禁' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="created_at" label="注册时间" width="200">
         <template #default="scope">
           <!-- 防止空数据报错 -->
           {{ scope.row.created_at ? new Date(scope.row.created_at).toLocaleString() : '-' }}
         </template>
      </el-table-column>
      <el-table-column label="操作" width="150">
        <template #default>
          <el-button size="small" type="danger" link>封禁账号</el-button>
        </template>
      </el-table-column>
    </el-table>
  </el-card>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { ElMessage } from 'element-plus'

const loading = ref(false)
const tableData = ref([])
// 确保这里的环境变量在 Render 配置了，或者用默认值
const API_BASE = import.meta.env.VITE_API_BASE || 'https://mall-api-13mx.onrender.com/api/v1' 

const fetchUsers = async () => {
  loading.value = true
  try {
    // 获取 token
    const token = localStorage.getItem('admin_token')
    
    // 发起请求，带上 Authorization 头
    const res = await axios.get(`${API_BASE}/admin/users`, {
      headers: {
        Authorization: `Bearer ${token}`
      }
    })
    
    // 如果后端返回的数据在 data 字段里
    tableData.value = res.data.data || []
    
    // 如果数据是空的，给个友好提示
    if (tableData.value.length === 0) {
      ElMessage.info('目前还没有买家注册')
    }
  } catch (error) {
    console.error('请求失败详情:', error)
    if (error.response && error.response.status === 401) {
      ElMessage.error('登录已过期，请重新登录')
    } else {
      ElMessage.error(error.response?.data?.msg || '获取用户失败，请检查网络')
    }
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchUsers()
})
</script>
