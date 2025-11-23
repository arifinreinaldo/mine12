export type AccessLevel = "admin" | "manager" | "staff" | "readonly"

export interface User {
  id: number
  username: string
  email: string
  full_name: string
  access_level: AccessLevel
  is_active: boolean
  created_at: string
  updated_at: string
}

export interface CreateUserRequest {
  username: string
  email: string
  password: string
  full_name: string
  access_level: AccessLevel
}

export interface UpdateUserRequest {
  username?: string
  email?: string
  full_name?: string
  access_level?: AccessLevel
  is_active?: boolean
}

export interface Item {
  id: number
  code: string
  name: string
  description: string
  price: number
  stock: number
  unit: string
  is_active: boolean
  created_at: string
  updated_at: string
}

export interface CreateItemRequest {
  code: string
  name: string
  description: string
  price: number
  stock: number
  unit: string
}

export interface UpdateItemRequest {
  code?: string
  name?: string
  description?: string
  price?: number
  stock?: number
  unit?: string
  is_active?: boolean
}

export interface ApiResponse<T> {
  success: boolean
  data?: T
  message?: string
  error?: string
  total?: number
}
