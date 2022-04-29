package transfers

import (
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/steschwa/hopper-analytics-api/controllers"
	"github.com/steschwa/hopper-analytics-api/utils"
	"github.com/steschwa/hopper-analytics-collector/constants"
	"github.com/steschwa/hopper-analytics-collector/graph"
)

func NewHistoryRouteHandler() controllers.RouteHandler {
	return func(ctx *fiber.Ctx) error {
		transferDirection := constants.TransferDirectionFromString(ctx.Query("direction", constants.TransferDirectionToUser.String()))
		user := ctx.Query("user")
		transferMethod := constants.TransferMethodFromString(ctx.Query("type", constants.TransferMethodAny.String()))

		transfersFilter := TransfersFilter{
			Direction: transferDirection,
			User:      user,
			Method:    transferMethod,
		}
		err := ValidateFilter(transfersFilter)
		if err != nil {
			log.Println(err)
			return controllers.CreateValidationError(ctx)
		}

		queryClient := graph.NewTransfersGraphClient()
		transfers, err := queryClient.FetchFilteredTransfers(graph.TransfersFilter{
			User:      transfersFilter.User,
			Direction: transfersFilter.Direction,
			MethodId:  transfersFilter.Method.ToMethodId(),
		})

		if err != nil {
			log.Println(err)
			sentry.CaptureException(err)
			return controllers.CreateServerError(ctx)
		}

		return ctx.JSON(fiber.Map{
			"fitler": fiber.Map{
				"user":      transfersFilter.User,
				"direction": transfersFilter.Direction.String(),
				"type":      transfersFilter.Method.String(),
			},
			"data": formatTransfers(transfers),
		})
	}
}

// ----------------------------------------
// Response formatters
// ----------------------------------------

func formatTransfers(transfers []graph.Transfer) []fiber.Map {
	data := []fiber.Map{}

	for _, transfer := range transfers {
		amount, _ := utils.ToDecimal(transfer.Amount.String(), 18).Float64()

		method := constants.TransferMethodFromMethodId(transfer.MethodId)
		if method == constants.TransferMethodAny {
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
