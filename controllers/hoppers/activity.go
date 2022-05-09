package hoppers

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/steschwa/hopper-analytics-api/controllers"
	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewHoppersActivityRouteHandler(dbClient *db.MongoDbClient) controllers.RouteHandler {
	return func(ctx *fiber.Ctx) error {

		activityCollection := &db.HoppersActivityCollection{
			Client: dbClient,
		}
		limit := int64(1)
		cursor, err := activityCollection.GetCollection().Find(
			context.Background(),
			bson.D{{}},
			&options.FindOptions{
				Sort: bson.D{{
					Key:   "timestamp",
					Value: -1,
				}},
				Limit: &limit,
			},
		)
		if err != nil {
			log.Println(err)
			return controllers.CreateServerError(ctx)
		}

		activities := []models.HoppersActivityDocument{}
		if err = cursor.All(context.Background(), &activities); err != nil {
			log.Println(err)
			return controllers.CreateServerError(ctx)
		}

		if len(activities) != 1 {
			log.Println(err)
			return controllers.CreateServerError(ctx)
		}

		activity := activities[0]

		return ctx.JSON(fiber.Map{
			"data": fiber.Map{
				"adventure":   activity.Adventure,
				"breeding":    activity.Breeding,
				"pond":        activity.Pond,
				"stream":      activity.Stream,
				"swamp":       activity.Swamp,
				"river":       activity.River,
				"forest":      activity.Forest,
				"greatLake":   activity.GreatLake,
				"marketplace": activity.Marketplace,
				"idle":        activity.Idle,
			},
		})
	}
}
