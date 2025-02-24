package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imamzfrr/amartha-go-struct/entity"
	"github.com/imamzfrr/amartha-go-struct/services"
)

type DiscountController struct {
	discountService services.DiscountService
}

func NewDiscountController(service services.DiscountService) *DiscountController {
	return &DiscountController{service}
}

func (c *DiscountController) Create(ctx *fiber.Ctx) error {
	var discount entity.Discount
	if err := ctx.BodyParser(&discount); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.discountService.Create(&discount); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(discount)
}

func (c *DiscountController) GetAll(ctx *fiber.Ctx) error {
	discounts, err := c.discountService.GetAll()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(discounts)
}

func (c *DiscountController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	discount, err := c.discountService.GetById(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Discount not found"})
	}
	return ctx.JSON(discount)
}

func (c *DiscountController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var discount entity.Discount
	if err := ctx.BodyParser(&discount); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.discountService.Update(id, &discount); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(discount)
}

func (c *DiscountController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.discountService.Delete(id); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}
