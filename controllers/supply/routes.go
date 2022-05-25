package supply

import (
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/steschwa/hopper-analytics-api/controllers"
	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
)

func NewRouteHandler(dbClient *db.MongoDbClient) controllers.RouteHandler {
	return func(ctx *fiber.Ctx) error {
		loader := NewSupplyHistoryLoader(dbClient)

		supplies, err := loader.Load()
		if err != nil {
			sentry.CaptureException(err)
			log.Println(err)
			return controllers.CreateServerError(ctx)
		}

		return ctx.JSON(fiber.Map{
			"data": formatSupplies(supplies),
		})
	}
}

// ----------------------------------------
// Response formatters
// ----------------------------------------

func formatSupplies(supplies []models.FlySupplyDocument) []fiber.Map {
	data := make([]fiber.Map, len(supplies))
	for i, supply := range supplies {
		data[i] = fiber.Map{
			"date":        supply.Timestamp.Format("2006-01-02"),
			"minted":      supply.Minted,
			"burned":      supply.Burned,
			"circulating": supply.Circulating,
			"staked":      supply.Staked,
			"free":        supply.Free,
		}
	}

	return data
}
