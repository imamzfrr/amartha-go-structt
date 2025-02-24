package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imamzfrr/amartha-go-struct/entity"
	"github.com/imamzfrr/amartha-go-struct/services"
)

type InventoryController struct {
	inventoryService services.InventoryService
}

func NewInventoryController(service services.InventoryService) *InventoryController {
	return &InventoryController{service}
}

func (c *InventoryController) Create(ctx *fiber.Ctx) error {
	var inventory entity.Inventory
	if err := ctx.BodyParser(&inventory); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.inventoryService.Create(&inventory); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(inventory)
}

func (c *InventoryController) GetAll(ctx *fiber.Ctx) error {
	inventories, err := c.inventoryService.GetAll()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(inventories)
}

func (c *InventoryController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	inventory, err := c.inventoryService.GetById(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Inventory not found"})
	}
	return ctx.JSON(inventory)
}

func (c *InventoryController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var inventory entity.Inventory
	if err := ctx.BodyParser(&inventory); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.inventoryService.Update(id, &inventory); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(inventory)
}

func (c *InventoryController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.inventoryService.Delete(id); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}
