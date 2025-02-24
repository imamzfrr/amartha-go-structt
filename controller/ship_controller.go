package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imamzfrr/amartha-go-struct/entity"
	"github.com/imamzfrr/amartha-go-struct/services"
)

type ShipController struct {
	shipService services.ShipService
}

func NewShipController(service services.ShipService) *ShipController {
	return &ShipController{service}
}

func (c *ShipController) Create(ctx *fiber.Ctx) error {
	var ship entity.ShipEntity
	if err := ctx.BodyParser(&ship); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.shipService.Create(&ship); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(ship)
}

func (c *ShipController) GetAll(ctx *fiber.Ctx) error {
	ships, err := c.shipService.GetAll()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(ships)
}

func (c *ShipController) GetByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	ship, err := c.shipService.GetById(uint64(id))
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Ship not found"})
	}
	return ctx.JSON(ship)
}

func (c *ShipController) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	var ship entity.ShipEntity
	if err := ctx.BodyParser(&ship); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.shipService.Update(uint64(id), &ship); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(ship)
}

func (c *ShipController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	if err := c.shipService.Delete(uint64(id)); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}
