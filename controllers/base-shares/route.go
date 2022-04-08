package baseshares

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/steschwa/hopper-analytics-api/controllers"
	"github.com/steschwa/hopper-analytics-api/controllers/hoppers"
	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewRouteHandler(mongoClient *mongo.Client) controllers.RouteHandler {
	return func(ctx *fiber.Ctx) error {
		adventure := hoppers.AdventureFilterFromString(ctx.Query("adventure", hoppers.AnyAdventure.String()))

		baseSharesCollection := &db.BaseSharesCollection{
			Connection: mongoClient,
		}

		var baseSharesLimit int64 = 1
		if adventure == hoppers.AnyAdventure {
			baseSharesLimit = 6
		}
		cursor, err := baseSharesCollection.GetCollection().Find(
			context.Background(),
			getBaseSharesFilter(adventure),
			&options.FindOptions{
				Sort: bson.D{{
					Key:   "updated",
					Value: -1,
				}},
				Limit: &baseSharesLimit,
			},
		)
		if err != nil {
			log.Println(err)
			return controllers.CreateServerError(ctx)
		}

		baseShares := []models.BaseSharesDocument{}
		if err = cursor.All(context.Background(), &baseShares); err != nil {
			log.Println(err)
			return controllers.CreateServerError(ctx)
		}

		return ctx.JSON(fiber.Map{
			"filter": fiber.Map{
				"adventure": adventure.String(),
			},
			"data": formatBaseShares(baseShares),
		})
	}
}

// ----------------------------------------
// Mongo filters
// ----------------------------------------

func getBaseSharesFilter(adventure hoppers.AdventureFilter) bson.D {
	filter := bson.D{}

	filter = append(filter, getAdventureFilter(adventure))

	return filter
}
func getAdventureFilter(adventureFilter hoppers.AdventureFilter) bson.E {
	switch adventureFilter {
	case hoppers.PondAdventure:
		return bson.E{Key: "adventure", Value: "pond"}
	case hoppers.StreamAdventure:
		return bson.E{Key: "adventure", Value: "stream"}
	case hoppers.SwampAdventure:
		return bson.E{Key: "adventure", Value: "swamp"}
	case hoppers.RiverAdventure:
		return bson.E{Key: "adventure", Value: "river"}
	case hoppers.ForestAdventure:
		return bson.E{Key: "adventure", Value: "forest"}
	case hoppers.GreatLakeAdventure:
		return bson.E{Key: "adventure", Value: "great-lake"}
	default:
		return bson.E{}
	}
}

// ----------------------------------------
// Response formatters
// ----------------------------------------

func formatBaseShares(baseShares []models.BaseSharesDocument) []fiber.Map {
	data := make([]fiber.Map, len(baseShares))
	for i, baseShare := range baseShares {
		data[i] = fiber.Map{
			"adventure":  baseShare.Adventure,
			"baseShares": baseShare.TotalBaseShares,
		}
	}

	return data
}
