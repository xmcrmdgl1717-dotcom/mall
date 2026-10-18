<template>
  <div class="my-coupons">
    <van-nav-bar title="My Coupons" left-arrow @click-left="$router.back()" fixed placeholder />
    <div class="list">
      <div class="card" v-for="uc in list" :key="uc.id" :class="{ disabled: uc.status !== 0 }">
        <div class="left">
          <div class="amount">${{ uc.coupon.amount }}</div>
          <div class="condition" v-if="uc.coupon.min_amount > 0">满 ${{ uc.coupon.min_amount }} 可用</div>
        </div>
        <div class="right">
          <h3>{{ uc.coupon.name }}</h3>
          <p>有效期至: {{ new Date(uc.coupon.end_time).toLocaleDateString() }}</p>
          <van-tag :type="uc.status === 0 ? 'success' : 'default'">
            {{ uc.status === 0 ? '未使用' : '已失效' }}
          </van-tag>
        </div>
      </div>
    </div>
    <van-empty v-if="list.length === 0" description="You have no coupons" />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getMyCoupons } from '../api/coupon'

const list = ref([])
onMounted(async () => {
  const res = await getMyCoupons()
  list.value = res.data.data || []
})
</script>

<style scoped>
.my-coupons { min-height: 100vh; background: #f7f8fa; padding: 10px; }
.card { display: flex; background: #fff; border-radius: 8px; margin-bottom: 10px; }
.card.disabled { opacity: 0.5; filter: grayscale(1); }
.left { width: 100px; background: #ff976a; color: #fff; display: flex; flex-direction: column; justify-content: center; align-items: center; border-radius: 8px 0 0 8px; }
.amount { font-size: 24px; font-weight: bold; }
.condition { font-size: 11px; margin-top: 5px; }
.right { flex: 1; padding: 15px; }
.right h3 { margin: 0 0 5px; font-size: 15px; color: #333; }
.right p { margin: 0 0 10px; font-size: 12px; color: #999; }
</style>
