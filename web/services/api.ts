import axios from 'axios'
import type {
  User,
  CreateUserRequest,
  UpdateUserRequest,
  Item,
  CreateItemRequest,
  UpdateItemRequest,
  Customer,
  CreateCustomerRequest,
  UpdateCustomerRequest,
  Sale,
  CreateSaleRequest,
  UpdateSaleRequest,
  SalesStats,
  ApiResponse,
} from '@/types'

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:3000'

const api = axios.create({
  baseURL: `${API_URL}/api/v1`,
  headers: {
    'Content-Type': 'application/json',
  },
})

// User API
export const userApi = {
  getAll: async (): Promise<User[]> => {
    const { data } = await api.get<ApiResponse<User[]>>('/users')
    return data.data || []
  },

  getById: async (id: number): Promise<User> => {
    const { data } = await api.get<ApiResponse<User>>(`/users/${id}`)
    if (!data.data) throw new Error('User not found')
    return data.data
  },

  create: async (userData: CreateUserRequest): Promise<User> => {
    const { data } = await api.post<ApiResponse<User>>('/users', userData)
    if (!data.data) throw new Error('Failed to create user')
    return data.data
  },

  update: async (id: number, userData: UpdateUserRequest): Promise<User> => {
    const { data } = await api.put<ApiResponse<User>>(`/users/${id}`, userData)
    if (!data.data) throw new Error('Failed to update user')
    return data.data
  },

  delete: async (id: number): Promise<void> => {
    await api.delete(`/users/${id}`)
  },
}

// Item API
export const itemApi = {
  getAll: async (): Promise<Item[]> => {
    const { data } = await api.get<ApiResponse<Item[]>>('/items')
    return data.data || []
  },

  getById: async (id: number): Promise<Item> => {
    const { data } = await api.get<ApiResponse<Item>>(`/items/${id}`)
    if (!data.data) throw new Error('Item not found')
    return data.data
  },

  create: async (itemData: CreateItemRequest): Promise<Item> => {
    const { data } = await api.post<ApiResponse<Item>>('/items', itemData)
    if (!data.data) throw new Error('Failed to create item')
    return data.data
  },

  update: async (id: number, itemData: UpdateItemRequest): Promise<Item> => {
    const { data } = await api.put<ApiResponse<Item>>(`/items/${id}`, itemData)
    if (!data.data) throw new Error('Failed to update item')
    return data.data
  },

  delete: async (id: number): Promise<void> => {
    await api.delete(`/items/${id}`)
  },
}

// Customer API
export const customerApi = {
  getAll: async (): Promise<Customer[]> => {
    const { data } = await api.get<ApiResponse<Customer[]>>('/customers')
    return data.data || []
  },

  getById: async (id: number): Promise<Customer> => {
    const { data } = await api.get<ApiResponse<Customer>>(`/customers/${id}`)
    if (!data.data) throw new Error('Customer not found')
    return data.data
  },

  create: async (customerData: CreateCustomerRequest): Promise<Customer> => {
    const { data } = await api.post<ApiResponse<Customer>>('/customers', customerData)
    if (!data.data) throw new Error('Failed to create customer')
    return data.data
  },

  update: async (id: number, customerData: UpdateCustomerRequest): Promise<Customer> => {
    const { data } = await api.put<ApiResponse<Customer>>(`/customers/${id}`, customerData)
    if (!data.data) throw new Error('Failed to update customer')
    return data.data
  },

  delete: async (id: number): Promise<void> => {
    await api.delete(`/customers/${id}`)
  },
}

// Sales API
export const salesApi = {
  getAll: async (): Promise<Sale[]> => {
    const { data } = await api.get<ApiResponse<Sale[]>>('/sales')
    return data.data || []
  },

  getById: async (id: number): Promise<Sale> => {
    const { data } = await api.get<ApiResponse<Sale>>(`/sales/${id}`)
    if (!data.data) throw new Error('Sale not found')
    return data.data
  },

  create: async (saleData: CreateSaleRequest): Promise<Sale> => {
    const { data } = await api.post<ApiResponse<Sale>>('/sales', saleData)
    if (!data.data) throw new Error('Failed to create sale')
    return data.data
  },

  update: async (id: number, saleData: UpdateSaleRequest): Promise<Sale> => {
    const { data } = await api.put<ApiResponse<Sale>>(`/sales/${id}`, saleData)
    if (!data.data) throw new Error('Failed to update sale')
    return data.data
  },

  delete: async (id: number): Promise<void> => {
    await api.delete(`/sales/${id}`)
  },

  getStats: async (): Promise<SalesStats> => {
    const { data } = await api.get<ApiResponse<SalesStats>>('/sales/stats')
    if (!data.data) throw new Error('Failed to fetch stats')
    return data.data
  },
}

export default api
