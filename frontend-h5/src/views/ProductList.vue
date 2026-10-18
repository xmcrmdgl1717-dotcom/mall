<template>
  <div class="list-page">
    <!-- 顶部搜索与返回 -->
    <van-nav-bar left-arrow @click-left="$router.back()" fixed placeholder>
      <template #title>
        <div class="search-box">
          <van-search
            v-model="searchKeyword"
            shape="round"
            background="transparent"
            placeholder="Search products"
            @search="onSearch"
          />
        </div>
      </template>
    </van-nav-bar>

    <!-- 排序与筛选栏 -->
    <van-tabs v-model:active="sortType" @change="onSortChange" sticky offset-top="46">
      <van-tab title="综合" name="default" />
      <van-tab title="销量优先" name="sales" />
      <van-tab title="价格从低到高" name="price_asc" />
      <van-tab title="价格从高到低" name="price_desc" />
    </van-tabs>

    <!-- 商品瀑布流 -->
    <div class="product-list" v-if="products.length > 0">
      <div class="product-card" v-for="p in products" :key="p.id" @click="$router.push('/product/' + p.id)">
        <img :src="p.main_image" class="p-img" />
        <div class="p-title van-multi-ellipsis--l2">{{ p.title }}</div>
        <div class="p-price">${{ p.min_price }}</div>
        <div class="p-sales">Sold {{ p.total_sales }}</div>
      </div>
    </div>

    <!-- 空状态 -->
    <van-empty v-else-if="!loading" description="No products found" />

    <!-- 加载状态 -->
    <div class="loading-box" v-if="loading">
      <van-loading color="#1989fa" />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { getProducts } from '../api/product'

const route = useRoute()
const router = useRouter()

const searchKeyword = ref('')
const categoryId = ref(0)
const sortType = ref('default')

const products = ref([])
const loading = ref(false)
const page = ref(1)
const finished = ref(false)
const pageSize = 10

onMounted(() => {
  // 从路由参数中读取搜索词和分类
  if (route.query.keyword) {
    searchKeyword.value = route.query.keyword
  }
  if (route.query.category) {
    categoryId.value = Number(route.query.category)
  }
  fetchProducts(true)
})

const onSearch = () => {
  // 搜索时重置分页和列表
  page.value = 1
  finished.value = false
  products.value = []
  fetchProducts(true)
}

const onSortChange = () => {
  // 切换排序时重置分页
  page.value = 1
  finished.value = false
  products.value = []
  fetchProducts(true)
}

const fetchProducts = async (reset = false) => {
  if (loading.value || (finished.value && !reset)) return
  loading.value = true

  try {
    const params = {
      page: page.value,
      size: pageSize,
      keyword: searchKeyword.value,
      category_id: categoryId.value,
      sort: sortType.value
    }
    const res = await getProducts(params)
    const list = res.data.data.list || []
    
    if (reset) {
      products.value = list
    } else {
      products.value = [...products.value, ...list]
    }

    if (list.length < pageSize) {
      finished.value = true
    } else {
      page.value++
    }
  } catch (e) {
    console.error('加载失败', e)
  } finally {
    loading.value = false
  }
}

// 触底加载更多 (可以用 IntersectionObserver 或 Vant 的 List 组件，这里为了简单用原生滚动)
const handleScroll = () => {
  const bottom = document.documentElement.scrollHeight - document.documentElement.scrollTop <= document.documentElement.clientHeight + 50
  if (bottom) {
    fetchProducts(false)
  }
}

onMounted(() => {
  window.addEventListener('scroll', handleScroll)
})

import { onUnmounted } from 'vue'
onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll)
})
</script>

<style scoped>
.list-page { min-height: 100vh; background: #f5f5f5; padding-bottom: 20px; }
.search-box { width: 100%; padding: 0 10px; }
.product-list { display: flex; flex-wrap: wrap; padding: 10px; gap: 10px; justify-content: space-between; }
.product-card { width: calc(50% - 5px); background: #fff; border-radius: 8px; padding: 10px; box-sizing: border-box; cursor: pointer; }
.p-img { width: 100%; aspect-ratio: 1; object-fit: cover; border-radius: 6px; margin-bottom: 8px; }
.p-title { font-size: 13px; color: #333; margin-bottom: 6px; line-height: 18px; }
.p-price { color: #ee0a24; font-size: 16px; font-weight: bold; }
.p-sales { font-size: 11px; color: #999; margin-top: 4px; }
.loading-box { text-align: center; padding: 20px 0; }
</style>
