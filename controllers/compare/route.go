package compare

import (
	"context"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/steschwa/hopper-analytics-api/controllers"
	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouteHandler(mongoClient *mongo.Client) controllers.RouteHandler {
	return func(ctx *fiber.Ctx) error {
		tokenIds := parseTokenIds(ctx.Query("hoppers"))

		if len(tokenIds) == 0 {
			return ctx.JSON(fiber.Map{
				"hoppers": []string{},
				"data":    fiber.Map{},
			})
		}

		hoppersCollection := &db.HoppersCollection{
			Connection: mongoClient,
		}

		cursor, err := hoppersCollection.GetCollection().Find(
			context.Background(),
			getHoppersFilter(tokenIds),
		)
		if err != nil {
			log.Println(err)
			return controllers.CreateServerError(ctx)
		}

		hoppers := []models.HopperDocument{}
		if err = cursor.All(context.Background(), &hoppers); err != nil {
			log.Println(err)
			return controllers.CreateServerError(ctx)
		}

		return ctx.JSON(fiber.Map{
			"hoppers": tokenIds,
			"data":    formatHoppers(hoppers),
		})
	}
}

// ----------------------------------------
// Input parsing
// ----------------------------------------

func parseTokenIds(value string) []string {
	tokenIds := strings.Split(value, ",")

	for i, tokenId := range tokenIds {
		tokenIds[i] = strings.Replace(tokenId, " ", "", -1)
	}

	return tokenIds
}

// ----------------------------------------
// Mongo filters
// ----------------------------------------

func getHoppersFilter(tokenIds []string) bson.M {
	return bson.M{
		"tokenId": bson.M{
			"$in": tokenIds,
		},
	}
}

// ----------------------------------------
// Response formatters
// ----------------------------------------

func formatHoppers(hoppers []models.HopperDocument) []fiber.Map {
	data := make([]fiber.Map, len(hoppers))
	for i, hopper := range hoppers {
		data[i] = fiber.Map{
			"tokenId":      hopper.TokenId,
			"strength":     hopper.Strength,
			"agility":      hopper.Agility,
			"vitality":     hopper.Vitality,
			"intelligence": hopper.Intelligence,
			"fertility":    hopper.Fertility,
			"level":        hopper.Level,
			"image":        hopper.Image,
			"adventure":    hopper.Adventure,
			"rating": fiber.Map{
				"pond":      hopper.RatingPond,
				"stream":    hopper.RatingStream,
				"swamp":     hopper.RatingSwamp,
				"river":     hopper.RatingRiver,
				"forest":    hopper.RatingForest,
				"greatLake": hopper.RatingGreatLake,
			},
			"baseFly": fiber.Map{
				"pond":      hopper.BaseFlyPond,
				"stream":    hopper.BaseFlyStream,
				"swamp":     hopper.BaseFlySwamp,
				"river":     hopper.BaseFlyRiver,
				"forest":    hopper.BaseFlyForest,
				"greatLake": hopper.BaseFlyGreatLake,
			},
			"listing": fiber.Map{
				"active": hopper.ListingActive,
				"price":  hopper.ListingPrice,
			},
		}
	}

	return data
}
