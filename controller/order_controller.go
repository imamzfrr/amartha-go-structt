package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imamzfrr/amartha-go-struct/entity"
	"github.com/imamzfrr/amartha-go-struct/services"
)

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController(service services.OrderService) *OrderController {
	return &OrderController{service}
}

func (c *OrderController) Create(ctx *fiber.Ctx) error {
	var order entity.Order
	if err := ctx.BodyParser(&order); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.orderService.Create(&order); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(order)
}

func (c *OrderController) GetAll(ctx *fiber.Ctx) error {
	orders, err := c.orderService.GetAll()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(orders)
}

func (c *OrderController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	order, err := c.orderService.GetById(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Order not found"})
	}
	return ctx.JSON(order)
}

func (c *OrderController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var order entity.Order
	if err := ctx.BodyParser(&order); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.orderService.Update(id, &order); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(order)
}

func (c *OrderController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.orderService.Delete(id); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}
