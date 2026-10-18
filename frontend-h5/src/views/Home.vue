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
      <!-- Banner 区 -->
      <div class="banner-section">
        <div class="banner-left">
          <img src="https://picsum.photos/800/350?grayscale" alt="Banner" />
        </div>
        <div class="banner-right">
          <img src="https://picsum.photos/400/350" alt="Small Banner" />
        </div>
      </div>

      <!-- Shop by Category -->
      <section class="section">
        <div class="section-header">
          <h2>Shop by Category</h2>
          <a href="#" class="view-all">View All ></a>
        </div>
        <div class="category-list">
          <div class="category-item" v-for="i in 12" :key="i">
            <div class="cat-icon"></div>
            <span>Category {{ i }}</span>
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
import { getProducts } from '../api/product'
import { getCategories } from '../api/category'
import { getTopMerchants } from '../api/merchant'

const dailyDeals = ref([])
const popularItems = ref([])
const hotDeals = ref([])
const popularStores = ref([])

onMounted(async () => {
  try {
    const dailyRes = await getProducts({ deal_type: 'daily', size: 10 })
    dailyDeals.value = dailyRes.data.data.list || []

    const popularRes = await getProducts({ sort: 'sales', size: 10 })
    popularItems.value = popularRes.data.data.list || []

    const hotRes = await getProducts({ deal_type: 'hot', size: 10 })
    hotDeals.value = hotRes.data.data.list || []

    const storeRes = await getTopMerchants()
    popularStores.value = storeRes.data.data || []
    
    if (dailyDeals.value.length === 0) {
       dailyDeals.value = Array(5).fill({ id: 1, title: 'YLTNEUQER Jade Roller Gua Sha Set', min_price: '5.62', total_sales: 1358, main_image: 'https://picsum.photos/200' })
    }
  } catch (e) {
    console.error(e)
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
.banner-section { display: flex; gap: 10px; margin-bottom: 20px; }
.banner-left { flex: 2; }
.banner-left img { width: 100%; border-radius: 8px; display: block; }
.banner-right { flex: 1; }
.banner-right img { width: 100%; height: 100%; object-fit: cover; border-radius: 8px; display: block; }

.section { background: #fff; padding: 20px; margin-bottom: 20px; border-radius: 8px; }
.section-header { display: flex; justify-content: space-between; margin-bottom: 15px; border-bottom: 1px solid #eee; padding-bottom: 10px; }
.section-header h2 { font-size: 18px; margin: 0; }
.view-all { font-size: 12px; color: #666; text-decoration: none; }

.category-list { display: flex; justify-content: space-between; text-align: center; }
.category-item { flex: 1; }
.cat-icon { width: 50px; height: 50px; background: #eee; border-radius: 50%; margin: 0 auto 8px; }
.category-item span { font-size: 12px; color: #333; }

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
