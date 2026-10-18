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

    <el-dialog v-model="dialogVisible" title="添加轮播图" width="600px">
      <el-form :model="form" label-width="100px">
        
        <!-- 改造点：本地图片上传 -->
        <el-form-item label="图片上传" required>
          <el-upload
            class="upload-demo"
            :action="`${API_BASE}/admin/upload`"
            :show-file-list="false"
            :on-success="handleBannerUploadSuccess"
            :before-upload="beforeUpload"
          >
            <el-button type="primary">点击上传本地图片</el-button>
          </el-upload>
          <div style="font-size: 12px; color: #999; margin-top: 5px;">或者手动填入在线链接：</div>
          <el-input v-model="form.image_url" placeholder="https://..." style="margin-top: 5px;" />
          <!-- 预览区域 -->
          <el-image v-if="form.image_url" :src="form.image_url" style="width: 100%; height: 150px; margin-top: 10px; border-radius: 8px;" fit="cover" />
        </el-form-item>

        <el-form-item label="跳转链接">
          <el-input v-model="form.link_url" placeholder="如 /product/1 (商品详情页)" />
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

// 上传前校验
const beforeUpload = (file) => {
  const isImage = file.type.startsWith('image/')
  if (!isImage) {
    ElMessage.error('只能上传图片文件!')
    return false
  }
  const isLt2M = file.size / 1024 / 1024 < 2
  if (!isLt2M) {
    ElMessage.error('图片大小不能超过 2MB!')
    return false
  }
  return true
}

// 上传成功回调
const handleBannerUploadSuccess = (res) => {
  if (res.code === 0) {
    form.image_url = res.data.url // 后端返回的 URL
    ElMessage.success('图片上传成功')
  } else {
    ElMessage.error(res.msg || '上传失败')
  }
}

const fetchData = async () => {
  // ... 保持原有逻辑
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
  form.image_url = ''
  form.link_url = ''
  form.sort = 1
  dialogVisible.value = true
}

const submitForm = async () => {
  if (!form.image_url) return ElMessage.warning('请上传图片或填入图片链接')
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
