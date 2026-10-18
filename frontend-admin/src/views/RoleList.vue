<template>
  <el-card>
    <div style="margin-bottom: 20px; display: flex; justify-content: space-between;">
      <h3>角色与权限管理</h3>
      <el-button type="primary">+ 新增角色 (开发中)</el-button>
    </div>
    
    <el-table :data="tableData" border style="width: 100%">
      <el-table-column prop="id" label="角色ID" width="100" />
      <el-table-column prop="name" label="角色名称" />
      <el-table-column prop="description" label="角色描述" />
      <el-table-column label="操作" width="150">
        <template #default>
          <el-button size="small" type="primary" link>分配权限</el-button>
        </template>
      </el-table-column>
    </el-table>
    
    <el-alert title="说明：当前系统内置了超级管理员角色。未来可在此处配置运营、客服等不同角色所能看到的菜单。" type="info" style="margin-top: 20px;" />
  </el-card>
</template>

<script setup>
import { ref, onMounted } from 'vue'
// 此处暂未写后端接口，直接使用 mock 数据展示
const tableData = ref([
  { id: 1, name: '超级管理员', description: '拥有系统所有权限' }
])
</script>
