<template>
  <el-card>
    <div style="margin-bottom: 20px; display: flex; justify-content: space-between;">
      <h3>订单管理</h3>
      <el-button type="primary" @click="fetchOrders">刷新列表</el-button>
    </div>

    <el-table :data="tableData" v-loading="loading" border style="width: 100%">
      <el-table-column prop="order_no" label="订单号" width="180" />
      <el-table-column label="收货信息" width="250">
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
        <template #default="scope"><span style="color: #f56c6c; font-weight: bold;">$ {{ scope.row.total_amount }}</span></template>
      </el-table-column>
      <el-table-column label="状态" width="100">
        <template #default="scope">
          <el-tag :type="scope.row.status === 1 ? 'success' : 'warning'">
            {{ scope.row.status === 1 ? '已付款' : '待处理' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="120">
        <template #default>
          <el-button size="small" type="primary" link>发货</el-button>
        </template>
      </el-table-column>
    </el-table>
    
    <el-empty v-if="tableData.length === 0 && !loading" description="暂无订单数据" />
  </el-card>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { ElMessage } from 'element-plus'

const loading = ref(false)
const tableData = ref([])
const API_BASE = import.meta.env.VITE_API_BASE || 'http://localhost:8080/api/v1'

const fetchOrders = async () => {
  loading.value = true
  try {
    // 这里调用的是后台获取全部订单的接口 (注意：你之前后端写的获取订单列表可能默认查的是 user_id=1 的订单，这里实际开发需要写一个新的后台查全量订单接口，这里先占位)
    const res = await axios.get(`${API_BASE}/orders`) 
    tableData.value = res.data.data || []
  } catch (error) {
    ElMessage.error('获取订单失败')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchOrders()
})
</script>
