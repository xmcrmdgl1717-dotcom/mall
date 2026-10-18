<template>
  <div class="coupon-center">
    <van-nav-bar title="Coupon Center" left-arrow @click-left="$router.back()" fixed placeholder />
    
    <div class="coupon-list">
      <div class="coupon-card" v-for="c in list" :key="c.id">
        <div class="left">
          <div class="amount"><span>$</span>{{ c.amount }}</div>
          <div class="condition" v-if="c.min_amount > 0">满 ${{ c.min_amount }} 可用</div>
          <div class="condition" v-else>无门槛</div>
        </div>
        <div class="right">
          <h3>{{ c.name }}</h3>
          <p>有效期至: {{ new Date(c.end_time).toLocaleDateString() }}</p>
          <van-button type="danger" size="small" round @click="onClaim(c.id)">立即领取</van-button>
        </div>
      </div>
    </div>
    
    <van-empty v-if="list.length === 0 && !loading" description="暂无优惠券" />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getActiveCoupons, claimCoupon } from '../api/coupon'
import { showToast } from 'vant'

const list = ref([])
const loading = ref(false)

const loadData = async () => {
  loading.value = true
  try {
    const res = await getActiveCoupons()
    list.value = res.data.data || []
  } finally {
    loading.value = false
  }
}

const onClaim = async (id) => {
  try {
    const res = await claimCoupon(id)
    if (res.data.code === 0) {
      showToast('领取成功')
      loadData()
    } else {
      showToast(res.data.msg)
    }
  } catch (e) {
    showToast(e.response?.data?.msg || '领取失败')
  }
}

onMounted(() => { loadData() })
</script>

<style scoped>
.coupon-center { min-height: 100vh; background: #f7f8fa; padding: 10px; }
.coupon-card { display: flex; background: #fff; border-radius: 8px; margin-bottom: 10px; overflow: hidden; }
.left { width: 120px; background: #ee0a24; color: #fff; display: flex; flex-direction: column; justify-content: center; align-items: center; padding: 15px 0; }
.amount { font-size: 28px; font-weight: bold; }
.amount span { font-size: 16px; }
.condition { font-size: 12px; margin-top: 5px; }
.right { flex: 1; padding: 15px; display: flex; flex-direction: column; justify-content: space-between; }
.right h3 { margin: 0; font-size: 16px; color: #333; }
.right p { margin: 0; font-size: 12px; color: #999; }
</style>
