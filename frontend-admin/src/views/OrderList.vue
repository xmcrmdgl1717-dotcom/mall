<template>
  <el-card>
    <div style="margin-bottom: 20px; display: flex; justify-content: space-between;">
      <h3>订单管理</h3>
      <el-button type="primary" @click="fetchOrders">刷新列表</el-button>
    </div>

    <el-table :data="tableData" v-loading="loading" border style="width: 100%">
      <el-table-column prop="order_no" label="订单号" width="180" />
      <el-table-column label="收货人" width="200">
        <template #default="scope">
          <div>{{ scope.row.address_name }} ({{ scope.row.address_phone }})</div>
          <div style="font-size: 12px; color: #999;">{{ scope.row.address_detail }}</div>
        </template>
      </el-table-column>
      <el-table-column label="购买商品">
        <template #default="scope">
          <div v-for="item in scope.row.items" :key="item.id" style="font-size: 13px; margin-bottom: 4px;">
            {{ item.title }} ({{ item.specs }}) x {{ item.quantity }}
          </div>
        </template>
      </el-table-column>
      <el-table-column prop="total_amount" label="总价" width="120">
        <template #default="scope"><span style="color: #f56c6c;">$ {{ scope.row.total_amount }}</span></template>
      </el-table-column>
      <el-table-column label="状态" width="100">
        <template #default="scope">
          <el-tag :type="getStatusType(scope.row.status)">{{ getStatusText(scope.row.status) }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="150" fixed="right">
        <template #default="scope">
          <el-button 
            v-if="scope.row.status === 1" 
            size="small" 
            type="success" 
            @click="handleShip(scope.row.id)"
          >
            点击发货
          </el-button>
          <span v-else style="color: #999; font-size: 12px;">无需操作</span>
        </template>
      </el-table-column>
    </el-table>
  </el-card>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { ElMessage, ElMessageBox } from 'element-plus'

const loading = ref(false)
const tableData = ref([])
const API_BASE = import.meta.env.VITE_API_BASE

const fetchOrders = async () => {
  loading.value = true
  try {
    const res = await axios.get(`${API_BASE}/orders`)
    tableData.value = res.data.data || []
  } catch (error) {
    ElMessage.error('获取订单失败')
  } finally {
    loading.value = false
  }
}

// 处理发货
const handleShip = (orderId) => {
  ElMessageBox.confirm('确认该订单已打包并发出？', '发货确认', { type: 'warning' }).then(async () => {
    try {
      await axios.put(`${API_BASE}/orders/${orderId}/ship`)
      ElMessage.success('发货成功！')
      fetchOrders() // 刷新列表
    } catch (e) {
      ElMessage.error('发货失败')
    }
  }).catch(() => {})
}

const getStatusText = (status) => {
  const map = { 0: '待付款', 1: '待发货', 2: '已发货', 3: '已完成' }
  return map[status] || '未知'
}
const getStatusType = (status) => {
  const map = { 0: 'info', 1: 'warning', 2: 'success', 3: 'success' }
  return map[status] || 'info'
}

onMounted(() => { fetchOrders() })
</script>
