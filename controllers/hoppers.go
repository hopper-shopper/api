package controllers

import (
	"context"
	"log"

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

		hoppersCollection := &db.HoppersCollection{
			Connection: mongoClient,
		}

		cursor, err := hoppersCollection.GetCollection().Find(
			context.Background(),
			getHoppersFilter(adventure),
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
			},
			"data": formatHoppers(hoppers),
		})
	}
}

func getHoppersFilter(adventure constants.Adventure) bson.D {
	switch adventure {
	case constants.AdventurePond:
		return bson.D{{Key: "canEnterPond", Value: true}}
	case constants.AdventureStream:
		return bson.D{{Key: "canEnterStream", Value: true}}
	case constants.AdventureSwamp:
		return bson.D{{Key: "canEnterSwamp", Value: true}}
	case constants.AdventureRiver:
		return bson.D{{Key: "canEnterRiver", Value: true}}
	case constants.AdventureForest:
		return bson.D{{Key: "canEnterForest", Value: true}}
	case constants.AdventureGreatLake:
		return bson.D{{Key: "canEnterGreatLake", Value: true}}
	default:
		return bson.D{}
	}
}

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
