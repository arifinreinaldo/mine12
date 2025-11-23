package handlers

import (
	"strconv"

	"github.com/arifinreinaldo/mine12/database"
	"github.com/arifinreinaldo/mine12/models"
	"github.com/gofiber/fiber/v2"
)

// GetAllItems retrieves all items
func GetAllItems(c *fiber.Ctx) error {
	var items []models.Item

	result := database.GetDB().Find(&items)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve items",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    items,
		"total":   len(items),
	})
}

// GetItemByID retrieves a single item by ID
func GetItemByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid item ID",
		})
	}

	var item models.Item
	result := database.GetDB().First(&item, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Item not found",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    item,
	})
}

// CreateItem creates a new item
func CreateItem(c *fiber.Ctx) error {
	var req models.CreateItemRequest

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

	item := models.Item{
		Code:        req.Code,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Unit:        req.Unit,
		IsActive:    true,
	}

	result := database.GetDB().Create(&item)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create item",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Item created successfully",
		"data":    item,
	})
}

// UpdateItem updates an existing item
func UpdateItem(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid item ID",
		})
	}

	var item models.Item
	result := database.GetDB().First(&item, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Item not found",
		})
	}

	var req models.UpdateItemRequest
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
	if req.Code != "" {
		item.Code = req.Code
	}
	if req.Name != "" {
		item.Name = req.Name
	}
	if req.Description != "" {
		item.Description = req.Description
	}
	if req.Price != nil {
		item.Price = *req.Price
	}
	if req.Stock != nil {
		item.Stock = *req.Stock
	}
	if req.Unit != "" {
		item.Unit = req.Unit
	}
	if req.IsActive != nil {
		item.IsActive = *req.IsActive
	}

	result = database.GetDB().Save(&item)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update item",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Item updated successfully",
		"data":    item,
	})
}

// DeleteItem deletes an item (soft delete)
func DeleteItem(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid item ID",
		})
	}

	var item models.Item
	result := database.GetDB().First(&item, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Item not found",
		})
	}

	result = database.GetDB().Delete(&item)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete item",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Item deleted successfully",
	})
}
