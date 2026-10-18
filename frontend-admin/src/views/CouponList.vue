<template>
  <el-card>
    <div style="margin-bottom: 20px; display: flex; justify-content: space-between;">
      <h3>优惠券管理</h3>
      <div>
        <el-button type="primary" @click="dialogVisible = true">+ 发布优惠券</el-button>
        <el-button @click="fetchData">刷新</el-button>
      </div>
    </div>

    <el-table :data="tableData" v-loading="loading" border style="width: 100%">
      <el-table-column prop="id" label="ID" width="80" />
      <el-table-column prop="name" label="优惠券名称" />
      <el-table-column prop="amount" label="抵扣金额" width="100">
        <template #default="scope">$ {{ scope.row.amount }}</template>
      </el-table-column>
      <el-table-column prop="min_amount" label="使用门槛" width="120">
        <template #default="scope">满 ${{ scope.row.min_amount }} 可用</template>
      </el-table-column>
      <el-table-column label="发行/剩余" width="120">
        <template #default="scope">{{ scope.row.total_count }} / {{ scope.row.left_count }}</template>
      </el-table-column>
      <el-table-column label="有效期至" width="180">
        <template #default="scope">{{ new Date(scope.row.end_time).toLocaleDateString() }}</template>
      </el-table-column>
    </el-table>

    <el-dialog v-model="dialogVisible" title="发布优惠券" width="500px">
      <el-form :model="form" label-width="100px">
        <el-form-item label="券名称" required>
          <el-input v-model="form.name" placeholder="如：新人专享券" />
        </el-form-item>
        <el-form-item label="抵扣金额" required>
          <el-input-number v-model="form.amount" :min="1" :precision="2" />
        </el-form-item>
        <el-form-item label="满减门槛">
          <el-input-number v-model="form.min_amount" :min="0" :precision="2" />
        </el-form-item>
        <el-form-item label="发行量" required>
          <el-input-number v-model="form.total_count" :min="1" />
        </el-form-item>
        <el-form-item label="有效天数" required>
          <el-input-number v-model="form.days" :min="1" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="submitForm">发布</el-button>
      </template>
    </el-dialog>
  </el-card>
</template>

<script setup>
import { ref, onMounted, reactive } from 'vue'
import axios from 'axios'
import { ElMessage } from 'element-plus'

const API_BASE = import.meta.env.VITE_API_BASE || 'https://mall-api-l3mx.onrender.com/api/v1'
const loading = ref(false)
const tableData = ref([])
const dialogVisible = ref(false)

const form = reactive({ name: '', amount: 5, min_amount: 50, total_count: 1000, days: 30 })

const fetchData = async () => {
  loading.value = true
  try {
    const res = await axios.get(`${API_BASE}/admin/coupons`)
    tableData.value = res.data.data || []
  } catch (error) {
    ElMessage.error('获取失败')
  } finally {
    loading.value = false
  }
}

const submitForm = async () => {
  if (!form.name) return ElMessage.warning('请输入名称')
  try {
    await axios.post(`${API_BASE}/admin/coupons`, form)
    ElMessage.success('发布成功')
    dialogVisible.value = false
    fetchData()
  } catch (error) {
    ElMessage.error('发布失败')
  }
}

onMounted(() => { fetchData() })
</script>
