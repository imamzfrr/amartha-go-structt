package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imamzfrr/amartha-go-struct/entity"
	"github.com/imamzfrr/amartha-go-struct/services"
)

type TaxController struct {
	taxService services.TaxService
}

func NewTaxController(service services.TaxService) *TaxController {
	return &TaxController{service}
}

func (c *TaxController) Create(ctx *fiber.Ctx) error {
	var tax entity.Tax
	if err := ctx.BodyParser(&tax); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.taxService.Create(&tax); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(tax)
}

func (c *TaxController) GetAll(ctx *fiber.Ctx) error {
	taxes, err := c.taxService.GetAll()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(taxes)
}

func (c *TaxController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	tax, err := c.taxService.GetById(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Tax not found"})
	}
	return ctx.JSON(tax)
}

func (c *TaxController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var tax entity.Tax
	if err := ctx.BodyParser(&tax); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.taxService.Update(id, &tax); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(tax)
}

func (c *TaxController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.taxService.Delete(id); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}
