package markets

import (
	"context"

	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	MarketListingsLoader struct {
		Mongo *mongo.Client
	}
)

func NewMarketListingsLoader(mongoClient *mongo.Client) *MarketListingsLoader {
	return &MarketListingsLoader{
		Mongo: mongoClient,
	}
}

func (loader *MarketListingsLoader) Load(filter MarketsFilter) ([]models.ListingDocument, error) {
	listingsCollection := &db.MarketsCollection{
		Connection: loader.Mongo,
	}
	cursor, err := listingsCollection.GetCollection().Find(
		context.Background(),
		getListingsFilter(filter),
	)
	if err != nil {
		return []models.ListingDocument{}, err
	}

	listings := []models.ListingDocument{}
	if err = cursor.All(context.Background(), &listings); err != nil {
		return []models.ListingDocument{}, err
	}

	return listings, nil
}

func getListingsFilter(marketFilter MarketsFilter) bson.D {
	filter := bson.D{}

	filter = append(filter, getTokenIdsFilter(marketFilter.TokenIds))
	filter = append(filter, getSoldFilter(marketFilter.Sold))

	return filter
}
func getTokenIdsFilter(tokenIds []string) bson.E {
	if len(tokenIds) == 0 {
		return bson.E{}
	}

	return bson.E{
		Key: "hopperId",
		Value: bson.M{
			"$in": tokenIds,
		},
	}
}
func getSoldFilter(sold SoldFilter) bson.E {
	switch sold {
	case AlreadySold:
		return bson.E{Key: "sold", Value: true}
	case NotSold:
		return bson.E{Key: "sold", Value: false}
	default:
		return bson.E{}
	}
}
