<template>
  <div class="home-container">
    <!-- 顶部导航 (Header) -->
    <header class="header">
      <div class="top-bar">
        <div class="logo">FRESH<span>CO</span></div>
        <div class="search-bar">
          <input type="text" placeholder="Search for brand/product/suppliers" />
          <button>Search</button>
        </div>
        <div class="user-actions">
          <span>👤 Log in or Register</span>
          <span>🛒 Cart</span>
        </div>
      </div>
      <nav class="nav-links">
        <a href="#">Home</a>
        <a href="#">Category</a>
        <a href="#">Product</a>
        <a href="#">Discounts</a>
        <a href="#">Partnership</a>
        <a href="#">Credit Loan Service</a>
      </nav>
    </header>

    <main class="main-content">
      <!-- 动态 Banner 轮播区 (带跳转) -->
      <div class="banner-section" v-if="bannerList.length > 0">
        <van-swipe class="my-swipe" :autoplay="3000" indicator-color="white" style="width: 100%; border-radius: 8px; height: 350px;">
          <van-swipe-item 
            v-for="item in bannerList" 
            :key="item.id" 
            @click="item.link_url && $router.push(item.link_url)"
            :style="{ cursor: item.link_url ? 'pointer' : 'default' }"
          >
            <img :src="item.image_url" style="width: 100%; height: 350px; object-fit: cover;" />
          </van-swipe-item>
        </van-swipe>
      </div>
      <!-- 如果后台没配 Banner，显示占位图 -->
      <div class="banner-section" v-else>
        <img src="https://picsum.photos/1200/350?grayscale" style="width: 100%; border-radius: 8px; height: 350px; object-fit: cover;" />
      </div>

      <!-- Shop by Category -->
      <section class="section">
        <div class="section-header">
          <h2>Shop by Category</h2>
          <a href="#" class="view-all">View All ></a>
        </div>
        <div class="category-list">
          <div class="category-item" v-for="c in categories" :key="c.id" @click="$router.push('/products?category=' + c.id)">
            <div class="cat-icon"></div>
            <span>{{ c.name }}</span>
          </div>
          <div class="category-item" v-if="categories.length === 0">
             <div class="cat-icon"></div>
             <span>暂无分类</span>
          </div>
        </div>
      </section>

      <!-- Daily Deals -->
      <section class="section">
        <div class="section-header">
          <h2>Daily Deals</h2>
          <a href="#" class="view-all">More ></a>
        </div>
        <div class="product-grid">
          <div class="product-card" v-for="p in dailyDeals" :key="p.id" @click="$router.push('/product/' + p.id)">
            <img :src="p.main_image" class="p-img" />
            <div class="p-price">${{ p.min_price }}</div>
            <div class="p-sales">Sold {{ p.total_sales }}</div>
            <div class="p-title">{{ p.title }}</div>
            <div class="p-actions">
              <span>🛒 Purchase Now</span>
              <span class="heart">♡</span>
            </div>
          </div>
        </div>
      </section>

      <!-- 成为卖家 Banner -->
      <div class="seller-banner">
        <div class="text">
          <h2>Become a Seller</h2>
          <p>Total Commission</p>
          <h1>Max. $100,000</h1>
          <button class="join-btn">Join Now</button>
        </div>
      </div>

      <!-- Popular Items -->
      <section class="section">
        <div class="section-header">
          <h2>Popular Items</h2>
        </div>
        <div class="product-grid">
          <div class="product-card" v-for="p in popularItems" :key="p.id" @click="$router.push('/product/' + p.id)">
             <img :src="p.main_image" class="p-img" />
            <div class="p-price">${{ p.min_price }}</div>
            <div class="p-sales">Sold {{ p.total_sales }}</div>
            <div class="p-title">{{ p.title }}</div>
            <div class="p-actions">
              <span>🛒 Purchase Now</span>
              <span class="heart">♡</span>
            </div>
          </div>
        </div>
      </section>

      <!-- Popular Stores -->
      <section class="section">
        <div class="section-header">
          <h2>Popular Stores</h2>
        </div>
        <div class="store-grid">
          <div class="store-card" v-for="s in popularStores" :key="s.id">
            <div class="s-header">
              <div class="s-avatar"></div>
              <div class="s-name">{{ s.name }}</div>
            </div>
            <div class="s-stats">
              <p>Product: {{ s.product_count || 50 }}</p>
              <p>Total Sales: {{ s.total_sales }}</p>
              <p>Rate of good reviews: {{ s.good_rate }}%</p>
            </div>
            <button class="visit-btn">Visit Store ></button>
          </div>
        </div>
      </section>

      <!-- Hot Deals -->
      <section class="section">
        <div class="section-header">
          <h2>Hot Deals</h2>
        </div>
        <div class="product-grid">
          <div class="product-card" v-for="p in hotDeals" :key="p.id" @click="$router.push('/product/' + p.id)">
             <img :src="p.main_image" class="p-img" />
            <div class="p-price">${{ p.min_price }}</div>
            <div class="p-sales">Sold {{ p.total_sales }}</div>
            <div class="p-title">{{ p.title }}</div>
            <div class="p-actions">
              <span>🛒 Purchase Now</span>
              <span class="heart">♡</span>
            </div>
          </div>
        </div>
      </section>

      <!-- 底部信任保障栏 -->
      <div class="footer-trust">
        <div><span>♡</span><p>100% Original</p></div>
        <div><span>⏱</span><p>7 days return</p></div>
        <div><span>🚚</span><p>Freight discount</p></div>
        <div><span>🔒</span><p>Secure Payment</p></div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { getProducts } from '../api/product'
import { getCategories } from '../api/category'
import { getTopMerchants } from '../api/merchant'

const dailyDeals = ref([])
const popularItems = ref([])
const hotDeals = ref([])
const popularStores = ref([])
const categories = ref([])
const bannerList = ref([])

const API_BASE = import.meta.env.VITE_API_BASE || 'https://mall-api-l3mx.onrender.com/api/v1'

onMounted(async () => {
  try {
    // 1. 获取动态 Banner
    const bannerRes = await axios.get(`${API_BASE}/banners`)
    bannerList.value = bannerRes.data.data || []

    // 2. 获取动态分类
    const catRes = await getCategories()
    categories.value = catRes.data.data || []

    // 3. 获取商品与店铺
    const dailyRes = await getProducts({ deal_type: 'daily', size: 10 })
    dailyDeals.value = dailyRes.data.data.list || []

    const popularRes = await getProducts({ sort: 'sales', size: 10 })
    popularItems.value = popularRes.data.data.list || []

    const hotRes = await getProducts({ deal_type: 'hot', size: 10 })
    hotDeals.value = hotRes.data.data.list || []

    const storeRes = await getTopMerchants()
    popularStores.value = storeRes.data.data || []
    
    if (dailyDeals.value.length === 0) {
       dailyDeals.value = popularItems.value.slice(0, 5)
    }
  } catch (e) {
    console.error('首页数据加载失败', e)
  }
})
</script>

<style scoped>
.home-container { background: #f5f5f5; font-family: Arial, sans-serif; min-width: 1200px; }
.header { background: #fff; padding: 15px 0; border-bottom: 1px solid #eee; }
.top-bar { display: flex; justify-content: space-between; align-items: center; width: 1200px; margin: 0 auto; }
.logo { font-size: 24px; font-weight: 900; color: #333; }
.logo span { background: #ccff00; padding: 0 4px; }
.search-bar { display: flex; width: 50%; }
.search-bar input { flex: 1; padding: 10px; border: 2px solid #000; border-right: none; border-radius: 20px 0 0 20px; outline: none; }
.search-bar button { padding: 10px 20px; background: #000; color: #fff; border: 2px solid #000; border-radius: 0 20px 20px 0; cursor: pointer; }
.nav-links { display: flex; justify-content: center; gap: 30px; margin-top: 15px; font-size: 14px; font-weight: bold; }
.nav-links a { color: #333; text-decoration: none; }

.main-content { width: 1200px; margin: 20px auto; }
.banner-section { margin-bottom: 20px; }

.section { background: #fff; padding: 20px; margin-bottom: 20px; border-radius: 8px; }
.section-header { display: flex; justify-content: space-between; margin-bottom: 15px; border-bottom: 1px solid #eee; padding-bottom: 10px; }
.section-header h2 { font-size: 18px; margin: 0; }
.view-all { font-size: 12px; color: #666; text-decoration: none; }

.category-list { display: flex; justify-content: flex-start; gap: 30px; text-align: center; overflow-x: auto; padding-bottom: 10px; }
.category-item { flex: 0 0 80px; cursor: pointer; }
.cat-icon { width: 50px; height: 50px; background: #eee; border-radius: 50%; margin: 0 auto 8px; }
.category-item span { font-size: 12px; color: #333; display: block; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }

.product-grid { display: grid; grid-template-columns: repeat(5, 1fr); gap: 15px; }
.product-card { border: 1px solid #f0f0f0; border-radius: 8px; padding: 10px; transition: 0.3s; display: flex; flex-direction: column; cursor: pointer; }
.product-card:hover { box-shadow: 0 4px 12px rgba(0,0,0,0.1); }
.p-img { width: 100%; aspect-ratio: 1; object-fit: cover; border-radius: 4px; margin-bottom: 8px; background: #f9f9f9; }
.p-price { color: #e4393c; font-size: 16px; font-weight: bold; }
.p-sales { font-size: 12px; color: #999; margin-bottom: 5px; }
.p-title { font-size: 13px; color: #333; height: 36px; overflow: hidden; display: -webkit-box; -webkit-line-clamp: 2; -webkit-box-orient: vertical; line-height: 18px; margin-bottom: 10px; flex-grow: 1;}
.p-actions { display: flex; justify-content: space-between; font-size: 12px; color: #666; border-top: 1px solid #eee; padding-top: 10px; }

.seller-banner { background: #fdf1e9; border-radius: 8px; padding: 40px; margin-bottom: 20px; display: flex; align-items: center; justify-content: flex-end; background-image: url('https://picsum.photos/1200/200'); background-size: cover; }
.seller-banner .text { text-align: right; }
.join-btn { background: #333; color: #fff; border: none; padding: 10px 30px; border-radius: 20px; margin-top: 10px; cursor: pointer; }

.store-grid { display: grid; grid-template-columns: repeat(3, 1fr); gap: 20px; }
.store-card { border: 1px solid #eee; padding: 15px; border-radius: 8px; }
.s-header { display: flex; align-items: center; gap: 10px; margin-bottom: 10px; }
.s-avatar { width: 40px; height: 40px; border-radius: 50%; background: #ddd; }
.s-name { font-weight: bold; }
.s-stats p { font-size: 12px; color: #666; margin: 4px 0; }
.visit-btn { width: 100%; border: 1px solid #333; background: #fff; padding: 8px; border-radius: 20px; cursor: pointer; }

.footer-trust { display: flex; justify-content: space-around; background: #fff; padding: 30px 0; border-radius: 8px; }
.footer-trust div { text-align: center; font-size: 14px; font-weight: bold; }
.footer-trust span { font-size: 24px; display: block; margin-bottom: 10px; }
</style>
