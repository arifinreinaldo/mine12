package handlers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/arifinreinaldo/mine12/database"
	"github.com/arifinreinaldo/mine12/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetAllSales retrieves all sales with customer and items
func GetAllSales(c *fiber.Ctx) error {
	var sales []models.Sale

	result := database.GetDB().
		Preload("Customer").
		Preload("SaleItems.Item").
		Order("created_at DESC").
		Find(&sales)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve sales",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    sales,
		"total":   len(sales),
	})
}

// GetSaleByID retrieves a single sale by ID
func GetSaleByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid sale ID",
		})
	}

	var sale models.Sale
	result := database.GetDB().
		Preload("Customer").
		Preload("SaleItems.Item").
		First(&sale, id)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Sale not found",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    sale,
	})
}

// CreateSale creates a new sale with items
func CreateSale(c *fiber.Ctx) error {
	var req models.CreateSaleRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Verify customer exists
	var customer models.Customer
	if err := database.GetDB().First(&customer, req.CustomerID).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Customer not found",
		})
	}

	// Start transaction
	tx := database.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Generate order number
	orderNumber := fmt.Sprintf("ORD-%d-%d", time.Now().Unix(), req.CustomerID)

	// Create sale
	sale := models.Sale{
		OrderNumber: orderNumber,
		CustomerID:  req.CustomerID,
		OrderDate:   req.OrderDate,
		Status:      models.OrderStatusPending,
		Notes:       req.Notes,
	}

	if err := tx.Create(&sale).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create sale",
		})
	}

	// Create sale items and calculate total
	var totalAmount float64
	for _, itemReq := range req.Items {
		// Get item details
		var item models.Item
		if err := tx.First(&item, itemReq.ItemID).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": fmt.Sprintf("Item with ID %d not found", itemReq.ItemID),
			})
		}

		// Check stock availability
		if item.Stock < itemReq.Quantity {
			tx.Rollback()
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": fmt.Sprintf("Insufficient stock for item %s. Available: %d, Requested: %d",
					item.Name, item.Stock, itemReq.Quantity),
			})
		}

		// Calculate subtotal
		subtotal := float64(itemReq.Quantity) * item.Price

		// Create sale item
		saleItem := models.SaleItem{
			SaleID:    sale.ID,
			ItemID:    item.ID,
			Quantity:  itemReq.Quantity,
			UnitPrice: item.Price,
			Subtotal:  subtotal,
		}

		if err := tx.Create(&saleItem).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create sale item",
			})
		}

		// Update item stock
		item.Stock -= itemReq.Quantity
		if err := tx.Save(&item).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update item stock",
			})
		}

		totalAmount += subtotal
	}

	// Update sale total
	sale.TotalAmount = totalAmount
	if err := tx.Save(&sale).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update sale total",
		})
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to complete sale",
		})
	}

	// Reload sale with relationships
	database.GetDB().
		Preload("Customer").
		Preload("SaleItems.Item").
		First(&sale, sale.ID)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Sale created successfully",
		"data":    sale,
	})
}

// UpdateSale updates sale status and notes
func UpdateSale(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid sale ID",
		})
	}

	var sale models.Sale
	result := database.GetDB().First(&sale, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Sale not found",
		})
	}

	var req models.UpdateSaleRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Update fields
	if req.Status != "" {
		sale.Status = req.Status
	}
	if req.Notes != "" {
		sale.Notes = req.Notes
	}

	result = database.GetDB().Save(&sale)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update sale",
		})
	}

	// Reload with relationships
	database.GetDB().
		Preload("Customer").
		Preload("SaleItems.Item").
		First(&sale, sale.ID)

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Sale updated successfully",
		"data":    sale,
	})
}

// DeleteSale deletes a sale (soft delete)
func DeleteSale(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid sale ID",
		})
	}

	var sale models.Sale
	result := database.GetDB().First(&sale, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Sale not found",
		})
	}

	// Restore stock before deleting
	var saleItems []models.SaleItem
	database.GetDB().Where("sale_id = ?", id).Find(&saleItems)

	tx := database.GetDB().Begin()
	for _, saleItem := range saleItems {
		var item models.Item
		if err := tx.First(&item, saleItem.ItemID).Error; err == nil {
			item.Stock += saleItem.Quantity
			tx.Save(&item)
		}
	}

	result = tx.Delete(&sale)
	if result.Error != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete sale",
		})
	}

	tx.Commit()

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Sale deleted successfully",
	})
}

// GetSalesStats retrieves sales statistics
func GetSalesStats(c *fiber.Ctx) error {
	var stats struct {
		TotalSales     int64   `json:"total_sales"`
		CompletedSales int64   `json:"completed_sales"`
		PendingSales   int64   `json:"pending_sales"`
		TotalRevenue   float64 `json:"total_revenue"`
	}

	db := database.GetDB()

	db.Model(&models.Sale{}).Count(&stats.TotalSales)
	db.Model(&models.Sale{}).Where("status = ?", models.OrderStatusCompleted).Count(&stats.CompletedSales)
	db.Model(&models.Sale{}).Where("status = ?", models.OrderStatusPending).Count(&stats.PendingSales)

	db.Model(&models.Sale{}).
		Where("status = ?", models.OrderStatusCompleted).
		Select("COALESCE(SUM(total_amount), 0)").
		Scan(&stats.TotalRevenue)

	return c.JSON(fiber.Map{
		"success": true,
		"data":    stats,
	})
}
