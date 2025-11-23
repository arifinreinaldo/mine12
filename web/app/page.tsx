"use client"

import { useEffect, useState } from "react"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { userApi, itemApi, customerApi, salesApi } from "@/services/api"
import { Users, Package, ShoppingCart, TrendingUp, DollarSign, UserCheck } from "lucide-react"
import Link from "next/link"

export default function Dashboard() {
  const [stats, setStats] = useState({
    totalUsers: 0,
    totalItems: 0,
    activeItems: 0,
    totalStock: 0,
    totalCustomers: 0,
    totalSales: 0,
    totalRevenue: 0,
  })

  useEffect(() => {
    fetchStats()
  }, [])

  const fetchStats = async () => {
    try {
      const [users, items, customers, salesStats] = await Promise.all([
        userApi.getAll(),
        itemApi.getAll(),
        customerApi.getAll(),
        salesApi.getStats(),
      ])

      setStats({
        totalUsers: users.length,
        totalItems: items.length,
        activeItems: items.filter(item => item.is_active).length,
        totalStock: items.reduce((sum, item) => sum + item.stock, 0),
        totalCustomers: customers.length,
        totalSales: salesStats.total_sales,
        totalRevenue: salesStats.total_revenue,
      })
    } catch (error) {
      console.error("Failed to fetch stats:", error)
    }
  }

  return (
    <div className="px-4 sm:px-0">
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-gray-900">Dashboard</h1>
        <p className="mt-2 text-sm text-gray-600">
          Welcome to your sales management system
        </p>
      </div>

      <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-4">
        <Link href="/customers">
          <Card className="hover:shadow-lg transition-shadow cursor-pointer">
            <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
              <CardTitle className="text-sm font-medium">Customers</CardTitle>
              <UserCheck className="h-4 w-4 text-muted-foreground" />
            </CardHeader>
            <CardContent>
              <div className="text-2xl font-bold">{stats.totalCustomers}</div>
              <p className="text-xs text-muted-foreground">Registered customers</p>
            </CardContent>
          </Card>
        </Link>

        <Link href="/items">
          <Card className="hover:shadow-lg transition-shadow cursor-pointer">
            <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
              <CardTitle className="text-sm font-medium">Products</CardTitle>
              <Package className="h-4 w-4 text-muted-foreground" />
            </CardHeader>
            <CardContent>
              <div className="text-2xl font-bold">{stats.totalItems}</div>
              <p className="text-xs text-muted-foreground">In catalog</p>
            </CardContent>
          </Card>
        </Link>

        <Link href="/sales">
          <Card className="hover:shadow-lg transition-shadow cursor-pointer">
            <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
              <CardTitle className="text-sm font-medium">Total Sales</CardTitle>
              <ShoppingCart className="h-4 w-4 text-muted-foreground" />
            </CardHeader>
            <CardContent>
              <div className="text-2xl font-bold">{stats.totalSales}</div>
              <p className="text-xs text-muted-foreground">Orders placed</p>
            </CardContent>
          </Card>
        </Link>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Revenue</CardTitle>
            <DollarSign className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">${stats.totalRevenue.toFixed(2)}</div>
            <p className="text-xs text-muted-foreground">Total earnings</p>
          </CardContent>
        </Card>
      </div>

      <div className="mt-8 grid gap-6 md:grid-cols-2">
        <Card>
          <CardHeader>
            <CardTitle>Quick Actions</CardTitle>
            <CardDescription>Common tasks and operations</CardDescription>
          </CardHeader>
          <CardContent className="space-y-2">
            <Link
              href="/users"
              className="block p-3 rounded-lg border hover:bg-gray-50 transition-colors"
            >
              <div className="font-medium">Manage Users</div>
              <div className="text-sm text-gray-500">View and edit user accounts</div>
            </Link>
            <Link
              href="/items"
              className="block p-3 rounded-lg border hover:bg-gray-50 transition-colors"
            >
              <div className="font-medium">Manage Items</div>
              <div className="text-sm text-gray-500">Update product inventory</div>
            </Link>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>System Info</CardTitle>
            <CardDescription>Application details</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="space-y-2">
              <div className="flex justify-between">
                <span className="text-sm text-gray-500">Version</span>
                <span className="text-sm font-medium">1.0.0</span>
              </div>
              <div className="flex justify-between">
                <span className="text-sm text-gray-500">Framework</span>
                <span className="text-sm font-medium">Next.js 14</span>
              </div>
              <div className="flex justify-between">
                <span className="text-sm text-gray-500">Backend</span>
                <span className="text-sm font-medium">Go Fiber</span>
              </div>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
