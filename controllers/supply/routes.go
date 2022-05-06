package supply

import (
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/steschwa/hopper-analytics-api/controllers"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
)

func NewRouteHandler(dbClient *db.MongoDbClient) controllers.RouteHandler {
	return func(ctx *fiber.Ctx) error {
		loader := NewSupplyHistoryLoader(dbClient)

		aggregates, err := loader.Load()
		if err != nil {
			sentry.CaptureException(err)
			log.Println(err)
			return controllers.CreateServerError(ctx)
		}

		return ctx.JSON(fiber.Map{
			"data": formatAggregates(aggregates),
		})
	}
}

// ----------------------------------------
// Response formatters
// ----------------------------------------

func formatAggregates(aggregates []AggregatedSupply) []fiber.Map {
	data := make([]fiber.Map, len(aggregates))
	for i, supply := range aggregates {
		data[i] = fiber.Map{
			"date":      supply.Datetime.Format("2006-01-02"),
			"supply":    supply.Supply,
			"burned":    supply.Burned,
			"staked":    supply.Staked,
			"available": supply.Available,
			"free":      supply.Free,
		}
	}

	return data
}
