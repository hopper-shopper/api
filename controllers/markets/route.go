package markets

import (
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/steschwa/hopper-analytics-api/controllers"
	"github.com/steschwa/hopper-analytics-collector/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewMarketHistoryByHopper(mongoClient *mongo.Client) controllers.RouteHandler {
	return func(ctx *fiber.Ctx) error {
		tokenIds := TokenIdsFilterFromString(ctx.Query("tokenIds", ""))
		sold := SoldFilterFromString(ctx.Query("sold", AnySold.String()))

		marketsFilter := MarketsFilter{
			TokenIds: tokenIds,
			Sold:     sold,
		}

		err := ValidateFilter(marketsFilter)
		if err != nil {
			log.Println(err)
			return controllers.CreateValidationError(ctx)
		}

		loader := NewMarketListingsLoader(mongoClient)
		listings, err := loader.Load(marketsFilter)
		if err != nil {
			log.Println(err)
			sentry.CaptureException(err)
			return controllers.CreateServerError(ctx)
		}

		return ctx.JSON(fiber.Map{
			"filter": fiber.Map{
				"tokenIds": tokenIds,
			},
			"data": formatListings(listings),
		})
	}
}

// ----------------------------------------
// Response formatters
// ----------------------------------------

func formatListings(listings []models.ListingDocument) []fiber.Map {
	data := make([]fiber.Map, len(listings))
	for i, listing := range listings {
		data[i] = fiber.Map{
			"id":        listing.Id,
			"sold":      listing.Sold,
			"enabled":   listing.Enabled,
			"price":     listing.Price,
			"buyer":     listing.Buyer,
			"seller":    listing.Seller,
			"timestamp": listing.Timestamp,
			"tokenId":   listing.HopperId,
		}
	}

	return data
}
