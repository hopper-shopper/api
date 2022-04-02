package compare

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
		tokenIds := parseTokenIds(ctx.Query("hoppers"))

		if len(tokenIds) == 0 {
			return ctx.JSON(fiber.Map{
				"hoppers": []string{},
				"data":    fiber.Map{},
			})
		}

		hoppersCollection := &db.HoppersCollection{
			Connection: mongoClient,
		}

		cursor, err := hoppersCollection.GetCollection().Find(
			context.Background(),
			getHoppersFilter(tokenIds),
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
			"hoppers": tokenIds,
			"data":    formatter.FormatAll(hoppers),
		})
	}
}

// ----------------------------------------
// Input parsing
// ----------------------------------------

func parseTokenIds(value string) []string {
	tokenIds := strings.Split(value, ",")

	for i, tokenId := range tokenIds {
		tokenIds[i] = strings.Replace(tokenId, " ", "", -1)
	}

	return tokenIds
}

// ----------------------------------------
// Mongo filters
// ----------------------------------------

func getHoppersFilter(tokenIds []string) bson.M {
	return bson.M{
		"tokenId": bson.M{
			"$in": tokenIds,
		},
	}
}
