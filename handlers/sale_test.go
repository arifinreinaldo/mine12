package handlers

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/arifinreinaldo/mine12/database"
	"github.com/arifinreinaldo/mine12/models"
	"github.com/gofiber/fiber/v2"
)

func TestCreateSale(t *testing.T) {
	setupTestDB(t)

	// Create test customer
	customer := models.Customer{
		Name:     "Test Customer",
		Email:    "customer@example.com",
		Phone:    "1234567890",
		Type:     "retail",
		IsActive: true,
	}
	database.DB.Create(&customer)

	// Create test item
	item := models.Item{
		Code:     "ITEM001",
		Name:     "Test Item",
		Price:    100.0,
		Stock:    50,
		Unit:     "pcs",
		IsActive: true,
	}
	database.DB.Create(&item)

	app := fiber.New()
	app.Post("/sales", CreateSale)

	saleData := models.CreateSaleRequest{
		CustomerID: customer.ID,
		OrderDate:  time.Now(),
		Notes:      "Test order",
		Items: []models.CreateSaleItem{
			{
				ItemID:   item.ID,
				Quantity: 5,
			},
		},
	}

	body, _ := json.Marshal(saleData)
	req := httptest.NewRequest("POST", "/sales", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != fiber.StatusCreated {
		t.Errorf("Expected status 201, got %d", resp.StatusCode)
	}

	// Verify stock was reduced
	var updatedItem models.Item
	database.DB.First(&updatedItem, item.ID)
	expectedStock := 45
	if updatedItem.Stock != expectedStock {
		t.Errorf("Expected stock %d, got %d", expectedStock, updatedItem.Stock)
	}
}

func TestCreateSaleInsufficientStock(t *testing.T) {
	setupTestDB(t)

	// Create test customer
	customer := models.Customer{
		Name:     "Test Customer",
		Email:    "customer@example.com",
		Phone:    "1234567890",
		Type:     "retail",
		IsActive: true,
	}
	database.DB.Create(&customer)

	// Create test item with low stock
	item := models.Item{
		Code:     "ITEM001",
		Name:     "Test Item",
		Price:    100.0,
		Stock:    3,
		Unit:     "pcs",
		IsActive: true,
	}
	database.DB.Create(&item)

	app := fiber.New()
	app.Post("/sales", CreateSale)

	saleData := models.CreateSaleRequest{
		CustomerID: customer.ID,
		OrderDate:  time.Now(),
		Items: []models.CreateSaleItem{
			{
				ItemID:   item.ID,
				Quantity: 10, // More than available stock
			},
		},
	}

	body, _ := json.Marshal(saleData)
	req := httptest.NewRequest("POST", "/sales", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != fiber.StatusBadRequest {
		t.Errorf("Expected status 400 for insufficient stock, got %d", resp.StatusCode)
	}
}

func TestGetAllSales(t *testing.T) {
	setupTestDB(t)

	app := fiber.New()
	app.Get("/sales", GetAllSales)

	req := httptest.NewRequest("GET", "/sales", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestGetSalesStats(t *testing.T) {
	setupTestDB(t)

	app := fiber.New()
	app.Get("/sales/stats", GetSalesStats)

	req := httptest.NewRequest("GET", "/sales/stats", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}
