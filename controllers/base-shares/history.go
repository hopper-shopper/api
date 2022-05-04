package baseshares

import (
	"context"
	"time"

	"github.com/steschwa/hopper-analytics-collector/constants"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	BaseSharesHistoryLoader struct {
		DbClient *db.MongoDbClient
	}

	AggregatedBaseShare struct {
		Datetime        time.Time `bson:"datetime"`
		Adventure       string    `bson:"adventure"`
		TotalBaseShares float64   `bson:"totalBaseShares"`
	}
)

func NewBaseSharesHistoryLoader(dbClient *db.MongoDbClient) *BaseSharesHistoryLoader {
	return &BaseSharesHistoryLoader{
		DbClient: dbClient,
	}
}

func (loader *BaseSharesHistoryLoader) Load(filter BaseShareByAdventureFilter) ([]AggregatedBaseShare, error) {
	collection := &db.BaseSharesCollection{
		Client: loader.DbClient,
	}
	cursor, err := collection.GetCollection().Aggregate(
		context.Background(),
		getAggregationPipeline(filter.Adventure),
	)
	if err != nil {
		return []AggregatedBaseShare{}, err
	}

	baseShares := []AggregatedBaseShare{}
	if err = cursor.All(context.Background(), &baseShares); err != nil {
		return []AggregatedBaseShare{}, err
	}

	return baseShares, nil
}

func getAggregationPipeline(adventure constants.Adventure) mongo.Pipeline {
	return mongo.Pipeline{
		getAggregationFilter(adventure),
		getAggregationGrouping(),
		getAggregationProjection(),
		getAggregationSort(),
	}
}
func getAggregationFilter(adventure constants.Adventure) bson.D {
	return bson.D{{
		Key: "$match",
		Value: bson.M{
			"updated": bson.M{
				"$gt": primitive.NewDateTimeFromTime(time.Now().Add(time.Hour * 24 * 30 * -1)),
			},
			"adventure": adventure.String(),
		},
	}}
}
func getAggregationGrouping() bson.D {
	dateFormat := "%Y-%m-%d"

	return bson.D{{
		Key: "$group",
		Value: bson.M{
			"_id": bson.M{
				"$dateToString": bson.D{
					{Key: "format", Value: dateFormat},
					{Key: "date", Value: "$updated"},
				},
			},
			"datetime": bson.M{
				"$first": bson.M{
					"$dateTrunc": bson.M{
						"date": "$updated",
						"unit": "day",
					},
				},
			},
			"adventure": bson.M{
				"$first": "$adventure",
			},
			"totalBaseShares": bson.M{
				"$avg": bson.M{
					"$toLong": "$totalBaseShares.v",
				},
			},
		},
	}}
}
func getAggregationProjection() bson.D {
	return bson.D{{
		Key: "$project",
		Value: bson.M{
			"_id":             0,
			"datetime":        "$datetime",
			"totalBaseShares": 1,
			"adventure":       1,
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
