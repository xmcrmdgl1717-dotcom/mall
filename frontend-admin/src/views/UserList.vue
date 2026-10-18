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
      <el-table-column prop="created_at" label="注册时间" width="200">
         <template #default="scope">
           {{ scope.row.created_at ? new Date(scope.row.created_at).toLocaleString() : '-' }}
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

// ✅ 标准的读取环境变量写法
const API_BASE = import.meta.env.VITE_API_BASE

const fetchUsers = async () => {
  loading.value = true
  try {
    const res = await axios.get(`${API_BASE}/admin/users`)
    tableData.value = res.data.data || []
    if (tableData.value.length === 0) ElMessage.info('目前没有注册买家')
  } catch (error) {
    console.error('用户请求失败:', error)
    ElMessage.error('获取用户失败')
  } finally {
    loading.value = false
  }
}

onMounted(() => { fetchUsers() })
</script>
