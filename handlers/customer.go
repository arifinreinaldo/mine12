package handlers

import (
	"strconv"

	"github.com/arifinreinaldo/mine12/database"
	"github.com/arifinreinaldo/mine12/models"
	"github.com/gofiber/fiber/v2"
)

// GetAllCustomers retrieves all customers
func GetAllCustomers(c *fiber.Ctx) error {
	var customers []models.Customer

	result := database.GetDB().Find(&customers)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve customers",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    customers,
		"total":   len(customers),
	})
}

// GetCustomerByID retrieves a single customer by ID
func GetCustomerByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid customer ID",
		})
	}

	var customer models.Customer
	result := database.GetDB().First(&customer, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Customer not found",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    customer,
	})
}

// CreateCustomer creates a new customer
func CreateCustomer(c *fiber.Ctx) error {
	var req models.CreateCustomerRequest

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

	customer := models.Customer{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Address:  req.Address,
		City:     req.City,
		Country:  req.Country,
		Type:     req.Type,
		IsActive: true,
	}

	result := database.GetDB().Create(&customer)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create customer",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Customer created successfully",
		"data":    customer,
	})
}

// UpdateCustomer updates an existing customer
func UpdateCustomer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid customer ID",
		})
	}

	var customer models.Customer
	result := database.GetDB().First(&customer, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Customer not found",
		})
	}

	var req models.UpdateCustomerRequest
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
	if req.Name != "" {
		customer.Name = req.Name
	}
	if req.Email != "" {
		customer.Email = req.Email
	}
	if req.Phone != "" {
		customer.Phone = req.Phone
	}
	if req.Address != "" {
		customer.Address = req.Address
	}
	if req.City != "" {
		customer.City = req.City
	}
	if req.Country != "" {
		customer.Country = req.Country
	}
	if req.Type != "" {
		customer.Type = req.Type
	}
	if req.IsActive != nil {
		customer.IsActive = *req.IsActive
	}

	result = database.GetDB().Save(&customer)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update customer",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Customer updated successfully",
		"data":    customer,
	})
}

// DeleteCustomer deletes a customer (soft delete)
func DeleteCustomer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid customer ID",
		})
	}

	var customer models.Customer
	result := database.GetDB().First(&customer, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Customer not found",
		})
	}

	result = database.GetDB().Delete(&customer)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete customer",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Customer deleted successfully",
	})
}
