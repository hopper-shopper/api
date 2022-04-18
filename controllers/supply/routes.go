package supply

import (
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/steschwa/hopper-analytics-api/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouteHandler(mongoClient *mongo.Client) controllers.RouteHandler {
	return func(ctx *fiber.Ctx) error {
		supplyType := SupplyTypeFromString(ctx.Query("type"))

		filter := SupplyFilter{
			Type: supplyType,
		}
		err := ValidateFilter(filter)
		if err != nil {
			log.Println(err)
			return controllers.CreateValidationError(ctx)
		}

		loader := NewSupplyHistoryLoader(mongoClient)

		aggregates, err := loader.Load(filter)
		if err != nil {
			sentry.CaptureException(err)
			log.Println(err)
			return controllers.CreateServerError(ctx)
		}

		return ctx.JSON(fiber.Map{
			"filter": fiber.Map{
				"type": supplyType,
			},
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
			"datetime": supply.Datetime,
			"supply":   supply.Supply,
		}
	}

	return data
}
