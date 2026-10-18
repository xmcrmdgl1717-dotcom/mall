import axios from 'axios'
const http = axios.create({ baseURL: import.meta.env.VITE_API_BASE || 'http://localhost:8080/api/v1' })
export const getCategories = () => http.get('/categories')
