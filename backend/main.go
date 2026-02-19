package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	godotenv.Load()

	app := fiber.New()

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"message": "Backend is running",
		})
	})

	// Auth routes (untuk nanti)
	auth := app.Group("/auth")
	auth.Post("/register", RegisterHandler)
	auth.Post("/login", LoginHandler)
	auth.Post("/logout", LogoutHandler)

	// Product routes (untuk nanti)
	products := app.Group("/products")
	products.Get("/", GetProducts)
	products.Get("/:id", GetProduct)
	products.Post("/", CreateProduct) // Admin only
	products.Put("/:id", UpdateProduct) // Admin only
	products.Delete("/:id", DeleteProduct) // Admin only

	// Order routes (untuk nanti)
	orders := app.Group("/orders")
	orders.Get("/", GetOrders)
	orders.Get("/:id", GetOrder)
	orders.Post("/", CreateOrder)
	orders.Put("/:id", UpdateOrder)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s\n", port)
	log.Fatal(app.Listen(":" + port))
}

// Auth handlers (placeholder)
func RegisterHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Register endpoint"})
}

func LoginHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Login endpoint"})
}

func LogoutHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Logout endpoint"})
}

// Product handlers (placeholder)
func GetProducts(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get products"})
}

func GetProduct(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get product"})
}

func CreateProduct(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Create product"})
}

func UpdateProduct(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Update product"})
}

func DeleteProduct(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Delete product"})
}

// Order handlers (placeholder)
func GetOrders(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get orders"})
}

func GetOrder(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get order"})
}

func CreateOrder(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Create order"})
}

func UpdateOrder(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Update order"})
}
