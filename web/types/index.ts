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

export type CustomerType = "retail" | "wholesale" | "corporate"

export interface Customer {
  id: number
  name: string
  email: string
  phone: string
  address: string
  city: string
  country: string
  type: CustomerType
  is_active: boolean
  created_at: string
  updated_at: string
}

export interface CreateCustomerRequest {
  name: string
  email: string
  phone: string
  address: string
  city: string
  country: string
  type: CustomerType
}

export interface UpdateCustomerRequest {
  name?: string
  email?: string
  phone?: string
  address?: string
  city?: string
  country?: string
  type?: CustomerType
  is_active?: boolean
}

export type OrderStatus = "pending" | "completed" | "cancelled" | "refunded"

export interface SaleItem {
  id: number
  sale_id: number
  item_id: number
  item?: Item
  quantity: number
  unit_price: number
  subtotal: number
  created_at: string
  updated_at: string
}

export interface Sale {
  id: number
  order_number: string
  customer_id: number
  customer?: Customer
  order_date: string
  status: OrderStatus
  total_amount: number
  notes: string
  items?: SaleItem[]
  created_at: string
  updated_at: string
}

export interface CreateSaleItem {
  item_id: number
  quantity: number
}

export interface CreateSaleRequest {
  customer_id: number
  order_date: string
  notes: string
  items: CreateSaleItem[]
}

export interface UpdateSaleRequest {
  status?: OrderStatus
  notes?: string
}

export interface SalesStats {
  total_sales: number
  completed_sales: number
  pending_sales: number
  total_revenue: number
}
