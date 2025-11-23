"use client"

import { useState } from "react"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Switch } from "@/components/ui/switch"
import type { Item, CreateItemRequest, UpdateItemRequest } from "@/types"

interface ItemFormProps {
  item?: Item
  onSubmit: (data: CreateItemRequest | UpdateItemRequest) => Promise<void>
  onCancel: () => void
  isEditing?: boolean
}

export function ItemForm({ item, onSubmit, onCancel, isEditing = false }: ItemFormProps) {
  const [formData, setFormData] = useState({
    code: item?.code || "",
    name: item?.name || "",
    description: item?.description || "",
    price: item?.price || 0,
    stock: item?.stock || 0,
    unit: item?.unit || "",
    is_active: item?.is_active ?? true,
  })
  const [loading, setLoading] = useState(false)

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setLoading(true)

    try {
      if (isEditing) {
        const updateData: UpdateItemRequest = {
          code: formData.code,
          name: formData.name,
          description: formData.description,
          price: formData.price,
          stock: formData.stock,
          unit: formData.unit,
          is_active: formData.is_active,
        }
        await onSubmit(updateData)
      } else {
        const createData: CreateItemRequest = {
          code: formData.code,
          name: formData.name,
          description: formData.description,
          price: formData.price,
          stock: formData.stock,
          unit: formData.unit,
        }
        await onSubmit(createData)
      }
    } finally {
      setLoading(false)
    }
  }

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <div className="space-y-2">
        <Label htmlFor="code">Item Code</Label>
        <Input
          id="code"
          value={formData.code}
          onChange={(e) => setFormData({ ...formData, code: e.target.value })}
          required
          minLength={2}
          placeholder="ITEM001"
        />
      </div>

      <div className="space-y-2">
        <Label htmlFor="name">Name</Label>
        <Input
          id="name"
          value={formData.name}
          onChange={(e) => setFormData({ ...formData, name: e.target.value })}
          required
          placeholder="Product name"
        />
      </div>

      <div className="space-y-2">
        <Label htmlFor="description">Description</Label>
        <Input
          id="description"
          value={formData.description}
          onChange={(e) => setFormData({ ...formData, description: e.target.value })}
          placeholder="Product description"
        />
      </div>

      <div className="grid grid-cols-2 gap-4">
        <div className="space-y-2">
          <Label htmlFor="price">Price</Label>
          <Input
            id="price"
            type="number"
            step="0.01"
            min="0"
            value={formData.price}
            onChange={(e) => setFormData({ ...formData, price: parseFloat(e.target.value) })}
            required
          />
        </div>

        <div className="space-y-2">
          <Label htmlFor="stock">Stock</Label>
          <Input
            id="stock"
            type="number"
            min="0"
            value={formData.stock}
            onChange={(e) => setFormData({ ...formData, stock: parseInt(e.target.value) })}
            required
          />
        </div>
      </div>

      <div className="space-y-2">
        <Label htmlFor="unit">Unit</Label>
        <Input
          id="unit"
          value={formData.unit}
          onChange={(e) => setFormData({ ...formData, unit: e.target.value })}
          required
          placeholder="pcs, kg, liter, etc."
        />
      </div>

      {isEditing && (
        <div className="flex items-center space-x-2">
          <Switch
            id="is_active"
            checked={formData.is_active}
            onCheckedChange={(checked) =>
              setFormData({ ...formData, is_active: checked })
            }
          />
          <Label htmlFor="is_active">Active</Label>
        </div>
      )}

      <div className="flex gap-2 pt-4">
        <Button type="submit" disabled={loading}>
          {loading ? "Saving..." : isEditing ? "Update Item" : "Create Item"}
        </Button>
        <Button type="button" variant="outline" onClick={onCancel}>
          Cancel
        </Button>
      </div>
    </form>
  )
}
