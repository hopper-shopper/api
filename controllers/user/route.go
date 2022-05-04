package user

import (
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/steschwa/hopper-analytics-api/controllers"
	"github.com/steschwa/hopper-analytics-collector/constants"
	"github.com/steschwa/hopper-analytics-collector/contracts"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

// ----------------------------------------
// Response formatters
// ----------------------------------------

func NewUserCapRouteHandler(onChainClient *contracts.OnChainClient, dbClient *db.MongoDbClient) controllers.RouteHandler {
	return func(ctx *fiber.Ctx) error {
		user := ctx.Query("user")
		adventure := constants.AdventureFromString(ctx.Query("adventure"))

		filter := UserCapFilter{
			User:      user,
			Adventure: adventure,
		}
		err := ValidateFilter(filter)
		if err != nil {
			log.Println(err)
			return controllers.CreateValidationError(ctx)
		}

		userCapCalculator := NewUserFlyGenerationCalculator(onChainClient, dbClient)
		flyGeneration, err := userCapCalculator.CalculateFlyGeneration(adventure, user)
		if err != nil {
			log.Println(err)
			sentry.CaptureException(err)
			return controllers.CreateServerError(ctx)
		}

		return ctx.JSON(fiber.Map{
			"filter": fiber.Map{
				"user":      user,
				"adventure": adventure.String(),
			},
			"data": formatUserAdventureFlyGeneration(flyGeneration),
		})
	}
}

func NewUserEarnings(onChainClient *contracts.OnChainClient, dbClient *db.MongoDbClient) controllers.RouteHandler {
	return func(ctx *fiber.Ctx) error {
		user := ctx.Query("user")
		adventure := constants.AdventureFromString(ctx.Query("adventure"))

		filter := UserEarningsFilter{
			User:      user,
			Adventure: adventure,
		}
		err := ValidateFilter(filter)
		if err != nil {
			log.Println(err)
			return controllers.CreateValidationError(ctx)
		}

		baseFlyCalculator := NewUserBaseFlyCalculator(onChainClient, dbClient)
		baseFly, err := baseFlyCalculator.CalculateBaseFly(adventure, user)
		if err != nil {
			log.Println(err)
			sentry.CaptureException(err)
			return controllers.CreateServerError(ctx)
		}

		boostedFlyCalculator := NewUserBoostedFlyCalculator(onChainClient, dbClient)
		boostedFly, err := boostedFlyCalculator.CalculateBoostedFly(adventure, user)
		if err != nil {
			log.Println(err)
			sentry.CaptureException(err)
			return controllers.CreateServerError(ctx)
		}

		return ctx.JSON(fiber.Map{
			"filter": fiber.Map{
				"user":      user,
				"adventure": adventure.String(),
			},
			"data": fiber.Map{
				"base":  baseFly,
				"boost": boostedFly,
			},
		})
	}
}

// ----------------------------------------
// Mongo filters
// ----------------------------------------

func getBaseSharesFilter(adventure constants.Adventure) bson.D {
	filter := bson.D{}

	filter = append(filter, getAdventureFilter(adventure))

	return filter
}
func getVotesFilter(adventure constants.Adventure) bson.D {
	filter := bson.D{}

	filter = append(filter, getAdventureFilter(adventure))

	return filter
}
func getAdventureFilter(adventure constants.Adventure) bson.E {
	switch adventure {
	case constants.AdventurePond:
		return bson.E{Key: "adventure", Value: "pond"}
	case constants.AdventureStream:
		return bson.E{Key: "adventure", Value: "stream"}
	case constants.AdventureSwamp:
		return bson.E{Key: "adventure", Value: "swamp"}
	case constants.AdventureRiver:
		return bson.E{Key: "adventure", Value: "river"}
	case constants.AdventureForest:
		return bson.E{Key: "adventure", Value: "forest"}
	case constants.AdventureGreatLake:
		return bson.E{Key: "adventure", Value: "great-lake"}
	default:
		return bson.E{}
	}
}

// ----------------------------------------
// Response formatters
// ----------------------------------------

func formatUserAdventureFlyGeneration(flyGeneration UserAdventureFlyGeneration) fiber.Map {
	return fiber.Map{
		"cap":     flyGeneration.Cap,
		"current": flyGeneration.Current,
		"time":    flyGeneration.Time,
	}
}
