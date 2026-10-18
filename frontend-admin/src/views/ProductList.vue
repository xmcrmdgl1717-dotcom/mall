<template>
  <el-card>
    <div style="margin-bottom: 20px; display: flex; justify-content: space-between;">
      <h3>商品管理</h3>
      <div>
        <el-button type="primary" @click="openDialog">+ 发布新商品 (带SKU)</el-button>
        <el-button @click="fetchData">刷新列表</el-button>
      </div>
    </div>

    <el-table :data="tableData" v-loading="loading" border style="width: 100%">
      <el-table-column prop="id" label="ID" width="80" />
      <el-table-column label="主图" width="100">
        <template #default="scope">
          <el-image :src="scope.row.main_image" style="width: 50px; height: 50px;" fit="cover" />
        </template>
      </el-table-column>
      <el-table-column prop="title" label="商品名称" show-overflow-tooltip />
      <el-table-column prop="min_price" label="最低售价" width="120">
        <template #default="scope"><span style="color:#f56c6c">$ {{ scope.row.min_price }}</span></template>
      </el-table-column>
      <el-table-column prop="total_sales" label="销量" width="100" />
      <el-table-column label="状态" width="100">
        <template #default="scope">
          <el-tag :type="scope.row.status === 1 ? 'success' : 'info'">{{ scope.row.status === 1 ? '上架中' : '已下架' }}</el-tag>
        </template>
      </el-table-column>
    </el-table>

    <!-- 复杂的发布商品弹窗 -->
    <el-dialog v-model="dialogVisible" title="发布新商品" width="700px">
      <el-form :model="form" label-width="100px">
        <el-form-item label="商品标题" required>
          <el-input v-model="form.title" placeholder="请输入商品标题" />
        </el-form-item>
        <el-form-item label="主图链接">
          <el-input v-model="form.main_image" placeholder="https://..." />
        </el-form-item>
        
        <el-divider content-position="left">商品规格 (SKU)</el-divider>
        
        <!-- 动态 SKU 列表 -->
        <div v-for="(sku, index) in form.skus" :key="index" class="sku-row">
          <el-input v-model="sku.specs" placeholder="规格 (如: 红色 / 尺寸M)" style="width: 200px; margin-right: 10px;" />
          <el-input-number v-model="sku.price" :min="0" :precision="2" placeholder="价格" style="width: 150px; margin-right: 10px;" />
          <el-input-number v-model="sku.stock" :min="0" placeholder="库存" style="width: 150px; margin-right: 10px;" />
          <el-button type="danger" icon="Delete" circle @click="removeSku(index)" />
        </div>
        
        <el-button type="primary" plain style="width: 100%; margin-top: 10px;" @click="addSku">+ 添加一个规格</el-button>
      </el-form>
      
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="submitting" @click="submitForm">立即发布</el-button>
      </template>
    </el-dialog>
  </el-card>
</template>

<script setup>
import { ref, onMounted, reactive } from 'vue'
import axios from 'axios'
import { ElMessage } from 'element-plus'

const loading = ref(false)
const submitting = ref(false)
const tableData = ref([])
const dialogVisible = ref(false)
const API_BASE = import.meta.env.VITE_API_BASE

const form = reactive({
  title: '',
  main_image: 'https://picsum.photos/400',
  skus: [{ specs: '默认规格', price: 9.99, stock: 100 }] // 默认给一条
})

const addSku = () => {
  form.skus.push({ specs: '', price: 0, stock: 0 })
}
const removeSku = (index) => {
  form.skus.splice(index, 1)
}

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
  form.main_image = 'https://picsum.photos/400'
  form.skus = [{ specs: '默认规格', price: 9.99, stock: 100 }]
  dialogVisible.value = true
}

const submitForm = async () => {
  if (!form.title) return ElMessage.warning('请输入标题')
  if (form.skus.length === 0) return ElMessage.warning('至少添加一个规格')
  
  submitting.value = true
  try {
    await axios.post(`${API_BASE}/products`, form)
    ElMessage.success('发布成功！')
    dialogVisible.value = false
    fetchData()
  } catch (error) {
    ElMessage.error('发布失败')
  } finally {
    submitting.value = false
  }
}

onMounted(() => { fetchData() })
</script>

<style scoped>
.sku-row { display: flex; align-items: center; margin-bottom: 10px; }
</style>
