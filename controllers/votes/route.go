package votes

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
)

func NewRouteHandler(mongoClient *mongo.Client) controllers.RouteHandler {
	return func(ctx *fiber.Ctx) error {
		adventure := hoppers.AdventureFilterFromString(ctx.Query("adventure", hoppers.AnyAdventure.String()))

		votesCollection := &db.VotesCollection{
			Connection: mongoClient,
		}

		cursor, err := votesCollection.GetCollection().Find(
			context.Background(),
			getVotesFilter(VotesFilter{
				Adventure: adventure,
			}),
		)
		if err != nil {
			log.Println(err)
			return controllers.CreateServerError(ctx)
		}

		votes := []models.VoteDocument{}
		if err = cursor.All(context.Background(), &votes); err != nil {
			log.Println(err)
			return controllers.CreateServerError(ctx)
		}

		return ctx.JSON(fiber.Map{
			"filter": fiber.Map{
				"adventure": adventure.String(),
			},
			"data": formatVotes(votes),
		})
	}
}

// ----------------------------------------
// Mongo filters
// ----------------------------------------

func getVotesFilter(veShareFilter VotesFilter) bson.D {
	filter := bson.D{}

	filter = append(filter, getAdventureFilter(veShareFilter.Adventure))

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
			"share":     vote.VotesShare,
		}
	}

	return data
}
