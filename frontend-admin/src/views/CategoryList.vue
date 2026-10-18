<template>
  <el-card>
    <div style="margin-bottom: 20px; display: flex; justify-content: space-between;">
      <h3>类目管理</h3>
      <div>
        <el-input v-model="newCatName" placeholder="新类目名称" style="width: 200px; margin-right: 10px;" />
        <el-button type="primary" @click="addCategory">添加类目</el-button>
      </div>
    </div>
    
    <el-table :data="tableData" v-loading="loading" border style="width: 100%">
      <el-table-column prop="id" label="类目ID" width="100" />
      <el-table-column prop="name" label="类目名称" />
      <el-table-column prop="sort" label="排序权重" width="120" />
      <el-table-column label="操作" width="150">
        <template #default="scope">
          <el-button size="small" type="danger" link @click="deleteCategory(scope.row.id)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </el-card>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { ElMessage, ElMessageBox } from 'element-plus'

const API_BASE = import.meta.env.VITE_API_BASE || 'https://mall-api-l3mx.onrender.com/api/v1'
const loading = ref(false)
const tableData = ref([])
const newCatName = ref('')

const fetchData = async () => {
  loading.value = true
  try {
    const res = await axios.get(`${API_BASE}/categories`)
    tableData.value = res.data.data || []
  } catch (error) {
    ElMessage.error('获取失败')
  } finally {
    loading.value = false
  }
}

const addCategory = async () => {
  if (!newCatName.value) return ElMessage.warning('请输入名称')
  try {
    await axios.post(`${API_BASE}/admin/categories`, { name: newCatName.value, sort: 0 })
    ElMessage.success('添加成功')
    newCatName.value = ''
    fetchData()
  } catch (error) {
    ElMessage.error('添加失败')
  }
}

const deleteCategory = async (id) => {
  try {
    await ElMessageBox.confirm('确定删除该分类？', '警告', { type: 'warning' })
    await axios.delete(`${API_BASE}/admin/categories/${id}`)
    ElMessage.success('删除成功')
    fetchData()
  } catch (e) {}
}

onMounted(() => { fetchData() })
</script>
