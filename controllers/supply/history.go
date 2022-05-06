package supply

import (
	"context"
	"time"

	db "github.com/steschwa/hopper-analytics-collector/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	SupplyHistoryLoader struct {
		Client *db.MongoDbClient
	}

	AggregatedSupply struct {
		Datetime  time.Time `bson:"datetime"`
		Supply    float64   `bson:"supply"`
		Burned    float64   `bson:"burned"`
		Staked    float64   `bson:"staked"`
		Available float64   `bson:"available"`
		Free      float64   `bson:"free"`
	}
)

func NewSupplyHistoryLoader(dbClient *db.MongoDbClient) *SupplyHistoryLoader {
	return &SupplyHistoryLoader{
		Client: dbClient,
	}
}

func (loader *SupplyHistoryLoader) Load() ([]AggregatedSupply, error) {
	collection := &db.FlySuppliesCollection{
		Client: loader.Client,
	}
	cursor, err := collection.GetCollection().Aggregate(
		context.Background(),
		getAggregationPipeline(),
	)
	if err != nil {
		return []AggregatedSupply{}, err
	}

	supplies := []AggregatedSupply{}
	if err = cursor.All(context.Background(), &supplies); err != nil {
		return []AggregatedSupply{}, err
	}

	return supplies, nil
}

func getAggregationPipeline() mongo.Pipeline {
	return mongo.Pipeline{
		getAggregationGrouping(),
		getAggregationProjection(),
		getAggregationSort(),
	}
}
func getAggregationGrouping() bson.D {
	return bson.D{{
		Key: "$group",
		Value: bson.M{
			"_id": bson.M{
				"$dateToString": bson.D{
					{Key: "format", Value: "%Y-%m-%d"},
					{Key: "date", Value: "$timestamp"},
				},
			},
			"datetime": bson.M{
				"$first": bson.M{
					"$dateTrunc": bson.M{
						"date": "$timestamp",
						"unit": "day",
					},
				},
			},
			"supply": bson.M{
				"$avg": "$supply",
			},
			"burned": bson.M{
				"$avg": "$burned",
			},
			"staked": bson.M{
				"$avg": "$staked",
			},
			"available": bson.M{
				"$avg": "$available",
			},
			"free": bson.M{
				"$avg": "$free",
			},
		},
	}}
}
func getAggregationProjection() bson.D {
	return bson.D{{
		Key: "$project",
		Value: bson.M{
			"_id":       0,
			"datetime":  "$datetime",
			"supply":    1,
			"burned":    1,
			"staked":    1,
			"available": 1,
			"free":      1,
		},
	}}
}
func getAggregationSort() bson.D {
	return bson.D{{
		Key: "$sort",
		Value: bson.M{
			"datetime": 1,
		},
	}}
}
