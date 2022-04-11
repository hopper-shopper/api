package user

import (
	"log"
	"math/big"

	"github.com/gofiber/fiber/v2"
	"github.com/steschwa/hopper-analytics-api/controllers"
	"github.com/steschwa/hopper-analytics-api/utils"
	"github.com/steschwa/hopper-analytics-collector/constants"
	"github.com/steschwa/hopper-analytics-collector/contracts"
)

func NewRouteHandler(onChainClient *contracts.OnChainClient) controllers.RouteHandler {
	return func(ctx *fiber.Ctx) error {
		user := ctx.Query("user")
		adventure := constants.AdventureFromString(ctx.Query("adventure"))

		filter := UserFilter{
			User:      user,
			Adventure: adventure,
		}
		err := ValidateFilter(filter)
		if err != nil {
			log.Println(err)
			return controllers.CreateValidationError(ctx)
		}

		userCap, err := onChainClient.GetUserMaxFlyGeneration(adventure, user)
		if err != nil {
			log.Println(err)
			return controllers.CreateServerError(ctx)
		}

		userBaseGenerated, _, err := onChainClient.GetUserGeneratedFly(adventure, user)
		if err != nil {
			log.Println(err)
			return controllers.CreateServerError(ctx)
		}
		userBonusGenerated, _, err := onChainClient.GetUserBonusGeneratedFly(adventure, user)
		if err != nil {
			log.Println(err)
			return controllers.CreateServerError(ctx)
		}

		sumGenerated := big.NewInt(0).Add(userBaseGenerated, userBonusGenerated)

		return ctx.JSON(fiber.Map{
			"filter": fiber.Map{
				"user":      user,
				"adventure": adventure.String(),
			},
			"data": formatUserAdventureFlyGeneration(UserAdventureFlyGeneration{
				Cap:     userCap,
				Current: sumGenerated,
			}),
		})
	}
}

// ----------------------------------------
// Response formatters
// ----------------------------------------

type UserAdventureFlyGeneration struct {
	Cap     *big.Int
	Current *big.Int
}

func formatUserAdventureFlyGeneration(flyGeneration UserAdventureFlyGeneration) fiber.Map {
	formattedCap, _ := utils.ToDecimal(flyGeneration.Cap.String(), 30).Float64()
	formattedGenerated, _ := utils.ToDecimal(flyGeneration.Current.String(), 18).Float64()

	return fiber.Map{
		"cap":     formattedCap,
		"current": formattedGenerated,
	}
}
