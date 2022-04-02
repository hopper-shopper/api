package prices

import (
	"context"

	"github.com/steschwa/hopper-analytics-collector/constants"
	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoadAvaxUsdPrice(mongoClient *mongo.Client) float64 {
	pricesCollection := &db.PricesCollection{
		Connection: mongoClient,
	}
	result := pricesCollection.GetCollection().FindOne(
		context.Background(),
		bson.D{
			{Key: "coin", Value: constants.COINGECKO_AVAX},
			{Key: "currency", Value: constants.COINGECKO_USD},
		},
	)

	priceDocument := models.PriceDocument{}
	if err := result.Decode(&priceDocument); err != nil {
		return 0
	}

	return priceDocument.Price
}

func LoadFlyUsdPrice(mongoClient *mongo.Client) float64 {
	pricesCollection := &db.PricesCollection{
		Connection: mongoClient,
	}
	result := pricesCollection.GetCollection().FindOne(
		context.Background(),
		bson.D{
			{Key: "coin", Value: constants.COINGECKO_FLY},
			{Key: "currency", Value: constants.COINGECKO_USD},
		},
	)

	priceDocument := models.PriceDocument{}
	if err := result.Decode(&priceDocument); err != nil {
		return 0
	}

	return priceDocument.Price
}
