package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imamzfrr/amartha-go-struct/entity"
	"github.com/imamzfrr/amartha-go-struct/services"
)

type EmployeeController struct {
	employeeService services.EmployeeService
}

func NewEmployeeController(service services.EmployeeService) *EmployeeController {
	return &EmployeeController{service}
}

func (c *EmployeeController) Create(ctx *fiber.Ctx) error {
	var Employee entity.Employee
	if err := ctx.BodyParser(&Employee); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.employeeService.Create(&Employee); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(Employee)
}

func (c *EmployeeController) GetAll(ctx *fiber.Ctx) error {
	Employees, err := c.employeeService.GetAll()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(Employees)
}

func (c *EmployeeController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	Employee, err := c.employeeService.GetById(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Employee not found"})
	}
	return ctx.JSON(Employee)
}

func (c *EmployeeController) UpdateEmployee(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var Employee entity.Employee
	if err := ctx.BodyParser(&Employee); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.employeeService.Update(id, &Employee); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(Employee)
}

func (c *EmployeeController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.employeeService.Delete(id); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}
