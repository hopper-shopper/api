package prices

import (
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/steschwa/hopper-analytics-api/controllers"
	"github.com/steschwa/hopper-analytics-api/prices"
	"github.com/steschwa/hopper-analytics-collector/constants"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
)

func NewHistoricalPriceRouteHandler(dbClient *db.MongoDbClient) controllers.RouteHandler {
	return func(ctx *fiber.Ctx) error {
		filter := HistoricalPriceFilter{
			Coin:     "avax",
			Currency: "usd",
		}

		if err := ctx.BodyParser(&filter); err != nil {
			log.Println(err)
			sentry.CaptureException(err)
			return controllers.CreateBadRequestError(ctx)
		}

		if err := ValidateFilter(filter); err != nil {
			log.Println(err)
			return controllers.CreateValidationError(ctx)
		}

		loader := prices.NewPriceLoader(dbClient)
		priceFilter := prices.PriceFilter{
			Coin:     constants.CoingeckoIdFromString(filter.Coin, constants.COINGECKO_AVAX),
			Currency: constants.CoingeckoCurrenyFromString(filter.Currency),
		}

		prices := fiber.Map{}
		for _, dateStr := range filter.Dates {
			time, err := time.Parse(time.RFC3339, dateStr)
			if err != nil {
				log.Println(err)
				sentry.CaptureException(err)
				return controllers.CreateServerError(ctx)
			}

			price := loader.LoadHistoricalPrice(priceFilter, time)
			prices[dateStr] = price
		}

		return ctx.JSON(fiber.Map{
			"filter": fiber.Map{
				"dates":    filter.Dates,
				"coin":     priceFilter.Coin.String(),
				"currency": priceFilter.Currency.String(),
			},
			"data": prices,
		})
	}
}
