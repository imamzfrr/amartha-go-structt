package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imamzfrr/amartha-go-struct/entity"
	"github.com/imamzfrr/amartha-go-struct/services"
)

type ReceiptController struct {
	receiptService services.ReceiptService
}

func NewReceiptController(service services.ReceiptService) *ReceiptController {
	return &ReceiptController{service}
}

func (c *ReceiptController) Create(ctx *fiber.Ctx) error {
	var receipt entity.Receipt
	if err := ctx.BodyParser(&receipt); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.receiptService.Create(&receipt); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(receipt)
}

func (c *ReceiptController) GetAll(ctx *fiber.Ctx) error {
	receipts, err := c.receiptService.GetAll()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(receipts)
}

func (c *ReceiptController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	receipt, err := c.receiptService.GetById(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Receipt not found"})
	}
	return ctx.JSON(receipt)
}

func (c *ReceiptController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var receipt entity.Receipt
	if err := ctx.BodyParser(&receipt); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.receiptService.Update(id, &receipt); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(receipt)
}

func (c *ReceiptController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.receiptService.Delete(id); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}
