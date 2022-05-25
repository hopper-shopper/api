package supply

import (
	"context"

	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	SupplyHistoryLoader struct {
		Client *db.MongoDbClient
	}
)

func NewSupplyHistoryLoader(dbClient *db.MongoDbClient) *SupplyHistoryLoader {
	return &SupplyHistoryLoader{
		Client: dbClient,
	}
}

func (loader *SupplyHistoryLoader) Load() ([]models.FlySupplyDocument, error) {
	collection := &db.FlySuppliesCollection{
		Client: loader.Client,
	}

	cursor, err := collection.GetCollection().Find(
		context.Background(),
		bson.D{{}},
		&options.FindOptions{
			Sort: bson.D{{
				Key:   "timestamp",
				Value: 1,
			}},
		},
	)
	if err != nil {
		return []models.FlySupplyDocument{}, err
	}

	supplies := []models.FlySupplyDocument{}
	if err = cursor.All(context.Background(), &supplies); err != nil {
		return []models.FlySupplyDocument{}, err
	}

	return supplies, nil
}
