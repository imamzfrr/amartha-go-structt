package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imamzfrr/amartha-go-struct/entity"
	"github.com/imamzfrr/amartha-go-struct/services"
)

type SupplierController struct {
	supplierService services.SupplierService
}

func NewSupplierController(service services.SupplierService) *SupplierController {
	return &SupplierController{service}
}

func (c *SupplierController) Create(ctx *fiber.Ctx) error {
	var supplier entity.SupplierEntity
	if err := ctx.BodyParser(&supplier); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if supplier.Name == "" || supplier.Address == "" {
		return ctx.Status(400).JSON(fiber.Map{"error": "Name and Address cannot be empty"})
	}
	if err := c.supplierService.Create(&supplier); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(supplier)
}

func (c *SupplierController) GetAll(ctx *fiber.Ctx) error {
	suppliers, err := c.supplierService.GetAll()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(suppliers)
}

func (c *SupplierController) GetByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}
	supplier, err := c.supplierService.GetById(uint64(id))
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Supplier not found"})
	}
	return ctx.JSON(supplier)
}

func (c *SupplierController) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}
	var supplier entity.SupplierEntity
	if err := ctx.BodyParser(&supplier); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if supplier.Name == "" || supplier.Address == "" {
		return ctx.Status(400).JSON(fiber.Map{"error": "Name and Address cannot be empty"})
	}
	if err := c.supplierService.Update(uint64(id), &supplier); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(supplier)
}

func (c *SupplierController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}
	if err := c.supplierService.Delete(uint64(id)); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}
