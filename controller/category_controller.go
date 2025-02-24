package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imamzfrr/amartha-go-struct/entity"
	"github.com/imamzfrr/amartha-go-struct/services"
	"strconv"
)

type CategoryController struct {
	categoryService services.CategoryService
}

func NewCategoryController(service services.CategoryService) *CategoryController {
	return &CategoryController{service}
}

func (c *CategoryController) CreateCategory(ctx *fiber.Ctx) error {
	var category entity.CategoryEntity
	if err := ctx.BodyParser(&category); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.categoryService.CreateCategory(&category); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(category)
}

func (c *CategoryController) GetAllCategories(ctx *fiber.Ctx) error {
	categories, err := c.categoryService.GetAllCategories()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(categories)
}

func (c *CategoryController) GetCategoryByID(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	category, err := c.categoryService.GetCategoryByID(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Category not found"})
	}
	return ctx.JSON(category)
}

func (c *CategoryController) UpdateCategory(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	var category entity.CategoryEntity
	if err := ctx.BodyParser(&category); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.categoryService.UpdateCategory(id, &category); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(category)
}

func (c *CategoryController) DeleteCategory(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	if err := c.categoryService.DeleteCategory(id); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}
