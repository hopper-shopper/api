package controllers

import (
	"context"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/steschwa/hopper-analytics-collector/constants"
	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewHoppersController(mongoClient *mongo.Client) RouteHandler {
	return func(ctx *fiber.Ctx) error {
		adventure := constants.AdventureFromString(ctx.Query("adventure", constants.AdventurePond.String()))
		market := marketFilterFromString(ctx.Query("market", AnyMarket.String()))

		hoppersCollection := &db.HoppersCollection{
			Connection: mongoClient,
		}

		cursor, err := hoppersCollection.GetCollection().Find(
			context.Background(),
			getHoppersFilter(adventure, market),
		)
		if err != nil {
			log.Println(err)
			return CreateServerError(ctx)
		}

		hoppers := []models.HopperDocument{}
		if err = cursor.All(context.Background(), &hoppers); err != nil {
			log.Println(err)
			return CreateServerError(ctx)
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

func getHoppersFilter(adventure constants.Adventure, market MarketFilter) bson.D {
	filter := bson.D{}

	filter = append(filter, getAdventureFilter(adventure))
	filter = append(filter, getMarketFilter(market))

	return filter
}
func getAdventureFilter(adventure constants.Adventure) bson.E {
	switch adventure {
	case constants.AdventurePond:
		return bson.E{Key: "canEnterPond", Value: true}
	case constants.AdventureStream:
		return bson.E{Key: "canEnterStream", Value: true}
	case constants.AdventureSwamp:
		return bson.E{Key: "canEnterSwamp", Value: true}
	case constants.AdventureRiver:
		return bson.E{Key: "canEnterRiver", Value: true}
	case constants.AdventureForest:
		return bson.E{Key: "canEnterForest", Value: true}
	case constants.AdventureGreatLake:
		return bson.E{Key: "canEnterGreatLake", Value: true}
	default:
		return bson.E{}
	}
}
func getMarketFilter(marketFilter MarketFilter) bson.E {
	switch marketFilter {
	case OnMarket:
		return bson.E{Key: "market", Value: true}
	case OffMarket:
		return bson.E{Key: "market", Value: false}
	default:
		return bson.E{}
	}
}

// ----------------------------------------
// Formatters
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
			"market":       hopper.Market,
			"ratings": fiber.Map{
				"pond":      hopper.RatingPond,
				"stream":    hopper.RatingStream,
				"swamp":     hopper.RatingSwamp,
				"river":     hopper.RatingRiver,
				"forest":    hopper.RatingForest,
				"greatLake": hopper.RatingGreatLake,
			},
		}
	}

	return data
}

// ----------------------------------------
// Filters
// ----------------------------------------

type (
	MarketFilter int
)

const (
	OnMarket MarketFilter = iota
	OffMarket
	AnyMarket
)

func marketFilterFromString(market string) MarketFilter {
	lowerCased := strings.ToLower(market)

	switch lowerCased {
	case "1", "true", "on", "yes":
		return OnMarket
	case "0", "false", "off", "no":
		return OffMarket
	default:
		return AnyMarket
	}
}

func (marketFilter MarketFilter) String() string {
	switch marketFilter {
	case OnMarket:
		return "yes"
	case OffMarket:
		return "no"
	default:
		return "any"
	}
}
