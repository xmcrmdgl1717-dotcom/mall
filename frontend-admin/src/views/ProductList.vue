<template>
  <el-card>
    <div style="margin-bottom: 20px; display: flex; justify-content: space-between;">
      <el-button type="primary" @click="openDialog">+ 新增商品</el-button>
      <el-button @click="fetchData">刷新列表</el-button>
    </div>

    <el-table :data="tableData" v-loading="loading" border style="width: 100%">
      <el-table-column prop="id" label="ID" width="80" />
      <el-table-column label="图片" width="100">
        <template #default="scope">
          <el-image :src="scope.row.main_image" style="width: 50px; height: 50px; border-radius: 4px;" fit="cover" />
        </template>
      </el-table-column>
      <el-table-column prop="title" label="商品名称" show-overflow-tooltip />
      <el-table-column prop="min_price" label="价格" width="120">
        <template #default="scope">$ {{ scope.row.min_price }}</template>
      </el-table-column>
      <el-table-column prop="total_sales" label="销量" width="100" />
      <el-table-column label="状态" width="100">
        <template #default="scope">
          <el-tag :type="scope.row.status === 1 ? 'success' : 'info'">
            {{ scope.row.status === 1 ? '上架中' : '已下架' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="150" fixed="right">
        <template #default>
          <el-button size="small" type="primary" link>编辑</el-button>
          <el-button size="small" type="danger" link>下架</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 新增商品弹窗 -->
    <el-dialog v-model="dialogVisible" title="新增商品" width="500px">
      <el-form :model="form" label-width="80px">
        <el-form-item label="商品标题">
          <el-input v-model="form.title" placeholder="请输入商品标题" />
        </el-form-item>
        <el-form-item label="主图链接">
          <el-input v-model="form.main_image" placeholder="输入图片URL" />
        </el-form-item>
        <el-form-item label="售价">
          <el-input-number v-model="form.price" :precision="2" :step="1" />
        </el-form-item>
        <el-form-item label="库存">
          <el-input-number v-model="form.stock" :step="10" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="submitForm">确定发布</el-button>
        </span>
      </template>
    </el-dialog>
  </el-card>
</template>

<script setup>
import { ref, onMounted, reactive } from 'vue'
import axios from 'axios'
import { ElMessage } from 'element-plus'

const loading = ref(false)
const tableData = ref([])
const dialogVisible = ref(false)

const form = reactive({
  title: '',
  main_image: '',
  price: 0,
  stock: 100
})

const API_BASE = import.meta.env.VITE_API_BASE || 'http://localhost:8080/api/v1'

const fetchData = async () => {
  loading.value = true
  try {
    const res = await axios.get(`${API_BASE}/products`)
    tableData.value = res.data.data.list || []
  } catch (error) {
    ElMessage.error('获取数据失败')
  } finally {
    loading.value = false
  }
}

const openDialog = () => {
  form.title = ''
  form.main_image = 'https://picsum.photos/200'
  form.price = 0
  form.stock = 100
  dialogVisible.value = true
}

const submitForm = async () => {
  if (!form.title) return ElMessage.warning('请输入商品标题')
  try {
    await axios.post(`${API_BASE}/products`, {
      title: form.title,
      main_image: form.main_image,
      min_price: form.price
    })
    ElMessage.success('商品发布成功！')
    dialogVisible.value = false
    fetchData() // 刷新表格
  } catch (error) {
    ElMessage.error('发布失败，请检查后端')
  }
}

  
onMounted(() => {
  fetchData()
})
</script>
