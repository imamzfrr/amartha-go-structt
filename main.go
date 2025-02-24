package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/imamzfrr/amartha-go-struct/controller"
	"github.com/imamzfrr/amartha-go-struct/database"
	"github.com/imamzfrr/amartha-go-struct/repository"
	"github.com/imamzfrr/amartha-go-struct/routes"
	"github.com/imamzfrr/amartha-go-struct/services"
)

func main() {
	app := fiber.New()

	// Koneksi ke database
	database.ConnectDatabase()

	// Product
	productRepo := repository.NewProductRepository(database.DB)
	productService := services.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	// Customer
	customerRepo := repository.NewCustomerRepository(database.DB)
	customerService := services.NewCustomerService(customerRepo)
	customerController := controller.NewCustomerController(customerService)

	// Employee
	employeeRepo := repository.NewEmployeeRepository(database.DB)
	employeeService := services.NewEmployeeService(employeeRepo)
	employeeController := controller.NewEmployeeController(employeeService)

	// Category
	categoryRepo := repository.NewCategoryRepository(database.DB)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryController := controller.NewCategoryController(categoryService)

	// Order
	orderRepo := repository.NewOrderRepository(database.DB)
	orderService := services.NewOrderService(orderRepo)
	orderController := controller.NewOrderController(orderService)

	// Tax
	taxRepo := repository.NewTaxRepository(database.DB)
	taxService := services.NewTaxService(taxRepo)
	taxController := controller.NewTaxController(taxService)

	// Payment
	paymentRepo := repository.NewPaymentRepository(database.DB)
	paymentService := services.NewPaymentService(paymentRepo)
	paymentController := controller.NewPaymentController(paymentService)

	// Supplier
	supplierRepo := repository.NewSupplierRepository(database.DB)
	supplierService := services.NewSupplierService(supplierRepo)
	supplierController := controller.NewSupplierController(supplierService)

	//Receipt
	receiptRepo := repository.NewReceiptRepository(database.DB)
	receiptService := services.NewReceiptService(receiptRepo)
	receiptController := controller.NewReceiptController(receiptService)

	//Discount
	discountRepo := repository.NewDiscountRepository(database.DB)
	discountService := services.NewDiscountService(discountRepo)
	discountController := controller.NewDiscountController(discountService)

	//Inventory
	inventoryRepo := repository.NewInventoryRepository(database.DB)
	inventoryService := services.NewInventoryService(inventoryRepo)
	inventoryController := controller.NewInventoryController(inventoryService)

	//Inventory
	shipRepo := repository.NewShipRepository(database.DB)
	shipService := services.NewShipService(shipRepo)
	shipController := controller.NewShipController(shipService)

	// Setup Routes
	routes.SetupRoutes(
		app,
		productController,
		customerController,
		employeeController,
		categoryController,
		orderController,
		taxController,
		paymentController,
		supplierController,
		discountController,
		inventoryController,
		receiptController,
		shipController,
	)

	// Jalankan server
	err := app.Listen(":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
