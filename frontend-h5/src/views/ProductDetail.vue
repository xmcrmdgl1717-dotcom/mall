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

    <!-- 图片轮播区 -->
    <div class="swipe-box">
      <van-swipe class="my-swipe" :autoplay="3000" indicator-color="white">
        <van-swipe-item>
          <img :src="product.main_image" alt="main" />
        </van-swipe-item>
      </van-swipe>
    </div>

    <!-- 价格与标题 -->
    <div class="info-section">
      <div class="price-row">
        <span class="currency">$</span>
        <span class="price">{{ currentPrice }}</span>
        <span class="sales">Sold {{ product.total_sales }}</span>
      </div>
      <h2 class="title">{{ product.title }}</h2>
    </div>

    <!-- SKU 规格选择 -->
    <div class="sku-section" v-if="product.skus && product.skus.length > 0">
      <div class="section-title">Select Specification</div>
      <div class="sku-list">
        <div 
          class="sku-item" 
          v-for="sku in product.skus" 
          :key="sku.id"
          :class="{ active: selectedSkuId === sku.id }"
          @click="selectSku(sku)"
        >
          {{ sku.specs }}
        </div>
      </div>
      <div class="stock-info" v-if="selectedSku">
        Stock: {{ selectedSku.stock }} pieces available
      </div>
    </div>

    <!-- 店铺信息 -->
    <div class="store-section">
      <div class="store-info">
        <div class="avatar"></div>
        <div class="name">Verified Store</div>
      </div>
      <van-button size="small" plain type="primary" round>Visit Store</van-button>
    </div>

    <!-- 商品图文详情 (占位) -->
    <div class="desc-section">
      <div class="section-title">Product Description</div>
      <p style="color:#666; font-size:14px; line-height:1.6;">
        This is a high-quality product from FreshCo. 100% Original, 7 days return.
      </p>
      <img :src="product.main_image" style="width:100%; margin-top:10px; border-radius: 8px;" />
    </div>

    <!-- 底部悬浮操作栏 -->
    <van-action-bar>
      <van-action-bar-icon icon="chat-o" text="Chat" />
      <!-- 点击跳转购物车 -->
      <van-action-bar-icon icon="cart-o" text="Cart" @click="$router.push('/cart')" />
      <van-action-bar-button type="warning" text="Add to Cart" @click="addToCartAction" />
      <van-action-bar-button type="danger" text="Buy Now" @click="buyNow" />
    </van-action-bar>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { getProductDetail } from '../api/product'
import { addToCart } from '../api/cart' // 引入购物车 API
import { showToast } from 'vant'

const route = useRoute()
const router = useRouter()
const product = ref(null)
const selectedSku = ref(null)

const selectedSkuId = computed(() => selectedSku.value ? selectedSku.value.id : null)

const currentPrice = computed(() => {
  if (selectedSku.value) return selectedSku.value.price
  return product.value ? product.value.min_price : 0
})

onMounted(async () => {
  const id = route.params.id
  if (!id) return showToast('Product ID missing')
  
  try {
    const res = await getProductDetail(id)
    product.value = res.data.data
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

// 加入购物车逻辑
const addToCartAction = async () => {
  if (product.value.skus.length > 0 && !selectedSku.value) {
    return showToast('Please select a specification')
  }
  
  try {
    await addToCart({
      product_id: product.value.id,
      sku_id: selectedSku.value ? selectedSku.value.id : 0,
      quantity: 1
    })
    showToast('Added to Cart successfully!')
  } catch (e) {
    showToast('Failed to add to cart')
  }
}

const buyNow = () => {
  if (product.value.skus.length > 0 && !selectedSku.value) {
    return showToast('Please select a specification')
  }
  router.push('/checkout')
}
</script>

<style scoped>
.detail-page { background: #f7f8fa; min-height: 100vh; padding-bottom: 60px; }
.swipe-box { width: 100%; height: 375px; background: #fff; }
.my-swipe img { width: 100%; height: 375px; object-fit: cover; }

.info-section { background: #fff; padding: 15px; margin-bottom: 10px; }
.price-row { display: flex; align-items: baseline; color: #ee0a24; margin-bottom: 10px; }
.currency { font-size: 16px; font-weight: bold; }
.price { font-size: 28px; font-weight: bold; margin-right: auto; }
.sales { font-size: 12px; color: #999; }
.title { font-size: 16px; color: #333; line-height: 1.4; margin: 0; }

.sku-section, .store-section, .desc-section { background: #fff; padding: 15px; margin-bottom: 10px; }
.section-title { font-size: 14px; font-weight: bold; margin-bottom: 10px; }
.sku-list { display: flex; flex-wrap: wrap; gap: 10px; }
.sku-item { padding: 8px 15px; background: #f5f5f5; border-radius: 4px; font-size: 14px; border: 1px solid transparent; cursor: pointer; }
.sku-item.active { background: #fff5f5; border-color: #ee0a24; color: #ee0a24; }
.stock-info { font-size: 12px; color: #999; margin-top: 10px; }

.store-section { display: flex; justify-content: space-between; align-items: center; }
.store-info { display: flex; align-items: center; gap: 10px; }
.avatar { width: 40px; height: 40px; background: #ddd; border-radius: 50%; }
.name { font-weight: bold; font-size: 14px; }
</style>
