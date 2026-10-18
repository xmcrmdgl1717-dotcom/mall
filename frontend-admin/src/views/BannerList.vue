<template>
  <el-card>
    <div style="margin-bottom: 20px; display: flex; justify-content: space-between;">
      <h3>轮播图管理</h3>
      <div>
        <el-button type="primary" @click="openDialog">+ 添加轮播图</el-button>
        <el-button @click="fetchData">刷新</el-button>
      </div>
    </div>

    <el-table :data="tableData" v-loading="loading" border style="width: 100%">
      <el-table-column prop="id" label="ID" width="80" />
      <el-table-column label="预览" width="200">
        <template #default="scope">
          <el-image :src="scope.row.image_url" style="width: 150px; height: 60px; border-radius: 4px;" fit="cover" />
        </template>
      </el-table-column>
      <el-table-column prop="link_url" label="跳转链接" />
      <el-table-column prop="sort" label="排序权重" width="120" />
      <el-table-column label="操作" width="120">
        <template #default="scope">
          <el-button size="small" type="danger" link @click="deleteBanner(scope.row.id)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog v-model="dialogVisible" title="添加轮播图" width="500px">
      <el-form :model="form" label-width="100px">
        <el-form-item label="图片链接" required>
          <el-input v-model="form.image_url" placeholder="输入图片 URL (如 https://picsum.photos/800/350)" />
        </el-form-item>
        <el-form-item label="跳转链接">
          <el-input v-model="form.link_url" placeholder="如 /product/1" />
        </el-form-item>
        <el-form-item label="排序权重">
          <el-input-number v-model="form.sort" :min="1" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="submitForm">确认添加</el-button>
      </template>
    </el-dialog>
  </el-card>
</template>

<script setup>
import { ref, onMounted, reactive } from 'vue'
import axios from 'axios'
import { ElMessage, ElMessageBox } from 'element-plus'

const API_BASE = import.meta.env.VITE_API_BASE || 'https://mall-api-l3mx.onrender.com/api/v1'
const loading = ref(false)
const tableData = ref([])
const dialogVisible = ref(false)

const form = reactive({ image_url: '', link_url: '', sort: 1 })

const fetchData = async () => {
  loading.value = true
  try {
    const res = await axios.get(`${API_BASE}/admin/banners`)
    tableData.value = res.data.data || []
  } catch (error) {
    ElMessage.error('获取失败')
  } finally {
    loading.value = false
  }
}

const openDialog = () => {
  form.image_url = 'https://picsum.photos/800/350'
  form.link_url = ''
  form.sort = 1
  dialogVisible.value = true
}

const submitForm = async () => {
  if (!form.image_url) return ElMessage.warning('请输入图片链接')
  try {
    await axios.post(`${API_BASE}/admin/banners`, form)
    ElMessage.success('添加成功')
    dialogVisible.value = false
    fetchData()
  } catch (error) {
    ElMessage.error('添加失败')
  }
}

const deleteBanner = async (id) => {
  try {
    await ElMessageBox.confirm('确定删除该轮播图？', '警告', { type: 'warning' })
    await axios.delete(`${API_BASE}/admin/banners/${id}`)
    ElMessage.success('删除成功')
    fetchData()
  } catch (e) {}
}

onMounted(() => { fetchData() })
</script>
