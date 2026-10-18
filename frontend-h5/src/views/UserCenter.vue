<template>
  <div class="user-center">
    <div class="header">
      <van-icon name="user-circle-o" size="60" color="#fff" />
      <div class="info">
        <h2>Test User</h2>
        <p>ID: 100001</p>
      </div>
    </div>

    <div class="menu-section">
      <van-cell title="My Orders" is-link to="/orders" icon="orders-o" />
      <van-cell title="Coupon Center" is-link to="/coupons" icon="coupon-o" />
      <van-cell title="My Coupons" is-link to="/my-coupons" icon="gold-coin-o" />
      <van-cell title="My Addresses" is-link to="/addresses" icon="location-o" />
      <van-cell title="Customer Service" is-link icon="service-o" />
    </div>
    
    <div class="menu-section" style="margin-top: 10px;">
      <van-cell title="Settings" is-link icon="setting-o" />
    </div>
  </div>
</template>

<style scoped>
.user-center { min-height: 100vh; background: #f7f8fa; }
.header { background: #333; padding: 40px 20px; display: flex; align-items: center; gap: 15px; color: #fff; }
.info h2 { margin: 0 0 5px; font-size: 18px; }
.info p { margin: 0; font-size: 12px; color: #ccc; }
.menu-section { background: #fff; }
</style>
