package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arifinreinaldo/mine12/database"
	"github.com/arifinreinaldo/mine12/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) {
	var err error
	database.DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal("Failed to connect to test database:", err)
	}

	err = database.DB.AutoMigrate(
		&models.User{},
		&models.Item{},
		&models.Customer{},
		&models.Sale{},
		&models.SaleItem{},
	)
	if err != nil {
		t.Fatal("Failed to migrate test database:", err)
	}
}

func TestCreateUser(t *testing.T) {
	setupTestDB(t)

	app := fiber.New()
	app.Post("/users", CreateUser)

	userData := models.CreateUserRequest{
		Username:    "testuser",
		Email:       "test@example.com",
		Password:    "password123",
		FullName:    "Test User",
		AccessLevel: "staff",
	}

	body, _ := json.Marshal(userData)
	req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != fiber.StatusCreated {
		t.Errorf("Expected status 201, got %d", resp.StatusCode)
	}
}

func TestGetAllUsers(t *testing.T) {
	setupTestDB(t)

	// Create test user
	user := models.User{
		Username:    "testuser",
		Email:       "test@example.com",
		Password:    "hashedpassword",
		FullName:    "Test User",
		AccessLevel: "staff",
		IsActive:    true,
	}
	database.DB.Create(&user)

	app := fiber.New()
	app.Get("/users", GetAllUsers)

	req := httptest.NewRequest("GET", "/users", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestGetUserByID(t *testing.T) {
	setupTestDB(t)

	// Create test user
	user := models.User{
		Username:    "testuser",
		Email:       "test@example.com",
		Password:    "hashedpassword",
		FullName:    "Test User",
		AccessLevel: "staff",
		IsActive:    true,
	}
	database.DB.Create(&user)

	app := fiber.New()
	app.Get("/users/:id", GetUserByID)

	req := httptest.NewRequest("GET", "/users/1", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestUpdateUser(t *testing.T) {
	setupTestDB(t)

	// Create test user
	user := models.User{
		Username:    "testuser",
		Email:       "test@example.com",
		Password:    "hashedpassword",
		FullName:    "Test User",
		AccessLevel: "staff",
		IsActive:    true,
	}
	database.DB.Create(&user)

	app := fiber.New()
	app.Put("/users/:id", UpdateUser)

	updateData := models.UpdateUserRequest{
		FullName: "Updated Name",
	}

	body, _ := json.Marshal(updateData)
	req := httptest.NewRequest("PUT", "/users/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestDeleteUser(t *testing.T) {
	setupTestDB(t)

	// Create test user
	user := models.User{
		Username:    "testuser",
		Email:       "test@example.com",
		Password:    "hashedpassword",
		FullName:    "Test User",
		AccessLevel: "staff",
		IsActive:    true,
	}
	database.DB.Create(&user)

	app := fiber.New()
	app.Delete("/users/:id", DeleteUser)

	req := httptest.NewRequest("DELETE", "/users/1", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}
