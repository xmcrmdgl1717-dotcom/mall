<template>
  <div class="checkout-page">
    <van-nav-bar title="Checkout" left-arrow @click-left="$router.back()" fixed placeholder />

    <!-- 收货地址卡片 -->
    <div class="section address-section">
      <div v-if="selectedAddress" class="address-card" @click="$router.push('/addresses')">
        <div class="user">
          <span style="font-weight: bold; margin-right: 10px;">{{ selectedAddress.name }}</span>
          <span>{{ selectedAddress.phone }}</span>
        </div>
        <div class="detail">{{ selectedAddress.province }} {{ selectedAddress.city }} {{ selectedAddress.detail }}</div>
        <van-icon name="arrow" class="arrow" />
      </div>
      <div v-else class="no-address" @click="$router.push('/addresses')">
        <van-icon name="plus" /> Add Shipping Address
      </div>
    </div>

    <!-- 订单商品列表 -->
    <div class="section items-section">
      <h3>Order Items</h3>
      <div class="order-item" v-for="item in checkoutItems" :key="item.id">
        <img :src="item.product.main_image" />
        <div class="info">
          <div class="title">{{ item.product.title }}</div>
          <div class="specs">{{ item.sku?.specs }}</div>
          <div class="price">$ {{ item.sku ? item.sku.price : item.product.min_price }} x {{ item.quantity }}</div>
        </div>
      </div>
    </div>

    <!-- 底部提交栏 -->
    <van-submit-bar :price="totalPrice * 100" button-text="Place Order" @submit="submitOrder" />
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { getCartList } from '../api/cart'
import { getAddresses } from '../api/address'
import { createOrder } from '../api/order'
import { showToast, showSuccessToast } from 'vant'

const router = useRouter()
const checkoutItems = ref([])
const selectedAddress = ref(null)

onMounted(async () => {
  // 1. 拉取购物车数据
  const cartRes = await getCartList()
  checkoutItems.value = cartRes.data.data || []
  if (checkoutItems.value.length === 0) {
    showToast('No items to checkout')
    router.replace('/cart')
    return
  }

  // 2. 拉取地址数据（后端已按 is_default DESC 排序，默认选第一个）
  const addrRes = await getAddresses()
  const addresses = addrRes.data.data || []
  if (addresses.length > 0) {
    selectedAddress.value = addresses[0]
  }
})

// 计算总价
const totalPrice = computed(() => {
  return checkoutItems.value.reduce((sum, item) => {
    const price = item.sku ? item.sku.price : item.product.min_price
    return sum + (price * item.quantity)
  }, 0)
})

const submitOrder = async () => {
  if (!selectedAddress.value) {
    return showToast('Please add a shipping address')
  }

  try {
    const cartIds = checkoutItems.value.map(item => item.id)
    await createOrder({
      address_name: selectedAddress.value.name,
      address_phone: selectedAddress.value.phone,
      address_detail: selectedAddress.value.detail,
      cart_ids: cartIds
    })
    
    showSuccessToast('Order placed successfully!')
    setTimeout(() => { router.push('/orders') }, 1000)
  } catch (e) {
    showToast(e.response?.data?.msg || 'Order failed')
  }
}
</script>

<style scoped>
.checkout-page { min-height: 100vh; background: #f7f8fa; padding-bottom: 60px; }
.section { background: #fff; margin-bottom: 10px; padding: 15px; }
.section h3 { margin: 0 0 15px 0; font-size: 16px; color: #333; }

.address-card { position: relative; padding-right: 20px; cursor: pointer; }
.user { margin-bottom: 5px; font-size: 15px; }
.detail { font-size: 13px; color: #666; line-height: 1.4; }
.arrow { position: absolute; right: 0; top: 50%; transform: translateY(-50%); color: #999; }
.no-address { padding: 20px; text-align: center; border: 1px dashed #ee0a24; color: #ee0a24; border-radius: 8px; font-weight: bold; cursor: pointer; }

.order-item { display: flex; gap: 10px; margin-bottom: 15px; }
.order-item img { width: 70px; height: 70px; object-fit: cover; border-radius: 4px; }
.info { flex: 1; display: flex; flex-direction: column; justify-content: space-between; }
.title { font-size: 14px; color: #333; line-height: 1.4; display: -webkit-box; -webkit-line-clamp: 2; -webkit-box-orient: vertical; overflow: hidden; }
.specs { font-size: 12px; color: #999; background: #f5f5f5; padding: 2px 6px; border-radius: 4px; align-self: flex-start; }
.price { font-size: 14px; color: #ee0a24; font-weight: bold; text-align: right; }
</style>
