package sort

import (
	"strconv"

	"github.com/steschwa/hopper-analytics-api/models"
)

type (
	SortHopperOnMarketByTokenId []models.HopperOnMarket
	SortHopperOnMarketByPrice   []models.HopperOnMarket
)

// ----------------------------------------
// SortHopperByTokenId
// ----------------------------------------

func (hoppers SortHopperOnMarketByTokenId) Len() int {
	return len(hoppers)
}

func (hoppers SortHopperOnMarketByTokenId) Swap(i, j int) {
	hoppers[i], hoppers[j] = hoppers[j], hoppers[i]
}

func (hoppers SortHopperOnMarketByTokenId) Less(i, j int) bool {
	hopperId1, err1 := strconv.ParseInt(hoppers[i].TokenId, 10, 64)
	hopperId2, err2 := strconv.ParseInt(hoppers[j].TokenId, 10, 64)

	if err1 != nil {
		return true
	}
	if err2 != nil {
		return false
	}

	return hopperId1 < hopperId2
}

// ----------------------------------------
// SortHopperOnMarketByPrice
// ----------------------------------------

func (hoppers SortHopperOnMarketByPrice) Len() int {
	return len(hoppers)
}

func (hoppers SortHopperOnMarketByPrice) Swap(i, j int) {
	hoppers[i], hoppers[j] = hoppers[j], hoppers[i]
}

func (hoppers SortHopperOnMarketByPrice) Less(i, j int) bool {
	return hoppers[i].Price < hoppers[j].Price
}
