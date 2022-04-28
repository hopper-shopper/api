package user

import (
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/steschwa/hopper-analytics-api/controllers"
	"github.com/steschwa/hopper-analytics-api/utils"
	"github.com/steschwa/hopper-analytics-collector/constants"
	"github.com/steschwa/hopper-analytics-collector/contracts"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// ----------------------------------------
// Response formatters
// ----------------------------------------

func NewUserCapRouteHandler(onChainClient *contracts.OnChainClient, mongoClient *mongo.Client) controllers.RouteHandler {
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

		userCapCalculator := NewUserFlyGenerationCalculator(onChainClient, mongoClient)
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

func NewUserEarnings(onChainClient *contracts.OnChainClient, mongoClient *mongo.Client) controllers.RouteHandler {
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

		baseFlyCalculator := NewUserBaseFlyCalculator(onChainClient, mongoClient)
		baseFly, err := baseFlyCalculator.CalculateBaseFly(adventure, user)
		if err != nil {
			log.Println(err)
			sentry.CaptureException(err)
			return controllers.CreateServerError(ctx)
		}

		boostedFlyCalculator := NewUserBoostedFlyCalculator(onChainClient, mongoClient)
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

func NewUserEarningsAllTime() controllers.RouteHandler {
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

		calculator := NewUserAllTimeEarningsCalculator()
		allTimeEarnings, err := calculator.GetAllTimeEarnings(filter)
		if err != nil {
			log.Println(err)
			sentry.CaptureException(err)
			return controllers.CreateServerError(ctx)
		}
		allTimeEarningsF, _ := utils.ToDecimal(allTimeEarnings, 18).Float64()

		return ctx.JSON(fiber.Map{
			"filter": fiber.Map{
				"user":      filter.User,
				"adventure": filter.Adventure.String(),
			},
			"data": fiber.Map{
				"total": allTimeEarningsF,
			},
		})
	}
}

func NewUserEarningsHistory() controllers.RouteHandler {
	return func(ctx *fiber.Ctx) error {
		user := ctx.Query("user")
		adventure := constants.AdventureFromString(ctx.Query("adventure"))
		groupTransfersBy := GroupTransfersByFromString(ctx.Query("groupBy"))

		filter := UserEarningsHistoryFilter{
			User:      user,
			Adventure: adventure,
			GroupBy:   groupTransfersBy,
		}
		err := ValidateFilter(filter)
		if err != nil {
			log.Println(err)
			return controllers.CreateValidationError(ctx)
		}

		calculator := NewUserAllTimeEarningsCalculator()
		groups, err := calculator.GetTransfersHistory(filter)
		if err != nil {
			log.Println(err)
			sentry.CaptureException(err)
			return controllers.CreateServerError(ctx)
		}

		return ctx.JSON(fiber.Map{
			"filter": fiber.Map{
				"user":      filter.User,
				"adventure": filter.Adventure.String(),
				"groupBy":   filter.GroupBy.String(),
			},
			"data": formatUserEarningsHistory(groups),
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

func formatUserEarningsHistory(history map[string]*GroupedTransfers) fiber.Map {
	data := fiber.Map{}

	for _, entry := range history {
		formattedEarnings := make([]fiber.Map, len(entry.Transfers))
		for j, transfer := range entry.Transfers {
			amountF, _ := utils.ToDecimal(transfer.Amount, 18).Float64()

			formattedEarnings[j] = fiber.Map{
				"amount":    amountF,
				"timestamp": transfer.Timestamp,
			}
		}

		data[entry.Key] = formattedEarnings
	}

	return data
}
