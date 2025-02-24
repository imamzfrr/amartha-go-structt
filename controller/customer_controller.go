package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imamzfrr/amartha-go-struct/entity"
	"github.com/imamzfrr/amartha-go-struct/services"
)

type CustomerController struct {
	customerService services.CustomerService
}

func NewCustomerController(service services.CustomerService) *CustomerController {
	return &CustomerController{service}
}

func (c *CustomerController) CreateCustomer(ctx *fiber.Ctx) error {
	var customer entity.Customer
	if err := ctx.BodyParser(&customer); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.customerService.Create(&customer); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(customer)
}

func (c *CustomerController) GetAllCustomers(ctx *fiber.Ctx) error {
	customers, err := c.customerService.GetAll()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(customers)
}

func (c *CustomerController) GetCustomerByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	customer, err := c.customerService.GetCustomerById(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Customer not found"})
	}
	return ctx.JSON(customer)
}

func (c *CustomerController) UpdateCustomer(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var customer entity.Customer
	if err := ctx.BodyParser(&customer); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.customerService.Update(id, &customer); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(customer)
}

func (c *CustomerController) DeleteCustomer(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.customerService.Delete(id); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}
