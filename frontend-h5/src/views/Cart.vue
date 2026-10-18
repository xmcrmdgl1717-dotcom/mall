<template>
  <div class="cart-page">
    <van-nav-bar title="Shopping Cart" fixed placeholder />

    <div v-if="cartList.length === 0" class="empty-cart">
      <van-empty description="Your cart is empty" />
      <van-button type="primary" round to="/">Go Shopping</van-button>
    </div>

    <div v-else class="cart-list">
      <van-swipe-cell v-for="item in cartList" :key="item.id">
        <div class="cart-item">
          <van-checkbox v-model="item.selected" :true-value="1" :false-value="0" />
          <img :src="item.product.main_image" class="item-img" />
          <div class="item-info">
            <div class="item-title van-multi-ellipsis--l2">{{ item.product.title }}</div>
            <div class="item-sku" v-if="item.sku">{{ item.sku.specs }}</div>
            <div class="item-bottom">
              <div class="item-price">$ {{ item.sku ? item.sku.price : item.product.min_price }}</div>
              <van-stepper v-model="item.quantity" min="1" />
            </div>
          </div>
        </div>
        <template #right>
          <van-button square text="Delete" type="danger" class="delete-btn" @click="onDelete(item.id)" />
        </template>
      </van-swipe-cell>
    </div>

    <!-- 底部结算栏 -->
    <van-submit-bar
      :price="totalPrice * 100"
      button-text="Checkout"
      @submit="onSubmit"
    >
      <van-checkbox v-model="allChecked" @click="toggleAll">Select All</van-checkbox>
    </van-submit-bar>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { getCartList, deleteCartItem } from '../api/cart'
import { showToast, showConfirmDialog } from 'vant'

const cartList = ref([])

const fetchCart = async () => {
  try {
    const res = await getCartList()
    cartList.value = res.data.data || []
    // 默认全选
    cartList.value.forEach(item => item.selected = 1)
  } catch (e) {
    showToast('Failed to load cart')
  }
}

onMounted(() => {
  fetchCart()
})

// 计算总价 (这里为了方便简单乘100，实际应该转换成分为单位)
const totalPrice = computed(() => {
  return cartList.value
    .filter(item => item.selected === 1)
    .reduce((sum, item) => {
      const price = item.sku ? item.sku.price : item.product.min_price
      return sum + (price * item.quantity)
    }, 0)
})

// 全选逻辑
const allChecked = computed({
  get: () => cartList.value.length > 0 && cartList.value.every(item => item.selected === 1),
  set: (val) => {
    cartList.value.forEach(item => item.selected = val ? 1 : 0)
  }
})

const toggleAll = () => {
  // 由 computed 自动处理
}

const onDelete = async (id) => {
  try {
    await showConfirmDialog({ title: 'Delete', message: 'Are you sure?' })
    await deleteCartItem(id)
    showToast('Deleted')
    fetchCart()
  } catch (e) {
    // 用户取消
  }
}

const onSubmit = () => {
  const selectedItems = cartList.value.filter(item => item.selected === 1)
  if (selectedItems.length === 0) return showToast('Please select items')
  showToast('Going to checkout... (To be implemented)')
}
</script>

<style scoped>
.cart-page { min-height: 100vh; background: #f7f8fa; padding-bottom: 50px; }
.cart-list { padding: 10px; }
.cart-item { display: flex; align-items: center; background: #fff; padding: 15px 10px; border-radius: 8px; margin-bottom: 10px; }
.item-img { width: 80px; height: 80px; object-fit: cover; border-radius: 6px; margin: 0 10px; }
.item-info { flex: 1; display: flex; flex-direction: column; }
.item-title { font-size: 14px; color: #333; margin-bottom: 5px; }
.item-sku { font-size: 12px; color: #999; background: #f5f5f5; padding: 2px 6px; border-radius: 4px; align-self: flex-start; margin-bottom: 5px; }
.item-bottom { display: flex; justify-content: space-between; align-items: center; margin-top: auto; }
.item-price { font-size: 16px; color: #ee0a24; font-weight: bold; }
.delete-btn { height: 100%; }
.empty-cart { padding-top: 100px; text-align: center; display: flex; flex-direction: column; align-items: center; gap: 20px; }
</style>
