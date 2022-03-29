package hoppers

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/steschwa/hopper-analytics-api/controllers"
	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouteHandler(mongoClient *mongo.Client) controllers.RouteHandler {
	return func(ctx *fiber.Ctx) error {
		adventure := AdventureFilterFromString(ctx.Query("adventure", AnyAdventure.String()))
		market := MarketFilterFromString(ctx.Query("market", AnyMarket.String()))

		hoppersCollection := &db.HoppersCollection{
			Connection: mongoClient,
		}

		cursor, err := hoppersCollection.GetCollection().Find(
			context.Background(),
			getHoppersFilter(HoppersFilter{
				Adventure: adventure,
				Market:    market,
			}),
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
			"filter": fiber.Map{
				"adventure": adventure.String(),
				"market":    market.String(),
			},
			"data": formatHoppers(hoppers),
		})
	}
}

// ----------------------------------------
// Mongo filters
// ----------------------------------------

func getHoppersFilter(hoppersFilter HoppersFilter) bson.D {
	filter := bson.D{}

	filter = append(filter, getAdventureFilter(hoppersFilter.Adventure))
	filter = append(filter, getMarketFilter(hoppersFilter.Market))

	return filter
}
func getAdventureFilter(adventureFilter AdventureFilter) bson.E {
	switch adventureFilter {
	case RiverAdventure:
		return bson.E{Key: "canEnterRiver", Value: true}
	case ForestAdventure:
		return bson.E{Key: "canEnterForest", Value: true}
	case GreatLakeAdventure:
		return bson.E{Key: "canEnterGreatLake", Value: true}
	default:
		return bson.E{}
	}
}
func getMarketFilter(marketFilter MarketFilter) bson.E {
	switch marketFilter {
	case OnMarket:
		return bson.E{Key: "listingActive", Value: true}
	case OffMarket:
		return bson.E{Key: "listingActive", Value: false}
	default:
		return bson.E{}
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
