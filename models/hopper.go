package models

import (
	"math/big"
	"strings"

	"github.com/steschwa/hopper-analytics-collector/models"
)

type (
	Listing struct {
		Enabled bool
		Sold    bool
		Price   *big.Float
	}

	HopperOnMarket struct {
		models.HopperDocument
		Price  float64
		Rating int
	}
)

type (
	HopperOnMarketSort int
)

const (
	HopperOnMarketSortByTokenId HopperOnMarketSort = iota
	HopperOnMarketSortByPrice
)

func HopperOnMarketSortFromString(sort string) HopperOnMarketSort {
	lowerCased := strings.ToLower(sort)

	switch lowerCased {
	case "tokenId":
		return HopperOnMarketSortByTokenId
	case "price":
		return HopperOnMarketSortByPrice
	default:
		return HopperOnMarketSortByTokenId
	}
}
