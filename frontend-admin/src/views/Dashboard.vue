<template>
  <div class="dashboard">
    <el-row :gutter="20">
      <el-col :span="6">
        <el-card shadow="hover" class="stat-card">
          <div class="title">总销售额 (GMV)</div>
          <div class="value" style="color: #f56c6c;">$ {{ stats.total_sales.toFixed(2) }}</div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover" class="stat-card">
          <div class="title">累计订单量</div>
          <div class="value">{{ stats.order_count }}</div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover" class="stat-card">
          <div class="title">注册买家数</div>
          <div class="value">{{ stats.user_count }}</div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover" class="stat-card">
          <div class="title">在售商品数</div>
          <div class="value">{{ stats.product_count }}</div>
        </el-card>
      </el-col>
    </el-row>

    <el-card style="margin-top: 20px;">
      <h3>系统提示</h3>
      <p style="color: #666;">欢迎使用 FRESHCO Admin 商城管理系统。左侧菜单可管理商品、订单、类目与用户。</p>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted, reactive } from 'vue'
import axios from 'axios'
import { ElMessage } from 'element-plus'

const API_BASE = import.meta.env.VITE_API_BASE || 'https://mall-api-l3mx.onrender.com/api/v1'
const stats = reactive({ total_sales: 0, order_count: 0, user_count: 0, product_count: 0 })

onMounted(async () => {
  try {
    const res = await axios.get(`${API_BASE}/admin/dashboard`)
    Object.assign(stats, res.data.data)
  } catch (error) {
    ElMessage.error('获取数据失败')
  }
})
</script>

<style scoped>
.stat-card { text-align: center; padding: 20px 0; }
.stat-card .title { font-size: 14px; color: #999; margin-bottom: 15px; }
.stat-card .value { font-size: 28px; font-weight: bold; color: #333; }
</style>
