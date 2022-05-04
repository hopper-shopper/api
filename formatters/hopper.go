package formatters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/steschwa/hopper-analytics-api/prices"
	"github.com/steschwa/hopper-analytics-collector/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	HopperFormatter struct {
		AvaxPriceUsd float64
		FlyPriceUsd  float64
	}
)

func NewHopperFormatter(mongoClient *mongo.Client) *HopperFormatter {
	loader := prices.NewPriceLoader(mongoClient)
	avaxUsd := loader.LoadLatestAvaxUsdPrice()
	flyUsd := loader.LoadLatestFlyUsdPrice()

	return &HopperFormatter{
		AvaxPriceUsd: avaxUsd,
		FlyPriceUsd:  flyUsd,
	}
}

func (formatter *HopperFormatter) Format(hopper models.HopperDocument) fiber.Map {
	cumulatedLevelCosts := prices.CalculateCumulatedLevelCost(uint8(hopper.Level))
	avaxPerFly := formatter.FlyPriceUsd / formatter.AvaxPriceUsd
	avaxLevelCosts := cumulatedLevelCosts * avaxPerFly

	return fiber.Map{
		"tokenId":      hopper.TokenId,
		"strength":     hopper.Strength,
		"agility":      hopper.Agility,
		"vitality":     hopper.Vitality,
		"intelligence": hopper.Intelligence,
		"fertility":    hopper.Fertility,
		"level":        hopper.Level,
		"levelCosts":   avaxLevelCosts,
		"image":        hopper.Image,
		"activity":     hopper.Activity,
		"rating": fiber.Map{
			"pond":      hopper.RatingPond,
			"stream":    hopper.RatingStream,
			"swamp":     hopper.RatingSwamp,
			"river":     hopper.RatingRiver,
			"forest":    hopper.RatingForest,
			"greatLake": hopper.RatingGreatLake,
		},
		"baseFly": fiber.Map{
			"pond":      hopper.BaseFlyPond,
			"stream":    hopper.BaseFlyStream,
			"swamp":     hopper.BaseFlySwamp,
			"river":     hopper.BaseFlyRiver,
			"forest":    hopper.BaseFlyForest,
			"greatLake": hopper.BaseFlyGreatLake,
		},
		"price": hopper.MarketPrice,
	}
}

func (formatter *HopperFormatter) FormatAll(hoppers []models.HopperDocument) []fiber.Map {
	data := make([]fiber.Map, len(hoppers))
	for i, hopper := range hoppers {
		data[i] = formatter.Format(hopper)
	}

	return data
}
