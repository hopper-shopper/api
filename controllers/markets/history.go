package markets

import "go.mongodb.org/mongo-driver/mongo"

type (
	MarketHistoryLoader struct {
		Mongo *mongo.Client
	}
)
