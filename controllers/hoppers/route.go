package hoppers

import (
	"context"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/steschwa/hopper-analytics-api/controllers"
	"github.com/steschwa/hopper-analytics-api/formatters"
	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouteHandler(mongoClient *mongo.Client) controllers.RouteHandler {
	return func(ctx *fiber.Ctx) error {
		adventure := AdventureFilterFromString(ctx.Query("adventure", AnyAdventure.String()))
		permit := PermitFilterFromString(ctx.Query("permit", AnyPermit.String()))
		market := MarketFilterFromString(ctx.Query("market", AnyMarket.String()))
		tokenIds := TokenIdsFilterFromString(ctx.Query("tokenIds", ""))
		owner := ctx.Query("owner")

		hoppersFilter := HoppersFilter{
			Adventure: adventure,
			Market:    market,
			Permit:    permit,
			TokenIds:  tokenIds,
			Owner:     owner,
		}

		err := ValidateFilter(hoppersFilter)
		if err != nil {
			log.Println(err)
			return controllers.CreateValidationError(ctx)
		}

		hoppersCollection := &db.HoppersCollection{
			Connection: mongoClient,
		}
		cursor, err := hoppersCollection.GetCollection().Find(
			context.Background(),
			getMongoFilter(hoppersFilter),
		)
		if err != nil {
			log.Println(err)
			return controllers.CreateServerError(ctx)
		}

		hoppers := []models.HopperDocument{}
		if err = cursor.All(context.Background(), &hoppers); err != nil {
			log.Println(err)
			return controllers.CreateServerError(ctx)
		}

		formatter := formatters.NewHopperFormatter(mongoClient)

		return ctx.JSON(fiber.Map{
			"filter": fiber.Map{
				"adventure": adventure.String(),
				"permit":    permit.String(),
				"market":    market.String(),
				"tokenIds":  tokenIds,
			},
			"data": formatter.FormatAll(hoppers),
		})
	}
}

// ----------------------------------------
// Mongo filters
// ----------------------------------------

func getMongoFilter(hoppersFilter HoppersFilter) bson.D {
	filter := bson.D{}

	filter = append(filter, getAdventureFilter(hoppersFilter.Adventure))
	filter = append(filter, getPermitFilter(hoppersFilter.Permit))
	filter = append(filter, getMarketFilter(hoppersFilter.Market))
	filter = append(filter, getTokenIdsFilter(hoppersFilter.TokenIds))
	filter = append(filter, getOwnerFilter(hoppersFilter.Owner))

	return filter
}
func getAdventureFilter(adventureFilter AdventureFilter) bson.E {
	switch adventureFilter {
	case PondAdventure:
		return bson.E{Key: "adventure", Value: "pond"}
	case StreamAdventure:
		return bson.E{Key: "adventure", Value: "stream"}
	case SwampAdventure:
		return bson.E{Key: "adventure", Value: "swamp"}
	case RiverAdventure:
		return bson.E{Key: "adventure", Value: "river"}
	case ForestAdventure:
		return bson.E{Key: "adventure", Value: "forest"}
	case GreatLakeAdventure:
		return bson.E{Key: "adventure", Value: "great-lake"}
	default:
		return bson.E{}
	}
}
func getPermitFilter(permitFilter PermitFilter) bson.E {
	switch permitFilter {
	case RiverPermit:
		return bson.E{Key: "canEnterRiver", Value: true}
	case ForestPermit:
		return bson.E{Key: "canEnterForest", Value: true}
	case GreatLakePermit:
		return bson.E{Key: "canEnterGreatLake", Value: true}
	default:
		return bson.E{}
	}
}
func getMarketFilter(marketFilter MarketFilter) bson.E {
	switch marketFilter {
	case OnMarket:
		return bson.E{Key: "listingActive", Value: true}
	case OffMarket:
		return bson.E{Key: "listingActive", Value: false}
	default:
		return bson.E{}
	}
}
func getTokenIdsFilter(tokenIds []string) bson.E {
	if len(tokenIds) == 0 {
		return bson.E{}
	}

	return bson.E{
		Key: "tokenId",
		Value: bson.M{
			"$in": tokenIds,
		},
	}
}
func getOwnerFilter(owner string) bson.E {
	if owner == "" {
		return bson.E{}
	}

	return bson.E{
		Key:   "owner",
		Value: strings.ToLower(owner),
	}
}
