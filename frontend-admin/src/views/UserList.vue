<template>
  <el-card>
    <div style="margin-bottom: 20px;">
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
         <template #default="scope">{{ new Date(scope.row.created_at).toLocaleString() }}</template>
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
const API_BASE = import.meta.env.VITE_API_BASE || 'http://localhost:8080/api/v1'

const fetchUsers = async () => {
  loading.value = true
  try {
    const res = await axios.get(`${API_BASE}/admin/users`)
    tableData.value = res.data.data || []
  } catch (error) {
    ElMessage.error('获取用户失败')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchUsers()
})
</script>
