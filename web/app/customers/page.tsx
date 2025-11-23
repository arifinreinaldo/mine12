"use client"

import { useEffect, useState } from "react"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogDescription } from "@/components/ui/dialog"
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table"
import { useToast } from "@/components/ui/use-toast"
import { customerApi } from "@/services/api"
import type { Customer, CreateCustomerRequest, UpdateCustomerRequest } from "@/types"
import { CustomerForm } from "@/components/customers/customer-form"
import { Pencil, Trash2, Plus } from "lucide-react"

export default function CustomersPage() {
  const [customers, setCustomers] = useState<Customer[]>([])
  const [loading, setLoading] = useState(true)
  const [dialogOpen, setDialogOpen] = useState(false)
  const [editingCustomer, setEditingCustomer] = useState<Customer | null>(null)
  const { toast } = useToast()

  useEffect(() => {
    fetchCustomers()
  }, [])

  const fetchCustomers = async () => {
    try {
      setLoading(true)
      const data = await customerApi.getAll()
      setCustomers(data)
    } catch (error) {
      toast({ title: "Error", description: "Failed to fetch customers", variant: "destructive" })
    } finally {
      setLoading(false)
    }
  }

  const handleCreate = async (data: CreateCustomerRequest | UpdateCustomerRequest) => {
    try {
      await customerApi.create(data as CreateCustomerRequest)
      toast({ title: "Success", description: "Customer created successfully" })
      setDialogOpen(false)
      fetchCustomers()
    } catch (error) {
      toast({ title: "Error", description: "Failed to create customer", variant: "destructive" })
    }
  }

  const handleUpdate = async (data: CreateCustomerRequest | UpdateCustomerRequest) => {
    if (!editingCustomer) return
    try {
      await customerApi.update(editingCustomer.id, data as UpdateCustomerRequest)
      toast({ title: "Success", description: "Customer updated successfully" })
      setDialogOpen(false)
      setEditingCustomer(null)
      fetchCustomers()
    } catch (error) {
      toast({ title: "Error", description: "Failed to update customer", variant: "destructive" })
    }
  }

  const handleDelete = async (id: number) => {
    if (!confirm("Delete this customer?")) return
    try {
      await customerApi.delete(id)
      toast({ title: "Success", description: "Customer deleted" })
      fetchCustomers()
    } catch (error) {
      toast({ title: "Error", description: "Failed to delete customer", variant: "destructive" })
    }
  }

  return (
    <div className="px-4 sm:px-0">
      <Card>
        <CardHeader>
          <div className="flex justify-between items-center">
            <div>
              <CardTitle>Customers</CardTitle>
              <CardDescription>Manage customer information</CardDescription>
            </div>
            <Button onClick={() => { setEditingCustomer(null); setDialogOpen(true) }}>
              <Plus className="mr-2 h-4 w-4" />Add Customer
            </Button>
          </div>
        </CardHeader>
        <CardContent>
          {loading ? <div className="text-center py-8">Loading...</div> : (
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>Name</TableHead>
                  <TableHead>Email</TableHead>
                  <TableHead>Phone</TableHead>
                  <TableHead>Type</TableHead>
                  <TableHead>Status</TableHead>
                  <TableHead className="text-right">Actions</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {customers.map((customer) => (
                  <TableRow key={customer.id}>
                    <TableCell className="font-medium">{customer.name}</TableCell>
                    <TableCell>{customer.email}</TableCell>
                    <TableCell>{customer.phone}</TableCell>
                    <TableCell><span className="capitalize">{customer.type}</span></TableCell>
                    <TableCell>
                      <span className={`px-2 py-1 rounded-full text-xs ${customer.is_active ? "bg-green-100 text-green-800" : "bg-gray-100 text-gray-800"}`}>
                        {customer.is_active ? "Active" : "Inactive"}
                      </span>
                    </TableCell>
                    <TableCell className="text-right">
                      <div className="flex justify-end gap-2">
                        <Button variant="outline" size="sm" onClick={() => { setEditingCustomer(customer); setDialogOpen(true) }}>
                          <Pencil className="h-4 w-4" />
                        </Button>
                        <Button variant="destructive" size="sm" onClick={() => handleDelete(customer.id)}>
                          <Trash2 className="h-4 w-4" />
                        </Button>
                      </div>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          )}
        </CardContent>
      </Card>

      <Dialog open={dialogOpen} onOpenChange={setDialogOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>{editingCustomer ? "Edit Customer" : "Create New Customer"}</DialogTitle>
            <DialogDescription>Customer information</DialogDescription>
          </DialogHeader>
          <CustomerForm customer={editingCustomer || undefined} onSubmit={editingCustomer ? handleUpdate : handleCreate} onCancel={() => { setDialogOpen(false); setEditingCustomer(null) }} isEditing={!!editingCustomer} />
        </DialogContent>
      </Dialog>
    </div>
  )
}
