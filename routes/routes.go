package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imamzfrr/amartha-go-struct/controller"
)

func SetupRoutes(app *fiber.App,
	productController *controller.ProductController,
	customerController *controller.CustomerController,
	employeeController *controller.EmployeeController,
	categoryController *controller.CategoryController,
	orderController *controller.OrderController,
	taxController *controller.TaxController,
	paymentController *controller.PaymentController,
	supplierController *controller.SupplierController,
	discountController *controller.DiscountController,
	inventoryController *controller.InventoryController,
	receiptController *controller.ReceiptController,
	shipController *controller.ShipController) {

	// Product Routes
	productGroup := app.Group("/products")
	productGroup.Post("/", productController.CreateProduct)
	productGroup.Get("/", productController.GetAllProducts)
	productGroup.Get("/:id", productController.GetProductByID)
	productGroup.Put("/:id", productController.UpdateProduct)
	productGroup.Delete("/:id", productController.DeleteProduct)

	// Customer Routes
	customerGroup := app.Group("/customers")
	customerGroup.Post("/", customerController.CreateCustomer)
	customerGroup.Get("/", customerController.GetAllCustomers)
	customerGroup.Get("/:id", customerController.GetCustomerByID)
	customerGroup.Put("/:id", customerController.UpdateCustomer)
	customerGroup.Delete("/:id", customerController.DeleteCustomer)

	// Employee Routes
	employeeGroup := app.Group("/employees")
	employeeGroup.Post("/", employeeController.Create)
	employeeGroup.Get("/", employeeController.GetAll)
	employeeGroup.Get("/:id", employeeController.GetByID)
	employeeGroup.Put("/:id", employeeController.UpdateEmployee)
	employeeGroup.Delete("/:id", employeeController.Delete)

	// Category Routes
	categoryGroup := app.Group("/categories")
	categoryGroup.Post("/", categoryController.CreateCategory)
	categoryGroup.Get("/", categoryController.GetAllCategories)
	categoryGroup.Get("/:id", categoryController.GetCategoryByID)
	categoryGroup.Put("/:id", categoryController.UpdateCategory)
	categoryGroup.Delete("/:id", categoryController.DeleteCategory)

	// Order Routes
	orderGroup := app.Group("/orders")
	orderGroup.Post("/", orderController.Create)
	orderGroup.Get("/", orderController.GetAll)
	orderGroup.Get("/:id", orderController.GetByID)
	orderGroup.Put("/:id", orderController.Update)
	orderGroup.Delete("/:id", orderController.Delete)

	// Tax Routes
	taxGroup := app.Group("/taxes")
	taxGroup.Post("/", taxController.Create)
	taxGroup.Get("/", taxController.GetAll)
	taxGroup.Get("/:id", taxController.GetByID)
	taxGroup.Put("/:id", taxController.Update)
	taxGroup.Delete("/:id", taxController.Delete)

	// Payment Routes
	paymentGroup := app.Group("/payments")
	paymentGroup.Post("/", paymentController.Create)
	paymentGroup.Get("/", paymentController.GetAll)
	paymentGroup.Get("/:id", paymentController.GetByID)
	paymentGroup.Put("/:id", paymentController.UpdatePayment)
	paymentGroup.Delete("/:id", paymentController.Delete)

	// Supplier Routes
	supplierGroup := app.Group("/suppliers")
	supplierGroup.Post("/", supplierController.Create)
	supplierGroup.Get("/", supplierController.GetAll)
	supplierGroup.Get("/:id", supplierController.GetByID)
	supplierGroup.Put("/:id", supplierController.Update)
	supplierGroup.Delete("/:id", supplierController.Delete)

	// Discount Routes
	discountGroup := app.Group("/discounts")
	discountGroup.Post("/", discountController.Create)
	discountGroup.Get("/", discountController.GetAll)
	discountGroup.Get("/:id", discountController.GetByID)
	discountGroup.Put("/:id", discountController.Update)
	discountGroup.Delete("/:id", discountController.Delete)

	// Inventory Routes
	inventoryGroup := app.Group("/inventory")
	inventoryGroup.Post("/", inventoryController.Create)
	inventoryGroup.Get("/", inventoryController.GetAll)
	inventoryGroup.Get("/:id", inventoryController.GetByID)
	inventoryGroup.Put("/:id", inventoryController.Update)
	inventoryGroup.Delete("/:id", inventoryController.Delete)

	// Receipt Routes
	receiptGroup := app.Group("/receipts")
	receiptGroup.Post("/", receiptController.Create)
	receiptGroup.Get("/", receiptController.GetAll)
	receiptGroup.Get("/:id", receiptController.GetByID)
	receiptGroup.Put("/:id", receiptController.Update)
	receiptGroup.Delete("/:id", receiptController.Delete)

	// Ship Routes
	shipGroup := app.Group("/ships")
	shipGroup.Post("/", shipController.Create)
	shipGroup.Get("/", shipController.GetAll)
	shipGroup.Get("/:id", shipController.GetByID)
	shipGroup.Put("/:id", shipController.Update)
	shipGroup.Delete("/:id", shipController.Delete)
}
