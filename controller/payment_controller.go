package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imamzfrr/amartha-go-struct/entity"
)

type PaymentController struct {
	paymentService services.PaymentService
}

func NewPaymentController(service services.PaymentService) *PaymentController {
	return &PaymentController{service}
}

func (c *PaymentController) Create(ctx *fiber.Ctx) error {
	var payment entity.Payment
	if err := ctx.BodyParser(&payment); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.paymentService.Create(&payment); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(payment)
}

func (c *PaymentController) GetAll(ctx *fiber.Ctx) error {
	payments, err := c.paymentService.GetAll()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(payments)
}

func (c *PaymentController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	payment, err := c.paymentService.GetById(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Payment not found"})
	}
	return ctx.JSON(payment)
}

func (c *PaymentController) UpdatePayment(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var payment entity.Payment
	if err := ctx.BodyParser(&payment); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.paymentService.Update(id, &payment); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(payment)
}

func (c *PaymentController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.paymentService.Delete(id); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}
