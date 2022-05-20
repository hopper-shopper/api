package hoppers

import (
	"context"
	"errors"
	"time"

	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	HoppersActivityLoader struct {
		DbClient *db.MongoDbClient
	}

	HoppersActivityAggregate struct {
		Adventure   uint      `bson:"adventure"`
		Pond        uint      `bson:"pond"`
		Stream      uint      `bson:"stream"`
		Swamp       uint      `bson:"swamp"`
		River       uint      `bson:"river"`
		Forest      uint      `bson:"forest"`
		GreatLake   uint      `bson:"greatLake"`
		Breeding    uint      `bson:"breeding"`
		Marketplace uint      `bson:"marketplace"`
		Idle        uint      `bson:"idle"`
		Date        time.Time `bson:"date"`
	}
)

func NewHoppersActivityLoader(dbClient *db.MongoDbClient) *HoppersActivityLoader {
	return &HoppersActivityLoader{
		DbClient: dbClient,
	}
}

func (loader *HoppersActivityLoader) Latest() (models.HoppersActivityDocument, error) {
	collection := &db.HoppersActivityCollection{
		Client: loader.DbClient,
	}

	limit := int64(1)
	cursor, err := collection.GetCollection().Find(
		context.Background(),
		bson.D{{}},
		&options.FindOptions{
			Sort: bson.D{{
				Key:   "timestamp",
				Value: -1,
			}},
			Limit: &limit,
		},
	)
	if err != nil {
		return models.HoppersActivityDocument{}, err
	}

	activities := []models.HoppersActivityDocument{}
	if err = cursor.All(context.Background(), &activities); err != nil {
		return models.HoppersActivityDocument{}, err
	}

	if len(activities) == 0 || len(activities) > 1 {
		return models.HoppersActivityDocument{}, errors.New("failed to fetch exactly one activity")
	}

	return activities[0], nil
}

func (loader *HoppersActivityLoader) History() ([]HoppersActivityAggregate, error) {
	collection := &db.HoppersActivityCollection{
		Client: loader.DbClient,
	}

	pipeline := mongo.Pipeline{
		getAggregationGrouping(),
		getAggregationProjection(),
		getAggregationSort(),
	}

	cursor, err := collection.GetCollection().Aggregate(
		context.Background(),
		pipeline,
	)
	if err != nil {
		return []HoppersActivityAggregate{}, err
	}

	aggregates := []HoppersActivityAggregate{}
	if err = cursor.All(context.Background(), &aggregates); err != nil {
		return []HoppersActivityAggregate{}, err
	}

	return aggregates, nil
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
			"timestamp": bson.M{
				"$first": bson.M{
					"$dateTrunc": bson.M{
						"date": "$timestamp",
						"unit": "day",
					},
				},
			},
			"adventure": bson.M{
				"$avg": "$adventure",
			},
			"pond": bson.M{
				"$avg": "$pond",
			},
			"stream": bson.M{
				"$avg": "$stream",
			},
			"swamp": bson.M{
				"$avg": "$swamp",
			},
			"river": bson.M{
				"$avg": "$river",
			},
			"forest": bson.M{
				"$avg": "$forest",
			},
			"greatLake": bson.M{
				"$avg": "$greatLake",
			},
			"breeding": bson.M{
				"$avg": "$breeding",
			},
			"marketplace": bson.M{
				"$avg": "$marketplace",
			},
			"idle": bson.M{
				"$avg": "$idle",
			},
		}}}
}

func getAggregationProjection() bson.D {
	return bson.D{{
		Key: "$project",
		Value: bson.M{
			"_id":  0,
			"date": "$timestamp",
			"adventure": bson.M{
				"$trunc": "$adventure",
			},
			"pond": bson.M{
				"$trunc": "$adventure",
			},
			"stream": bson.M{
				"$trunc": "$stream",
			},
			"swamp": bson.M{
				"$trunc": "$swamp",
			},
			"river": bson.M{
				"$trunc": "$river",
			},
			"forest": bson.M{
				"$trunc": "$forest",
			},
			"greatLake": bson.M{
				"$trunc": "$greatLake",
			},
			"breeding": bson.M{
				"$trunc": "$breeding",
			},
			"marketplace": bson.M{
				"$trunc": "$marketplace",
			},
			"idle": bson.M{
				"$trunc": "$idle",
			},
		},
	}}
}

func getAggregationSort() bson.D {
	return bson.D{{
		Key: "$sort",
		Value: bson.M{
			"date": 1,
		},
	}}
}
