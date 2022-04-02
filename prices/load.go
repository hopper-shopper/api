package prices

import (
	"context"

	"github.com/steschwa/hopper-analytics-collector/constants"
	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	PriceLoader struct {
		MongoClient *mongo.Client
	}
)

func NewPriceLoader(mongoClient *mongo.Client) *PriceLoader {
	return &PriceLoader{
		MongoClient: mongoClient,
	}
}

func (loader *PriceLoader) loadPrice(filter bson.D) float64 {
	pricesCollection := &db.PricesCollection{
		Connection: loader.MongoClient,
	}
	result := pricesCollection.GetCollection().FindOne(
		context.Background(),
		filter,
	)

	priceDocument := models.PriceDocument{}
	if err := result.Decode(&priceDocument); err != nil {
		return 0
	}

	return priceDocument.Price
}

func (loader *PriceLoader) LoadAvaxUsdPrice() float64 {
	return loader.loadPrice(bson.D{
		{Key: "coin", Value: constants.COINGECKO_AVAX},
		{Key: "currency", Value: constants.COINGECKO_USD},
	})
}
func (loader *PriceLoader) LoadAvaxEurPrice() float64 {
	return loader.loadPrice(bson.D{
		{Key: "coin", Value: constants.COINGECKO_AVAX},
		{Key: "currency", Value: constants.COINGECKO_EUR},
	})
}

func (loader *PriceLoader) LoadFlyUsdPrice() float64 {
	return loader.loadPrice(bson.D{
		{Key: "coin", Value: constants.COINGECKO_FLY},
		{Key: "currency", Value: constants.COINGECKO_USD},
	})
}
func (loader *PriceLoader) LoadFlyEurPrice() float64 {
	return loader.loadPrice(bson.D{
		{Key: "coin", Value: constants.COINGECKO_FLY},
		{Key: "currency", Value: constants.COINGECKO_EUR},
	})
}
