package formatters

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/steschwa/hopper-analytics-api/prices"
	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
)

type (
	HopperFormatter struct {
		AvaxPriceUsd float64
		FlyPriceUsd  float64
	}
)

func NewHopperFormatter(dbClient *db.MongoDbClient) *HopperFormatter {
	loader := prices.NewPriceLoader(dbClient)
	avaxUsd := loader.LoadLatestAvaxUsdPrice()
	flyUsd := loader.LoadLatestFlyUsdPrice()

	return &HopperFormatter{
		AvaxPriceUsd: avaxUsd,
		FlyPriceUsd:  flyUsd,
	}
}

func (formatter *HopperFormatter) GetMarketPrice(hopper models.HopperDocument) float64 {
	marketPrice := hopper.MarketPrice
	if hopper.Activity != models.HopperActivityMarketplace {
		marketPrice = 0
	}

	return marketPrice
}

func (formatter *HopperFormatter) GetLevelCosts(hopper models.HopperDocument) float64 {
	cumulatedLevelCosts := prices.CalculateCumulatedLevelCost(uint8(hopper.Level))
	avaxPerFly := formatter.FlyPriceUsd / formatter.AvaxPriceUsd
	return cumulatedLevelCosts * avaxPerFly
}

func (formatter *HopperFormatter) FormatJson(hopper models.HopperDocument) fiber.Map {
	return fiber.Map{
		"tokenId":      hopper.TokenId,
		"strength":     hopper.Strength,
		"agility":      hopper.Agility,
		"vitality":     hopper.Vitality,
		"intelligence": hopper.Intelligence,
		"fertility":    hopper.Fertility,
		"level":        hopper.Level,
		"levelCosts":   formatter.GetLevelCosts(hopper),
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
		"price": formatter.GetMarketPrice(hopper),
	}
}

func (formatter *HopperFormatter) FormatAllJson(hoppers []models.HopperDocument) []fiber.Map {
	data := make([]fiber.Map, len(hoppers))
	for i, hopper := range hoppers {
		data[i] = formatter.FormatJson(hopper)
	}

	return data
}

// TODO Make formatting more readable
func (formatter *HopperFormatter) FormatCsv(hopper models.HopperDocument) string {
	// "tokenId;strength;agility;vitality;intelligence;fertility;level;levelCosts;price;activity;image;ratingPong;ratingStream;ratingSwap;ratingRiver;ratingForest;ratingGreatLake;baseFlyPond;baseFlyStream;baseFlySwamp;baseFlyRiver;baseFlyForest;baseFlyGreatLake"
	return fmt.Sprintf(`"%s";%d;%d;%d;%d;%d;%d;%.4f;%.4f;"%s";"%s";%.4f;%.4f;%.4f;%.4f;%.4f;%.4f;%.4f;%.4f;%.4f;%.4f;%.4f;%.4f`,
		hopper.TokenId,
		hopper.Strength,
		hopper.Agility,
		hopper.Vitality,
		hopper.Intelligence,
		hopper.Fertility,
		hopper.Level,
		formatter.GetLevelCosts(hopper),
		formatter.GetMarketPrice(hopper),
		hopper.Activity,
		hopper.Image,
		hopper.RatingPond,
		hopper.RatingStream,
		hopper.RatingSwamp,
		hopper.RatingRiver,
		hopper.RatingForest,
		hopper.RatingGreatLake,
		hopper.BaseFlyPond,
		hopper.BaseFlyStream,
		hopper.BaseFlySwamp,
		hopper.BaseFlyRiver,
		hopper.BaseFlyForest,
		hopper.BaseFlyGreatLake,
	)
}

// Format hoppers to CSV. Header included
func (formatter *HopperFormatter) FormatAllCsv(hoppers []models.HopperDocument) []string {
	rows := make([]string, len(hoppers)+1)
	rows[0] = "tokenId;strength;agility;vitality;intelligence;fertility;level;levelCosts;price;activity;image;ratingPong;ratingStream;ratingSwap;ratingRiver;ratingForest;ratingGreatLake;baseFlyPond;baseFlyStream;baseFlySwamp;baseFlyRiver;baseFlyForest;baseFlyGreatLake"

	for i, hopper := range hoppers {
		rows[i+1] = formatter.FormatCsv(hopper)
	}

	return rows
}
