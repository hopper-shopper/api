package prices

import (
	"context"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/steschwa/hopper-analytics-collector/constants"
	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	PriceLoader struct {
		Mongo *mongo.Client
	}

	PriceFilter struct {
		Coin     constants.CoinGeckoId
		Currency constants.CoinGeckoCurrency
	}

	AggregatedPrice struct {
		Price float64 `bson:"price"`
	}
)

func NewPriceLoader(mongoClient *mongo.Client) *PriceLoader {
	return &PriceLoader{
		Mongo: mongoClient,
	}
}

func (loader *PriceLoader) loadPrice(filter PriceFilter) float64 {
	pricesCollection := &db.PricesCollection{
		Connection: loader.Mongo,
	}
	cursor := pricesCollection.GetCollection().FindOne(
		context.Background(),
		filter.ToMongoFilter(),
	)

	priceDocument := models.PriceDocument{}
	if err := cursor.Decode(&priceDocument); err != nil {
		log.Println(err)
		sentry.CaptureException(err)
		return 0
	}

	return priceDocument.Price
}

func (loader *PriceLoader) LoadHistoricalPrice(filter PriceFilter, at time.Time) float64 {
	pricesCollection := &db.PricesCollection{
		Connection: loader.Mongo,
	}

	grouping := bson.D{{
		Key: "$group",
		Value: bson.M{
			"_id": bson.M{
				"$dateToString": bson.D{
					{Key: "format", Value: "%Y-%m-%d"},
					{Key: "date", Value: "$timestamp"},
				},
			},
			"price": bson.M{
				"$avg": "$price",
			},
		},
	}}

	projection := bson.D{{
		Key: "$project",
		Value: bson.M{
			"_id":   0,
			"price": 1,
		},
	}}

	pipeline := mongo.Pipeline{}
	pipeline = append(pipeline, filter.ToMongoAggregationFilter(at))
	pipeline = append(pipeline, grouping)
	pipeline = append(pipeline, projection)

	cursor, err := pricesCollection.GetCollection().Aggregate(
		context.Background(),
		pipeline,
	)
	if err != nil {
		log.Println(err)
		sentry.CaptureException(err)
		return 0
	}
	defer cursor.Close(context.Background())

	prices := []AggregatedPrice{}
	if err = cursor.All(context.Background(), &prices); err != nil {
		log.Println(err)
		sentry.CaptureException(err)
		return 0
	}

	if len(prices) == 0 {
		return 0
	}

	if len(prices) == 1 {
		return prices[0].Price
	}

	total := 0.0
	for _, price := range prices {
		total += price.Price
	}

	return total / float64(len(prices))
}

func (loader *PriceLoader) LoadLatestAvaxUsdPrice() float64 {
	return loader.loadPrice(PriceFilter{
		Coin:     constants.COINGECKO_AVAX,
		Currency: constants.COINGECKO_USD,
	})
}
func (loader *PriceLoader) LoadLatestAvaxEurPrice() float64 {
	return loader.loadPrice(PriceFilter{
		Coin:     constants.COINGECKO_AVAX,
		Currency: constants.COINGECKO_EUR,
	})
}

func (loader *PriceLoader) LoadLatestFlyUsdPrice() float64 {
	return loader.loadPrice(PriceFilter{
		Coin:     constants.COINGECKO_FLY,
		Currency: constants.COINGECKO_USD,
	})
}
func (loader *PriceLoader) LoadLatestFlyEurPrice() float64 {
	return loader.loadPrice(PriceFilter{
		Coin:     constants.COINGECKO_FLY,
		Currency: constants.COINGECKO_EUR,
	})
}

func (loader *PriceLoader) LoadHistoricalAvaxUsdPrice(at time.Time) float64 {
	return loader.LoadHistoricalPrice(PriceFilter{
		Coin:     constants.COINGECKO_AVAX,
		Currency: constants.COINGECKO_USD,
	}, at)
}
func (loader *PriceLoader) LoadHistoricalAvaxEurPrice(at time.Time) float64 {
	return loader.LoadHistoricalPrice(PriceFilter{
		Coin:     constants.COINGECKO_AVAX,
		Currency: constants.COINGECKO_EUR,
	}, at)
}

func (loader *PriceLoader) LoadHistoricalFlyUsdPrice(at time.Time) float64 {
	return loader.LoadHistoricalPrice(PriceFilter{
		Coin:     constants.COINGECKO_FLY,
		Currency: constants.COINGECKO_USD,
	}, at)
}
func (loader *PriceLoader) LoadHistoricalFlyEurPrice(at time.Time) float64 {
	return loader.LoadHistoricalPrice(PriceFilter{
		Coin:     constants.COINGECKO_FLY,
		Currency: constants.COINGECKO_EUR,
	}, at)
}

func (filter PriceFilter) ToMongoFilter() bson.D {
	return bson.D{
		{Key: "coin", Value: filter.Coin},
		{Key: "currency", Value: filter.Currency},
	}
}

func (filter PriceFilter) ToMongoAggregationFilter(at time.Time) bson.D {
	from := at.Add(time.Minute * 10 * -1)
	to := at.Add(time.Minute * 10)

	return bson.D{{
		Key: "$match",
		Value: bson.M{
			"coin":     filter.Coin,
			"currency": filter.Currency,
			"timestamp": bson.M{
				"$gte": primitive.NewDateTimeFromTime(from),
				"$lte": primitive.NewDateTimeFromTime(to),
			},
		},
	}}
}
