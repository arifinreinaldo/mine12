import axios from 'axios'
import type {
  User,
  CreateUserRequest,
  UpdateUserRequest,
  Item,
  CreateItemRequest,
  UpdateItemRequest,
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

export default api
