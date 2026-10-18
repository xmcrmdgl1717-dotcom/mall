import axios from 'axios'
const http = axios.create({ baseURL: import.meta.env.VITE_API_BASE || 'https://mall-api-l3mx.onrender.com/api/v1' })

export const getActiveCoupons = () => http.get('/coupons')
export const claimCoupon = (id) => http.post(`/coupons/${id}/claim`)
export const getMyCoupons = () => http.get('/user/coupons')
