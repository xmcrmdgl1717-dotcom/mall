<template>
  <div class="store-page">
    <van-nav-bar title="Store Detail" left-arrow @click-left="$router.back()" fixed placeholder />
    
    <div class="store-header" v-if="store">
      <div class="avatar"></div>
      <div class="info">
        <h2>{{ store.name }}</h2>
        <p>好评率: {{ store.good_rate }}% | 销量: {{ store.total_sales }}</p>
      </div>
    </div>

    <div class="product-list">
      <h3>全部商品</h3>
      <div class="grid">
        <div class="card" v-for="p in products" :key="p.id" @click="$router.push('/product/' + p.id)">
          <img :src="p.main_image" />
          <div class="title">{{ p.title }}</div>
          <div class="price">${{ p.min_price }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import axios from 'axios'

const route = useRoute()
const API_BASE = import.meta.env.VITE_API_BASE || 'https://mall-api-l3mx.onrender.com/api/v1'
const store = ref(null)
const products = ref([])

onMounted(async () => {
  const storeId = route.params.id
  // 这里为了简化，直接复用获取商品列表的接口，并假设后端支持按商家ID筛选
  // 实际开发中，你可以去后端加一个 /merchants/:id 和 /products?merchant_id=xxx 的接口
  try {
    // Mock 数据演示
    store.value = { name: 'FreshStore', good_rate: 99, total_sales: 1189321 }
    const res = await axios.get(`${API_BASE}/products?size=10`)
    products.value = res.data.data.list || []
  } catch (e) {
    console.error(e)
  }
})
</script>

<style scoped>
.store-page { min-height: 100vh; background: #f5f5f5; }
.store-header { background: #fff; padding: 20px; display: flex; align-items: center; gap: 15px; margin-bottom: 10px; }
.avatar { width: 60px; height: 60px; background: #ddd; border-radius: 50%; }
.info h2 { margin: 0 0 5px; font-size: 18px; }
.info p { margin: 0; font-size: 13px; color: #666; }
.product-list { padding: 15px; background: #fff; }
.grid { display: grid; grid-template-columns: repeat(2, 1fr); gap: 15px; margin-top: 10px; }
.card { background: #f9f9f9; padding: 10px; border-radius: 8px; }
.card img { width: 100%; aspect-ratio: 1; object-fit: cover; border-radius: 4px; }
.title { font-size: 13px; margin: 5px 0; display: -webkit-box; -webkit-line-clamp: 2; -webkit-box-orient: vertical; overflow: hidden; }
.price { color: #ee0a24; font-weight: bold; font-size: 16px; }
</style>
