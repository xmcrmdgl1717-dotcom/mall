import axios from 'axios'
const http = axios.create({ baseURL: import.meta.env.VITE_API_BASE || 'https://mall-api-l3mx.onrender.com/api/v1' })

export const getAddresses = () => http.get('/addresses')
export const createAddress = (data) => http.post('/addresses', data)
export const deleteAddress = (id) => http.delete(`/addresses/${id}`)
