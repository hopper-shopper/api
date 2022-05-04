package votes

import (
	"context"
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/steschwa/hopper-analytics-api/controllers"
	"github.com/steschwa/hopper-analytics-api/controllers/hoppers"
	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO Use new HopperDocument schema `activity`

func NewRouteHandler(dbClient *db.MongoDbClient) controllers.RouteHandler {
	return func(ctx *fiber.Ctx) error {
		adventure := hoppers.AdventureFilterFromString(ctx.Query("adventure", string(hoppers.AnyAdventure)))

		votesCollection := &db.VotesCollection{
			Client: dbClient,
		}

		var votesLimit int64 = 1
		if adventure == hoppers.AnyAdventure {
			votesLimit = 6
		}
		cursor, err := votesCollection.GetCollection().Find(
			context.Background(),
			getVotesFilter(adventure),
			&options.FindOptions{
				Sort: bson.D{{
					Key:   "updated",
					Value: -1,
				}},
				Limit: &votesLimit,
			},
		)
		if err != nil {
			log.Println(err)
			sentry.CaptureException(err)
			return controllers.CreateServerError(ctx)
		}

		votes := []models.VoteDocument{}
		if err = cursor.All(context.Background(), &votes); err != nil {
			log.Println(err)
			return controllers.CreateServerError(ctx)
		}

		return ctx.JSON(fiber.Map{
			"filter": fiber.Map{
				"adventure": adventure,
			},
			"data": formatVotes(votes),
		})
	}
}

// ----------------------------------------
// Mongo filters
// ----------------------------------------

func getVotesFilter(adventure hoppers.AdventureFilter) bson.D {
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

func formatVotes(votes []models.VoteDocument) []fiber.Map {
	data := make([]fiber.Map, len(votes))
	for i, vote := range votes {
		data[i] = fiber.Map{
			"adventure": vote.Adventure,
			"votes":     vote.Votes,
		}
	}

	return data
}
