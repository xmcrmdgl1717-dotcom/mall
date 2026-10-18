<template>
  <el-card>
    <div style="margin-bottom: 20px; display: flex; justify-content: space-between;">
      <h3>商家/店铺管理</h3>
      <el-button type="primary" @click="fetchData">刷新列表</el-button>
    </div>

    <el-table :data="tableData" v-loading="loading" border style="width: 100%">
      <el-table-column prop="id" label="商家ID" width="80" />
      <el-table-column prop="name" label="店铺名称" />
      <el-table-column prop="total_sales" label="总销量" width="120" />
      <el-table-column prop="good_rate" label="好评率" width="120">
        <template #default="scope">{{ scope.row.good_rate }}%</template>
      </el-table-column>
      <el-table-column label="状态" width="120">
        <template #default="scope">
          <el-tag :type="getStatusType(scope.row.status)">{{ getStatusText(scope.row.status) }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="200">
        <template #default="scope">
          <el-button v-if="scope.row.status === 0" size="small" type="success" @click="updateStatus(scope.row.id, 1)">通过审核</el-button>
          <el-button v-if="scope.row.status === 1" size="small" type="danger" @click="updateStatus(scope.row.id, 3)">冻结店铺</el-button>
          <el-button v-if="scope.row.status === 3" size="small" type="warning" @click="updateStatus(scope.row.id, 1)">解封</el-button>
        </template>
      </el-table-column>
    </el-table>
  </el-card>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { ElMessage } from 'element-plus'

const API_BASE = import.meta.env.VITE_API_BASE || 'https://mall-api-l3mx.onrender.com/api/v1'
const loading = ref(false)
const tableData = ref([])

const getStatusText = (status) => {
  const map = { 0: '待审核', 1: '正常营业', 2: '已驳回', 3: '已冻结' }
  return map[status] || '未知'
}
const getStatusType = (status) => {
  const map = { 0: 'warning', 1: 'success', 2: 'danger', 3: 'info' }
  return map[status] || 'info'
}

const fetchData = async () => {
  loading.value = true
  try {
    const res = await axios.get(`${API_BASE}/admin/merchants`)
    tableData.value = res.data.data || []
  } catch (error) {
    ElMessage.error('获取失败')
  } finally {
    loading.value = false
  }
}

const updateStatus = async (id, status) => {
  try {
    await axios.put(`${API_BASE}/admin/merchants/${id}/status`, { status })
    ElMessage.success('状态更新成功')
    fetchData()
  } catch (error) {
    ElMessage.error('更新失败')
  }
}

onMounted(() => { fetchData() })
</script>
