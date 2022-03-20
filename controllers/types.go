package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type (
	RouteHandler func(ctx *fiber.Ctx) error
)
