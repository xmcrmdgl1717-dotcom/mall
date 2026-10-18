<template>
  <el-container class="layout-container">
    <!-- 侧边栏 -->
    <el-aside width="220px" class="aside">
      <div class="logo">FRESHCO Admin</div>
      <el-menu
        active-text-color="#409EFF"
        background-color="#304156"
        text-color="#bfcbd9"
        router
        :default-active="$route.path"
      >
        <el-menu-item index="/dashboard">
          <el-icon><Odometer /></el-icon>
          <span>数据看板</span>
        </el-menu-item>
        
        <el-sub-menu index="product">
          <template #title>
            <el-icon><Goods /></el-icon>
            <span>商品中心</span>
          </template>
          <el-menu-item index="/products">商品列表</el-menu-item>
          <el-menu-item index="/categories">类目管理</el-menu-item>
        </el-sub-menu>

        <el-menu-item index="/orders">
          <el-icon><List /></el-icon>
          <span>订单管理</span>
        </el-menu-item>

        <el-menu-item index="/users">
          <el-icon><User /></el-icon>
          <span>用户管理</span>
        </el-menu-item>

        <el-sub-menu index="system">
          <template #title>
            <el-icon><Setting /></el-icon>
            <span>系统管理</span>
          </template>
          <el-menu-item index="/roles">角色权限</el-menu-item>
          <el-menu-item index="/settings">系统设置</el-menu-item>
        </el-sub-menu>
      </el-menu>
    </el-aside>

    <!-- 右侧主体 -->
    <el-container>
      <el-header class="header">
        <div class="breadcrumb">欢迎使用多商家商城后台</div>
        <el-button type="danger" link @click="logout">退出登录</el-button>
      </el-header>
      <el-main class="main">
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup>
import { useRouter } from 'vue-router'
const router = useRouter()

const logout = () => {
  localStorage.removeItem('admin_token')
  router.push('/login')
}
</script>

<style scoped>
.layout-container { height: 100vh; }
.aside { background-color: #304156; }
.logo { height: 60px; line-height: 60px; text-align: center; font-size: 18px; font-weight: bold; color: #ccff00; background: #263445; }
.header { background-color: #fff; border-bottom: 1px solid #e6e6e6; display: flex; justify-content: space-between; align-items: center; }
.main { background-color: #f0f2f5; padding: 20px; }
</style>
