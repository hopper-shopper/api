package models

import (
	"math/big"
	"strings"
)

type (
	Hopper struct {
		TokenId      string
		Strength     int
		Agility      int
		Vitality     int
		Intelligence int
		Market       bool
		Level        int
		Adventure    bool
		Image        string
		Listings     []Listing
	}
	Listing struct {
		Enabled bool
		Sold    bool
		Price   *big.Float
	}

	HopperOnMarket struct {
		Hopper
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
