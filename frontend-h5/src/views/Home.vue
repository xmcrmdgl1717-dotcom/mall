<template>
  <div class="home">
    <van-search v-model="keyword" placeholder="搜索商品" @search="load" />
    <van-tabs v-model:active="sort" @change="load">
      <van-tab title="综合" name="default" />
      <van-tab title="销量" name="sales" />
      <van-tab title="价格↑" name="price_asc" />
      <van-tab title="价格↓" name="price_desc" />
    </van-tabs>
    <van-list v-model:loading="loading" :finished="finished" @load="load">
      <div class="grid">
        <div v-for="p in list" :key="p.id" class="card">
          <img :src="p.main_image" />
          <div class="title">{{ p.title }}</div>
          <div class="price">${{ p.min_price }}</div>
          <div class="sales">Sold {{ p.total_sales }}</div>
        </div>
      </div>
    </van-list>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { getProducts } from '../api/product'

const keyword = ref('')
const sort = ref('default')
const list = ref([])
const loading = ref(false)
const finished = ref(false)
let page = 1

async function load() {
  loading.value = true
  const { data } = await getProducts({ sort: sort.value, page, size: 20 })
  const rows = data.data.list || []
  list.value = page === 1 ? rows : [...list.value, ...rows]
  finished.value = rows.length < 20
  page++
  loading.value = false
}
</script>
