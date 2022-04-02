package prices

import (
	"github.com/gofiber/fiber/v2"
	"github.com/steschwa/hopper-analytics-api/controllers"
	"github.com/steschwa/hopper-analytics-api/prices"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouteHandler(mongoClient *mongo.Client) controllers.RouteHandler {
	return func(ctx *fiber.Ctx) error {
		loader := prices.NewPriceLoader(mongoClient)

		avaxUsd := loader.LoadAvaxUsdPrice()
		avaxEur := loader.LoadAvaxEurPrice()
		flyUsd := loader.LoadFlyUsdPrice()
		flyEur := loader.LoadFlyEurPrice()

		return ctx.JSON(fiber.Map{
			"data": fiber.Map{
				"AVAX": fiber.Map{
					"EUR": avaxEur,
					"USD": avaxUsd,
				},
				"FLY": fiber.Map{
					"EUR": flyEur,
					"USD": flyUsd,
				},
			},
		})
	}
}
