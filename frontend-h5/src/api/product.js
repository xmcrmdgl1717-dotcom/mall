import axios from 'axios'

const http = axios.create({ baseURL: import.meta.env.VITE_API_BASE || 'http://localhost:8080/api/v1' })

export const getProducts = (params) => http.get('/products', { params })

// 获取商品详情
export const getProductDetail = (id) => http.get(`/products/${id}`)
