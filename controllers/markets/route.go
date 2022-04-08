package markets

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/steschwa/hopper-analytics-api/controllers"
	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewMarketHistoryRouteHandler(mongoClient *mongo.Client) controllers.RouteHandler {
	return func(ctx *fiber.Ctx) error {
		tokenIds := TokenIdsFilterFromString(ctx.Query("tokenIds", ""))

		marketsFilter := MarketsFilter{
			TokenIds: tokenIds,
		}

		err := ValidateFilter(marketsFilter)
		if err != nil {
			log.Println(err)
			return controllers.CreateValidationError(ctx)
		}

		listingsCollection := &db.MarketsCollection{
			Connection: mongoClient,
		}
		cursor, err := listingsCollection.GetCollection().Find(
			context.Background(),
			getListingsFilter(marketsFilter),
		)
		if err != nil {
			log.Println(err)
			return controllers.CreateServerError(ctx)
		}

		listings := []models.ListingDocument{}
		if err = cursor.All(context.Background(), &listings); err != nil {
			log.Println(err)
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
// Mongo filters
// ----------------------------------------

func getListingsFilter(marketFilter MarketsFilter) bson.D {
	if len(marketFilter.TokenIds) == 0 {
		return bson.D{}
	}

	filter := bson.D{}

	filter = append(filter, bson.E{
		Key: "hopperId",
		Value: bson.M{
			"$in": marketFilter.TokenIds,
		},
	})
	filter = append(filter, bson.E{
		Key:   "sold",
		Value: true,
	})

	return filter
}

func formatListings(listings []models.ListingDocument) []fiber.Map {
	data := make([]fiber.Map, len(listings))
	for i, listing := range listings {
		data[i] = fiber.Map{
			"sold":      listing.Sold,
			"price":     listing.Price,
			"buyer":     listing.Buyer,
			"timestamp": listing.Timestamp,
			"tokenId":   listing.HopperId,
		}
	}

	return data
}
