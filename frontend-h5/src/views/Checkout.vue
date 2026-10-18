<template>
  <div class="checkout-page">
    <van-nav-bar title="Checkout" left-arrow @click-left="$router.back()" fixed placeholder />

    <div class="section address-section">
      <h3>Shipping Address</h3>
      <van-field v-model="address.name" label="Name" placeholder="Receiver name" />
      <van-field v-model="address.phone" label="Phone" placeholder="Phone number" />
      <van-field v-model="address.detail" label="Address" type="textarea" placeholder="Detailed address" />
    </div>

    <div class="section items-section">
      <h3>Order Items</h3>
      <div class="order-item" v-for="item in checkoutItems" :key="item.id">
        <img :src="item.product.main_image" />
        <div class="info">
          <div class="title">{{ item.product.title }}</div>
          <div class="price">$ {{ item.sku ? item.sku.price : item.product.min_price }} x {{ item.quantity }}</div>
        </div>
      </div>
    </div>

    <!-- 底部提交栏 -->
    <van-submit-bar
      :price="totalPrice * 100"
      button-text="Place Order"
      @submit="submitOrder"
    />
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { getCartList } from '../api/cart'
import { createOrder } from '../api/order'
import { showToast, showSuccessToast } from 'vant'

const router = useRouter()
const checkoutItems = ref([])
const address = ref({ name: '', phone: '', detail: '' })

onMounted(async () => {
  // 实际项目中应该只传勾选的商品，这里简化处理，拿全部
  const res = await getCartList()
  checkoutItems.value = res.data.data || []
  if (checkoutItems.value.length === 0) {
    showToast('No items to checkout')
    router.replace('/cart')
  }
})

const totalPrice = computed(() => {
  return checkoutItems.value.reduce((sum, item) => {
    const price = item.sku ? item.sku.price : item.product.min_price
    return sum + (price * item.quantity)
  }, 0)
})

const submitOrder = async () => {
  if (!address.value.name || !address.value.phone || !address.value.detail) {
    return showToast('Please fill in all address fields')
  }

  try {
    const cartIds = checkoutItems.value.map(item => item.id)
    await createOrder({
      address_name: address.value.name,
      address_phone: address.value.phone,
      address_detail: address.value.detail,
      cart_ids: cartIds
    })
    
    showSuccessToast('Order placed successfully!')
    setTimeout(() => {
      router.push('/orders')
    }, 1000)
  } catch (e) {
    showToast(e.response?.data?.msg || 'Order failed')
  }
}
</script>

<style scoped>
.checkout-page { min-height: 100vh; background: #f7f8fa; padding-bottom: 60px; }
.section { background: #fff; margin-bottom: 10px; padding: 15px; }
.section h3 { margin: 0 0 10px 0; font-size: 15px; color: #333; }
.order-item { display: flex; gap: 10px; margin-bottom: 10px; }
.order-item img { width: 60px; height: 60px; object-fit: cover; border-radius: 4px; }
.info { flex: 1; }
.title { font-size: 13px; color: #333; margin-bottom: 5px; }
.price { font-size: 13px; color: #666; }
</style>
