package controllers

import (
	"math"
	"sort"
	"time"

	"github.com/akyoto/cache"
	"github.com/gofiber/fiber/v2"
	"github.com/steschwa/hopper-analytics-api/constants"
	"github.com/steschwa/hopper-analytics-api/filter"
	"github.com/steschwa/hopper-analytics-api/graph"
	"github.com/steschwa/hopper-analytics-api/helpers"
	"github.com/steschwa/hopper-analytics-api/models"
	sortHoppers "github.com/steschwa/hopper-analytics-api/sort"
)

func NewMarketsByAdventureController() RouteHandler {
	return func(ctx *fiber.Ctx) error {
		adventure := constants.AdventureFromString(ctx.Params("adventure"))
		sortBy := models.HopperOnMarketSortFromString(ctx.Query("sortBy"))

		hoppers := loadHoppers()
		hoppersWithPermit := getHopperAdventurePermitFilter(adventure).Get(hoppers)

		marketFilter := filter.NewHoppersFilter().Where(filter.KeyMarket, filter.NewBoolFilter(true))
		hoppersWithPermitOnMarket := marketFilter.Get(hoppersWithPermit)

		markets := make([]models.HopperOnMarket, len(hoppersWithPermitOnMarket))
		for i, hopper := range hoppersWithPermitOnMarket {
			markets[i] = getHopperMarketByAdventure(adventure, hopper)
		}

		sharePermit := float64(len(hoppersWithPermit)) / float64(constants.HOPPERS_TOTAL_SUPPLY)
		sharePermitOnMarket := float64(len(hoppersWithPermitOnMarket)) / float64(constants.HOPPERS_TOTAL_SUPPLY)
		shareMarketToPermit := sharePermitOnMarket / sharePermit

		sortedHoppers := sortHoppersOnMarkets(sortBy, markets)

		response := fiber.Map{
			"adventure": adventure.String(),
			"total":     constants.HOPPERS_TOTAL_SUPPLY,
			"shares": fiber.Map{
				"withPermit":         sharePermit,
				"withPermitOnMarket": sharePermitOnMarket,
				"marketToPermit":     shareMarketToPermit,
			},
			"data": formatMarkets(sortedHoppers),
		}

		return ctx.JSON(response)
	}
}

var hoppersCache = cache.New(time.Minute)

func loadHoppers() []models.Hopper {
	data, found := hoppersCache.Get("hoppers")
	if found {
		if data, ok := data.([]models.Hopper); ok {
			return data
		}
	}

	client := graph.NewHoppersGraphClient()

	hoppers, err := client.FetchAllHoppers()
	if err != nil {
		return []models.Hopper{}
	}

	hoppersCache.Set("hoppers", hoppers, time.Minute*5)

	return hoppers
}

// Construct a query that filters all hoppers which are allowed
// to enter the given adventure
func getHopperAdventurePermitFilter(adventure constants.Adventure) filter.Filter {
	if adventure == constants.AdventurePond || adventure == constants.AdventureStream || adventure == constants.AdventureSwamp {
		return filter.NewHoppersFilter()
	} else if adventure == constants.AdventureRiver {
		return filter.NewHoppersFilter().
			Where(filter.KeyStrength, filter.NewIntFilter(5, filter.OpGe)).
			Where(filter.KeyIntelligence, filter.NewIntFilter(5, filter.OpGe))
	} else if adventure == constants.AdventureForest {
		return filter.NewHoppersFilter().
			Where(filter.KeyAgility, filter.NewIntFilter(5, filter.OpGe)).
			Where(filter.KeyIntelligence, filter.NewIntFilter(5, filter.OpGe)).
			Where(filter.KeyVitality, filter.NewIntFilter(5, filter.OpGe))
	} else if adventure == constants.AdventureGreatLake {
		return filter.NewHoppersFilter().
			Where(filter.KeyStrength, filter.NewIntFilter(5, filter.OpGe)).
			Where(filter.KeyAgility, filter.NewIntFilter(5, filter.OpGe)).
			Where(filter.KeyIntelligence, filter.NewIntFilter(5, filter.OpGe)).
			Where(filter.KeyVitality, filter.NewIntFilter(5, filter.OpGe))
	}

	return filter.NewHoppersFilter()
}

func getHopperMarketByAdventure(adventure constants.Adventure, hopper models.Hopper) models.HopperOnMarket {
	price := 0.0

	for _, listing := range hopper.Listings {
		if listing.Enabled && !listing.Sold {
			val, _ := listing.Price.Float64()
			price = val * math.Pow(10, -18)
			break
		}
	}

	rating := getHopperRatingByAdventure(adventure, hopper)

	return models.HopperOnMarket{
		Hopper: hopper,
		Price:  price,
		Rating: rating,
	}
}

func getHopperRatingByAdventure(adventure constants.Adventure, hopper models.Hopper) int {
	if !helpers.CanEnter(adventure, hopper) {
		return -1
	}

	switch adventure {
	case constants.AdventurePond:
		return hopper.Strength
	case constants.AdventureStream:
		return hopper.Agility
	case constants.AdventureSwamp:
		return hopper.Vitality
	case constants.AdventureRiver:
		return hopper.Strength * hopper.Intelligence
	case constants.AdventureForest:
		return hopper.Agility * hopper.Vitality * hopper.Intelligence
	case constants.AdventureGreatLake:
		return hopper.Strength * hopper.Agility * hopper.Vitality * hopper.Intelligence
	default:
		return 0
	}
}

func sortHoppersOnMarkets(sortBy models.HopperOnMarketSort, markets []models.HopperOnMarket) []models.HopperOnMarket {
	marketsCopy := make([]models.HopperOnMarket, len(markets))
	copy(marketsCopy, markets)

	switch sortBy {
	case models.HopperOnMarketSortByTokenId:
		sort.Sort(sortHoppers.SortHopperOnMarketByTokenId(marketsCopy))
	case models.HopperOnMarketSortByPrice:
		sort.Sort(sortHoppers.SortHopperOnMarketByPrice(marketsCopy))
	}

	return marketsCopy
}

func formatMarkets(markets []models.HopperOnMarket) []fiber.Map {
	response := make([]fiber.Map, len(markets))
	for i, market := range markets {
		response[i] = fiber.Map{
			"tokenId":      market.TokenId,
			"strength":     market.Strength,
			"agility":      market.Agility,
			"vitality":     market.Vitality,
			"intelligence": market.Intelligence,
			"level":        market.Level,
			"price":        market.Price,
			"rating":       market.Rating,
		}
	}

	return response
}
