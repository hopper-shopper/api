package transfers

import (
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/steschwa/hopper-analytics-api/controllers"
	"github.com/steschwa/hopper-analytics-api/graph"
	"github.com/steschwa/hopper-analytics-api/models"
	"github.com/steschwa/hopper-analytics-api/utils"
)

func NewRouteHandler() controllers.RouteHandler {
	return func(ctx *fiber.Ctx) error {
		transferDirection := TransferDirectionFromString(ctx.Query("direction", FromUser.String()))
		user := ctx.Query("user")

		transfersFilter := TransfersFilter{
			Direction: transferDirection,
			User:      user,
		}
		err := ValidateFilter(transfersFilter)
		if err != nil {
			log.Println(err)
			return controllers.CreateValidationError(ctx)
		}

		queryClient := graph.NewTransfersGraphClient()
		transfers := []models.Transfer{}
		if transfersFilter.Direction == FromUser {
			transfers, err = queryClient.FetchTransfersFromUser(transfersFilter.User)
		} else if transfersFilter.Direction == ToUser {
			transfers, err = queryClient.FetchTransfersToUser(transfersFilter.User)
		}

		if err != nil {
			log.Println(err)
			sentry.CaptureException(err)
			return controllers.CreateServerError(ctx)
		}

		return ctx.JSON(fiber.Map{
			"fitler": fiber.Map{
				"user":      user,
				"direction": transferDirection.String(),
			},
			"data": formatTransfers(transfers),
		})
	}
}

// ----------------------------------------
// Response formatters
// ----------------------------------------

func formatTransfers(transfers []models.Transfer) []fiber.Map {
	data := []fiber.Map{}

	for _, transfer := range transfers {
		amount, _ := utils.ToDecimal(transfer.Amount.String(), 18).Float64()

		method := GetContractMethod(transfer.MethodId)
		if method == MethodUnknown {
			continue
		}

		data = append(data, fiber.Map{
			"timestamp": transfer.Timestamp,
			"type":      method.String(),
			"amount":    amount,
		})
	}

	return data
}
