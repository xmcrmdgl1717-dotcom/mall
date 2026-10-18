<template>
  <div class="detail-page" v-if="product">
    <!-- 顶部导航 -->
    <van-nav-bar
      title="Product Detail"
      left-text="Back"
      left-arrow
      @click-left="$router.back()"
      fixed
      placeholder
    />

    <!-- 图片轮播 (暂时只用主图) -->
    <div class="image-swipe">
      <img :src="product.main_image" class="main-img" />
    </div>

    <!-- 价格与标题区 -->
    <div class="info-section">
      <div class="price-row">
        <span class="currency">$</span>
        <span class="price">{{ selectedSku ? selectedSku.price : product.min_price }}</span>
        <span class="sales">Sold {{ product.total_sales }}</span>
      </div>
      <h2 class="title">{{ product.title }}</h2>
    </div>

    <!-- SKU 规格选择区 -->
    <div class="sku-section" v-if="product.skus && product.skus.length > 0">
      <div class="section-title">Select Specification</div>
      <div class="sku-list">
        <div 
          class="sku-item" 
          v-for="sku in product.skus" 
          :key="sku.id"
          :class="{ active: selectedSku && selectedSku.id === sku.id }"
          @click="selectSku(sku)"
        >
          {{ sku.specs }}
        </div>
      </div>
      <div class="stock-info" v-if="selectedSku">
        Stock: {{ selectedSku.stock }} pieces available
      </div>
    </div>

    <!-- 店铺信息区 (占位) -->
    <div class="store-section">
      <div class="store-info">
        <div class="avatar"></div>
        <div class="name">Verified Store</div>
      </div>
      <van-button size="small" plain type="primary" round>Visit Store</van-button>
    </div>

    <!-- 详情图文 (占位) -->
    <div class="desc-section">
      <div class="section-title">Product Description</div>
      <p style="color:#666; font-size:14px; line-height:1.6;">
        This is a high-quality product from FreshCo. We guarantee 100% original items, 7 days return policy, and secure payment. (Detailed rich text description placeholder)
      </p>
      <img :src="product.main_image" style="width:100%; margin-top:10px;" />
    </div>

    <!-- 底部悬浮操作栏 -->
    <div class="bottom-bar">
      <div class="icon-btn">
        <van-icon name="chat-o" size="20" />
        <span>Chat</span>
      </div>
      <div class="icon-btn">
        <van-icon name="shopping-cart-o" size="20" />
        <span>Cart</span>
      </div>
      <div class="action-btns">
        <van-button type="warning" class="add-cart-btn" @click="addToCart">Add to Cart</van-button>
        <van-button type="danger" class="buy-now-btn" @click="buyNow">Buy Now</van-button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { getProductDetail } from '../api/product'
import { showToast } from 'vant'

const route = useRoute()
const product = ref(null)
const selectedSku = ref(null)

onMounted(async () => {
  const id = route.params.id
  if (!id) return showToast('Product ID missing')
  
  try {
    const res = await getProductDetail(id)
    product.value = res.data.data
    // 默认选中第一个 SKU
    if (product.value.skus && product.value.skus.length > 0) {
      selectedSku.value = product.value.skus[0]
    }
  } catch (e) {
    showToast('Failed to load product')
  }
})

const selectSku = (sku) => {
  selectedSku.value = sku
}

const addToCart = () => {
  if (!selectedSku.value) return showToast('Please select a specification')
  showToast('Added to Cart (To be implemented)')
}

const buyNow = () => {
  if (!selectedSku.value) return showToast('Please select a specification')
  showToast('Redirecting to Checkout... (To be implemented)')
}
</script>

<style scoped>
.detail-page { background: #f5f5f5; min-height: 100vh; padding-bottom: 60px; }
.image-swipe { width: 100%; height: 375px; background: #fff; }
.main-img { width: 100%; height: 100%; object-fit: cover; }

.info-section { background: #fff; padding: 15px; margin-bottom: 10px; }
.price-row { display: flex; align-items: baseline; color: #ee0a24; margin-bottom: 10px; }
.currency { font-size: 16px; font-weight: bold; }
.price { font-size: 28px; font-weight: bold; margin-right: auto; }
.sales { font-size: 12px; color: #999; }
.title { font-size: 16px; color: #333; line-height: 1.4; margin: 0; }

.sku-section, .store-section, .desc-section { background: #fff; padding: 15px; margin-bottom: 10px; }
.section-title { font-size: 14px; font-weight: bold; margin-bottom: 10px; }
.sku-list { display: flex; flex-wrap: wrap; gap: 10px; }
.sku-item { padding: 8px 15px; background: #f5f5f5; border-radius: 4px; font-size: 14px; border: 1px solid transparent; }
.sku-item.active { background: #fff5f5; border-color: #ee0a24; color: #ee0a24; }
.stock-info { font-size: 12px; color: #999; margin-top: 10px; }

.store-section { display: flex; justify-content: space-between; align-items: center; }
.store-info { display: flex; align-items: center; gap: 10px; }
.avatar { width: 40px; height: 40px; background: #ddd; border-radius: 50%; }
.name { font-weight: bold; font-size: 14px; }

.bottom-bar { position: fixed; bottom: 0; left: 0; width: 100%; height: 60px; background: #fff; display: flex; align-items: center; padding: 0 10px; box-sizing: border-box; box-shadow: 0 -2px 10px rgba(0,0,0,0.05); }
.icon-btn { display: flex; flex-direction: column; align-items: center; font-size: 10px; color: #666; margin-right: 20px; }
.action-btns { flex: 1; display: flex; gap: 10px; }
.add-cart-btn, .buy-now-btn { flex: 1; border-radius: 20px; }
</style>
