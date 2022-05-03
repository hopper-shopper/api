package prices

import (
	"log"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/steschwa/hopper-analytics-api/controllers"
	"github.com/steschwa/hopper-analytics-api/prices"
	"github.com/steschwa/hopper-analytics-collector/constants"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewLatestPriceRouteHandler(mongoClient *mongo.Client) controllers.RouteHandler {
	return func(ctx *fiber.Ctx) error {
		loader := prices.NewPriceLoader(mongoClient)

		avaxUsd := loader.LoadLatestAvaxUsdPrice()
		avaxEur := loader.LoadLatestAvaxEurPrice()
		flyUsd := loader.LoadLatestFlyUsdPrice()
		flyEur := loader.LoadLatestFlyEurPrice()
		var avaxFly = 0.0
		if flyUsd != 0 {
			avaxFly = avaxUsd / flyUsd
		}
		var flyAvax = 0.0
		if avaxUsd != 0 {
			flyAvax = flyUsd / avaxUsd
		}

		return ctx.JSON(fiber.Map{
			"data": fiber.Map{
				"AVAX": fiber.Map{
					"EUR": avaxEur,
					"USD": avaxUsd,
					"FLY": avaxFly,
				},
				"FLY": fiber.Map{
					"EUR":  flyEur,
					"USD":  flyUsd,
					"AVAX": flyAvax,
				},
			},
		})
	}
}

func NewHistoricalPriceRouteHandler(mongoClient *mongo.Client) controllers.RouteHandler {
	return func(ctx *fiber.Ctx) error {
		times := strings.Split(ctx.Query("at"), ",")
		coin := ctx.Query("coin", "avax")
		currency := ctx.Query("currency", "usd")

		filter := HistoricalPriceFilter{
			Times:    times,
			Coin:     coin,
			Currency: currency,
		}

		if err := ValidateFilter(filter); err != nil {
			log.Println(err)
			return controllers.CreateValidationError(ctx)
		}

		loader := prices.NewPriceLoader(mongoClient)
		priceFilter := prices.PriceFilter{
			Coin:     constants.CoingeckoIdFromString(filter.Coin, constants.COINGECKO_AVAX),
			Currency: constants.CoingeckoCurrenyFromString(filter.Currency),
		}

		prices := fiber.Map{}
		for _, timeStr := range times {
			time, err := time.Parse(time.RFC3339, timeStr)
			if err != nil {
				log.Println(err)
				sentry.CaptureException(err)
				return controllers.CreateServerError(ctx)
			}

			price := loader.LoadHistoricalPrice(priceFilter, time)
			prices[timeStr] = price
		}

		return ctx.JSON(fiber.Map{
			"filter": fiber.Map{
				"at":       filter.Times,
				"coin":     priceFilter.Coin.String(),
				"currency": priceFilter.Currency.String(),
			},
			"data": prices,
		})
	}
}
