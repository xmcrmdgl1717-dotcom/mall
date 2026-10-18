<template>
  <div class="home-page">
    <!-- 顶部搜索 -->
    <van-search v-model="keyword" placeholder="Search for products..." shape="round" background="#fff" />

    <!-- 分类九宫格 (横向滑动) -->
    <div class="category-section">
      <div class="section-title">Shop by Category <span class="more">View All ></span></div>
      <div class="category-scroll">
        <div class="category-item" v-for="c in categories" :key="c.id">
          <div class="c-icon">
            <!-- 用随机颜色代替真实图标 -->
            <div class="circle" :style="{ backgroundColor: getRandomColor() }"></div>
          </div>
          <span class="c-name">{{ c.name }}</span>
        </div>
      </div>
    </div>

    <!-- 每日特惠 (Daily Deals) -->
    <div class="deals-section">
      <div class="section-title">Daily Deals <span class="more">More ></span></div>
      <div class="product-grid">
        <div class="product-card" v-for="p in dailyDeals" :key="p.id">
          <img :src="p.main_image" class="p-img" />
          <div class="p-title van-multi-ellipsis--l2">{{ p.title }}</div>
          <div class="p-price">${{ p.min_price }}</div>
          <div class="p-sales">Sold {{ p.total_sales }}</div>
          <van-button type="primary" size="small" block round class="add-btn">Add to cart</van-button>
        </div>
      </div>
    </div>
    
    <!-- 后续可以按同样的结构继续写 Popular Items 和 Hot Deals -->
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getProducts } from '../api/product'
import { getCategories } from '../api/category'

const keyword = ref('')
const categories = ref([])
const dailyDeals = ref([])

// 生成随机颜色占位
const getRandomColor = () => {
  const colors = ['#ff9f43', '#0abde3', '#10ac84', '#ee5253', '#5f27cd']
  return colors[Math.floor(Math.random() * colors.length)]
}

onMounted(async () => {
  try {
    // 1. 获取分类
    const catRes = await getCategories()
    categories.value = catRes.data.data || []

    // 2. 获取每日特惠 (deal_type=daily)
    const dealRes = await getProducts({ deal_type: 'daily', size: 6 })
    dailyDeals.value = dealRes.data.data.list || []
  } catch (e) {
    console.error('Failed to load data', e)
  }
})
</script>

<style scoped>
.home-page { background-color: #f7f8fa; min-height: 100vh; padding-bottom: 20px; }
.section-title { font-size: 18px; font-weight: bold; color: #333; padding: 12px; display: flex; justify-content: space-between; align-items: center; }
.section-title .more { font-size: 12px; color: #999; font-weight: normal; }

/* 分类横向滚动 */
.category-section { background: #fff; padding-bottom: 12px; margin-bottom: 10px; }
.category-scroll { display: flex; overflow-x: auto; padding: 0 12px; gap: 15px; }
.category-scroll::-webkit-scrollbar { display: none; }
.category-item { display: flex; flex-direction: column; align-items: center; min-width: 60px; }
.circle { width: 50px; height: 50px; border-radius: 50%; margin-bottom: 6px; opacity: 0.7; }
.c-name { font-size: 12px; color: #666; text-align: center; line-height: 1.2; }

/* 商品瀑布流/网格 */
.deals-section { background: #fff; padding-bottom: 15px; margin-bottom: 10px; }
.product-grid { display: flex; flex-wrap: wrap; padding: 0 6px; }
.product-card { width: calc(50% - 12px); margin: 6px; background: #f9f9f9; border-radius: 8px; padding: 8px; box-sizing: border-box; display: flex; flex-direction: column; }
.p-img { width: 100%; aspect-ratio: 1; object-fit: cover; border-radius: 6px; background: #eee; }
.p-title { font-size: 13px; color: #333; margin: 6px 0; flex-grow: 1; }
.p-price { font-size: 16px; color: #ee0a24; font-weight: bold; }
.p-sales { font-size: 11px; color: #999; margin-bottom: 8px; }
.add-btn { background: #ff6000; border: none; height: 28px; }
</style>
