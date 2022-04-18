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
		Mongo *mongo.Client
	}

	AggregatedSupply struct {
		Datetime time.Time `bson:"datetime"`
		Supply   float64   `bson:"supply"`
	}
)

func NewSupplyHistoryLoader(mongoClient *mongo.Client) *SupplyHistoryLoader {
	return &SupplyHistoryLoader{
		Mongo: mongoClient,
	}
}

func (loader *SupplyHistoryLoader) Load(filter SupplyFilter) ([]AggregatedSupply, error) {
	collection := &db.SuppliesCollection{
		Connection: loader.Mongo,
	}
	cursor, err := collection.GetCollection().Aggregate(
		context.Background(),
		getAggregationPipeline(filter),
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

func getAggregationPipeline(filter SupplyFilter) mongo.Pipeline {
	return mongo.Pipeline{
		getAggregationFilter(filter),
		getAggregationGrouping(),
		getAggregationProjection(),
		getAggregationSort(),
	}
}
func getAggregationFilter(filter SupplyFilter) bson.D {
	return bson.D{{
		Key: "$match",
		Value: bson.M{
			"type": filter.Type.String(),
		},
	}}
}
func getAggregationGrouping() bson.D {
	dateFormat := "%Y-%m-%d-%H"

	return bson.D{{
		Key: "$group",
		Value: bson.M{
			"_id": bson.M{
				"$dateToString": bson.D{
					{Key: "format", Value: dateFormat},
					{Key: "date", Value: "$timestamp"},
				},
			},
			"datetime": bson.M{
				"$first": bson.M{
					"$dateTrunc": bson.M{
						"date": "$timestamp",
						"unit": "hour",
					},
				},
			},
			"supply": bson.M{
				"$avg": "$supply",
			},
		},
	}}
}
func getAggregationProjection() bson.D {
	return bson.D{{
		Key: "$project",
		Value: bson.M{
			"_id":      0,
			"datetime": "$datetime",
			"supply":   1,
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
