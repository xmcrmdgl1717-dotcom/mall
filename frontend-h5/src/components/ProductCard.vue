<template>
  <el-card>
    <div style="margin-bottom: 20px;">
      <el-button type="primary" @click="fetchData">刷新数据</el-button>
    </div>
    <el-table :data="tableData" v-loading="loading" border style="width: 100%">
      <el-table-column prop="id" label="ID" width="80" />
      <el-table-column prop="title" label="商品名称" show-overflow-tooltip />
      <el-table-column prop="min_price" label="价格" width="120">
        <template #default="scope">$ {{ scope.row.min_price }}</template>
      </el-table-column>
      <el-table-column prop="total_sales" label="销量" width="120" />
      <el-table-column label="操作" width="150">
        <template #default>
          <el-button size="small" type="primary" link>编辑</el-button>
          <el-button size="small" type="danger" link>下架</el-button>
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

// 注意：这里需要换成你 Render 部署的 Go 后端地址，如果本地测试可以填 localhost:8080
const API_BASE = import.meta.env.VITE_API_BASE || 'http://localhost:8080/api/v1'

const fetchData = async () => {
  loading.value = true
  try {
    const res = await axios.get(`${API_BASE}/products`)
    tableData.value = res.data.data.list || []
  } catch (error) {
    ElMessage.error('获取数据失败，请检查后端是否运行')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchData()
})
</script>
