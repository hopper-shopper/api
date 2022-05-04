package prices

import (
	"github.com/gofiber/fiber/v2"
	"github.com/steschwa/hopper-analytics-api/controllers"
	"github.com/steschwa/hopper-analytics-api/prices"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
)

func NewLatestPriceRouteHandler(dbClient *db.MongoDbClient) controllers.RouteHandler {
	return func(ctx *fiber.Ctx) error {
		loader := prices.NewPriceLoader(dbClient)

		avaxUsd := loader.LoadLatestAvaxUsdPrice()
		avaxEur := loader.LoadLatestAvaxEurPrice()
		flyUsd := loader.LoadLatestFlyUsdPrice()
		flyEur := loader.LoadLatestFlyEurPrice()
		var avaxFly = 0.0
		if flyUsd != 0 {
			avaxFly = avaxUsd / flyUsd
		}
		var flyAvax = 0.0
		if avaxUsd != 0 {
			flyAvax = flyUsd / avaxUsd
		}

		return ctx.JSON(fiber.Map{
			"data": fiber.Map{
				"AVAX": fiber.Map{
					"EUR": avaxEur,
					"USD": avaxUsd,
					"FLY": avaxFly,
				},
				"FLY": fiber.Map{
					"EUR":  flyEur,
					"USD":  flyUsd,
					"AVAX": flyAvax,
				},
			},
		})
	}
}
