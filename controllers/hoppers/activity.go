package hoppers

import (
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/steschwa/hopper-analytics-api/controllers"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
)

func NewHoppersActivityHistoryRouteHandler(dbClient *db.MongoDbClient) controllers.RouteHandler {
	return func(ctx *fiber.Ctx) error {
		loader := NewHoppersActivityLoader(dbClient)
		aggregates, err := loader.History()
		if err != nil {
			log.Println(err)
			sentry.CaptureException(err)
			return controllers.CreateServerError(ctx)
		}

		formatted := make([]fiber.Map, len(aggregates))
		for i, entry := range aggregates {
			formatted[i] = fiber.Map{
				"date":        entry.Date.Format("2006-01-02"),
				"adventure":   entry.Adventure,
				"breeding":    entry.Breeding,
				"pond":        entry.Pond,
				"stream":      entry.Stream,
				"swamp":       entry.Swamp,
				"river":       entry.River,
				"forest":      entry.Forest,
				"greatLake":   entry.GreatLake,
				"marketplace": entry.Marketplace,
				"idle":        entry.Idle,
			}
		}

		return ctx.JSON(fiber.Map{
			"data": formatted,
		})
	}
}
