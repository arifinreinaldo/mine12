"use client"

import { useEffect, useState } from "react"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table"
import { useToast } from "@/components/ui/use-toast"
import { salesApi } from "@/services/api"
import type { Sale } from "@/types"
import { DollarSign, ShoppingBag } from "lucide-react"
import Link from "next/link"

export default function SalesPage() {
  const [sales, setSales] = useState<Sale[]>([])
  const [loading, setLoading] = useState(true)
  const { toast } = useToast()

  useEffect(() => {
    fetchSales()
  }, [])

  const fetchSales = async () => {
    try {
      setLoading(true)
      const data = await salesApi.getAll()
      setSales(data)
    } catch (error) {
      toast({ title: "Error", description: "Failed to fetch sales", variant: "destructive" })
    } finally {
      setLoading(false)
    }
  }

  const formatCurrency = (value: number) => {
    return new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD' }).format(value)
  }

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString()
  }

  const getStatusBadge = (status: string) => {
    const colors: Record<string, string> = {
      pending: "bg-yellow-100 text-yellow-800",
      completed: "bg-green-100 text-green-800",
      cancelled: "bg-red-100 text-red-800",
      refunded: "bg-gray-100 text-gray-800",
    }
    return (
      <span className={`px-2 py-1 rounded-full text-xs font-medium ${colors[status]}`}>
        {status.charAt(0).toUpperCase() + status.slice(1)}
      </span>
    )
  }

  return (
    <div className="px-4 sm:px-0">
      <Card>
        <CardHeader>
          <div className="flex justify-between items-center">
            <div>
              <CardTitle>Sales Orders</CardTitle>
              <CardDescription>View all sales transactions</CardDescription>
            </div>
            <Link href="/sales/new">
              <Button>
                <ShoppingBag className="mr-2 h-4 w-4" />New Sale
              </Button>
            </Link>
          </div>
        </CardHeader>
        <CardContent>
          {loading ? <div className="text-center py-8">Loading...</div> : (
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>Order #</TableHead>
                  <TableHead>Customer</TableHead>
                  <TableHead>Date</TableHead>
                  <TableHead>Items</TableHead>
                  <TableHead>Total</TableHead>
                  <TableHead>Status</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {sales.map((sale) => (
                  <TableRow key={sale.id}>
                    <TableCell className="font-medium">{sale.order_number}</TableCell>
                    <TableCell>{sale.customer?.name || `Customer #${sale.customer_id}`}</TableCell>
                    <TableCell>{formatDate(sale.order_date)}</TableCell>
                    <TableCell>{sale.items?.length || 0} items</TableCell>
                    <TableCell className="font-semibold">{formatCurrency(sale.total_amount)}</TableCell>
                    <TableCell>{getStatusBadge(sale.status)}</TableCell>
                  </TableRow>
                ))}
                {sales.length === 0 && (
                  <TableRow>
                    <TableCell colSpan={6} className="text-center text-gray-500 py-8">
                      No sales yet. Create your first sale!
                    </TableCell>
                  </TableRow>
                )}
              </TableBody>
            </Table>
          )}
        </CardContent>
      </Card>
    </div>
  )
}
