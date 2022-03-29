package controllers

import "github.com/gofiber/fiber/v2"

func CreateServerError(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "Internal server error",
	})
}

func CreateBadRequestError(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": "Bad request",
	})
}
