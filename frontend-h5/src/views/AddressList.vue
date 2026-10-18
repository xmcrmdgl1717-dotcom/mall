<template>
  <div class="address-page">
    <van-nav-bar title="My Addresses" left-arrow @click-left="$router.back()" fixed placeholder />

    <div class="address-list" v-if="list.length > 0">
      <div class="address-card" v-for="item in list" :key="item.id">
        <div class="info">
          <div class="user">
            <span class="name">{{ item.name }}</span>
            <span class="phone">{{ item.phone }}</span>
            <van-tag v-if="item.is_default === 1" type="danger" style="margin-left: 10px;">Default</van-tag>
          </div>
          <div class="detail">{{ item.province }} {{ item.city }} {{ item.detail }}</div>
        </div>
        <van-icon name="delete-o" size="20" color="#999" @click="onDelete(item.id)" />
      </div>
    </div>
    <van-empty v-else description="No address found" />

    <!-- 底部悬浮添加按钮 -->
    <div class="bottom-btn">
      <van-button type="primary" block round @click="showAdd = true">+ Add New Address</van-button>
    </div>

    <!-- 新增地址弹窗 -->
    <van-popup v-model:show="showAdd" position="bottom" round style="height: 70%;">
      <div class="popup-content">
        <h3>Add New Address</h3>
        <van-form @submit="onSubmit">
          <van-cell-group inset>
            <van-field v-model="form.name" label="Name" placeholder="Receiver name" :rules="[{ required: true }]" />
            <van-field v-model="form.phone" label="Phone" placeholder="Phone number" :rules="[{ required: true }]" />
            <van-field v-model="form.province" label="Province" placeholder="Province" />
            <van-field v-model="form.city" label="City" placeholder="City" />
            <van-field v-model="form.detail" label="Detail" type="textarea" rows="2" placeholder="Street, building, etc." :rules="[{ required: true }]" />
            <van-field name="switch" label="Set as Default">
              <template #input>
                <van-switch v-model="form.is_default" :active-value="1" :inactive-value="0" size="20" />
              </template>
            </van-field>
          </van-cell-group>
          <div style="margin: 16px;">
            <van-button round block type="primary" native-type="submit">Save</van-button>
          </div>
        </van-form>
      </div>
    </van-popup>
  </div>
</template>

<script setup>
import { ref, onMounted, reactive } from 'vue'
import { getAddresses, createAddress, deleteAddress } from '../api/address'
import { showToast, showConfirmDialog } from 'vant'

const list = ref([])
const showAdd = ref(false)
const form = reactive({ name: '', phone: '', province: '', city: '', detail: '', is_default: 1 })

const loadData = async () => {
  const res = await getAddresses()
  list.value = res.data.data || []
}

const onSubmit = async () => {
  try {
    await createAddress(form)
    showToast('Saved successfully')
    showAdd.value = false
    // 重置表单
    form.name = ''; form.phone = ''; form.detail = ''
    loadData()
  } catch (e) {
    showToast('Failed to save')
  }
}

const onDelete = async (id) => {
  try {
    await showConfirmDialog({ title: 'Delete', message: 'Delete this address?' })
    await deleteAddress(id)
    showToast('Deleted')
    loadData()
  } catch (e) {}
}

onMounted(() => { loadData() })
</script>

<style scoped>
.address-page { min-height: 100vh; background: #f7f8fa; padding-bottom: 80px; }
.address-list { padding: 10px; }
.address-card { background: #fff; padding: 15px; border-radius: 8px; margin-bottom: 10px; display: flex; justify-content: space-between; align-items: center; }
.info { flex: 1; }
.user { margin-bottom: 5px; }
.name { font-weight: bold; margin-right: 10px; }
.phone { color: #666; }
.detail { font-size: 13px; color: #999; line-height: 1.4; }
.bottom-btn { position: fixed; bottom: 0; left: 0; width: 100%; padding: 10px; background: #fff; box-sizing: border-box; }
.popup-content { padding: 20px; }
</style>
