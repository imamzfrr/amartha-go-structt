package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imamzfrr/amartha-go-struct/entity"
	"github.com/imamzfrr/amartha-go-struct/services"
)

type CategoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController(service service.CategoryService) *CategoryController {
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
	id := ctx.Params("id")
	category, err := c.categoryService.GetCategoryByID(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Category not found"})
	}
	return ctx.JSON(category)
}

func (c *CategoryController) UpdateCategory(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
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
	id := ctx.Params("id")
	if err := c.categoryService.DeleteCategory(id); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}
