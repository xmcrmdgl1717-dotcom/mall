<template>
  <div class="orders-page">
    <van-nav-bar title="My Orders" left-arrow @click-left="$router.push('/')" fixed placeholder />
    
    <div v-if="orders.length === 0" class="empty">
      <van-empty description="No orders found" />
    </div>

    <div v-else class="order-list">
      <div class="order-card" v-for="order in orders" :key="order.id">
        <div class="order-header">
          <span class="no">Order No: {{ order.order_no }}</span>
          <span class="status">{{ getStatusText(order.status) }}</span>
        </div>
        <div class="order-body">
          <div class="item" v-for="item in order.items" :key="item.id">
            <img :src="item.image" />
            <div class="info">
              <div class="title">{{ item.title }}</div>
              <div class="specs">{{ item.specs }}</div>
            </div>
            <div class="price">${{ item.price }} x {{ item.quantity }}</div>
          </div>
        </div>
        <div class="order-footer">
          Total: <span>${{ order.total_amount }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getOrders } from '../api/order'

const orders = ref([])

onMounted(async () => {
  try {
    const res = await getOrders()
    orders.value = res.data.data || []
  } catch (e) {
    console.error(e)
  }
})

const getStatusText = (status) => {
  const map = { 0: 'Unpaid', 1: 'Paid', 2: 'Shipped', 3: 'Completed' }
  return map[status] || 'Unknown'
}
</script>

<style scoped>
.orders-page { min-height: 100vh; background: #f7f8fa; padding-bottom: 20px; }
.order-list { padding: 10px; }
.order-card { background: #fff; border-radius: 8px; margin-bottom: 15px; padding: 15px; }
.order-header { display: flex; justify-content: space-between; border-bottom: 1px solid #eee; padding-bottom: 10px; margin-bottom: 10px; font-size: 13px; color: #666; }
.status { color: #ee0a24; font-weight: bold; }
.item { display: flex; gap: 10px; margin-bottom: 10px; align-items: center; }
.item img { width: 50px; height: 50px; border-radius: 4px; object-fit: cover; }
.info { flex: 1; overflow: hidden; }
.title { font-size: 13px; color: #333; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }
.specs { font-size: 12px; color: #999; background: #f5f5f5; display: inline-block; padding: 2px 4px; border-radius: 2px; margin-top: 4px; }
.price { font-size: 14px; font-weight: bold; }
.order-footer { text-align: right; border-top: 1px solid #eee; padding-top: 10px; font-size: 14px; }
.order-footer span { color: #ee0a24; font-weight: bold; font-size: 16px; }
.empty { padding-top: 100px; }
</style>
