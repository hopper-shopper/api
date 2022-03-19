package graph

import (
	"context"
	"math/big"

	"github.com/machinebox/graphql"
	"github.com/steschwa/hopper-analytics-api/constants"
)

type (
	HoppersGraphClient struct {
		Graph *graphql.Client
	}
)

func NewHoppersGraphClient() *HoppersGraphClient {
	return &HoppersGraphClient{
		Graph: graphql.NewClient(constants.HOPPERS_GRAPH_URL),
	}
}

// ----------------------------------------
// Models
// ----------------------------------------
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
		Listings     []Listing
	}
	Listing struct {
		Enabled bool
		Sold    bool
		Price   *big.Float
	}
)

// ----------------------------------------
// Queries
// ----------------------------------------
const GET_HOPPERS_QUERY = `
query($skip: Int!) {
	hopperNFTs(
		first: 1000,
		orderBy: tokenId,
		orderDirection: asc,
		skip: $skip
	) {
		tokenId
		strength
		agility
		vitality
		intelligence
		market
		level
		adventure
		listings {
			enabled
			price
			sold
		}
	}
}`

// ----------------------------------------
// Graph responses
// ----------------------------------------
type (
	HopperGraph struct {
		TokenId      string         `json:"tokenId"`
		Strength     string         `json:"strength"`
		Agility      string         `json:"agility"`
		Vitality     string         `json:"vitality"`
		Intelligence string         `json:"intelligence"`
		Market       bool           `json:"market"`
		Level        string         `json:"level"`
		Adventure    bool           `json:"adventure"`
		Listings     []ListingGraph `json:"listings"`
	}
	ListingGraph struct {
		Enabled bool   `json:"enabled"`
		Sold    bool   `json:"sold"`
		Price   string `json:"price"`
	}
	HoppersResponse struct {
		HopperNFTs []HopperGraph `json:"hopperNFTs"`
	}
)

// ----------------------------------------
// Graph response converters
// ----------------------------------------
func parseHopper(hopperGraph HopperGraph) Hopper {
	listings := []Listing{}

	for _, listing := range hopperGraph.Listings {
		listings = append(listings, parseListing(listing))
	}

	return Hopper{
		TokenId:      hopperGraph.TokenId,
		Strength:     ParseInt(hopperGraph.Strength),
		Agility:      ParseInt(hopperGraph.Agility),
		Vitality:     ParseInt(hopperGraph.Vitality),
		Intelligence: ParseInt(hopperGraph.Intelligence),
		Market:       hopperGraph.Market,
		Level:        ParseInt(hopperGraph.Level),
		Adventure:    hopperGraph.Adventure,
		Listings:     listings,
	}
}

func parseListing(listingGraph ListingGraph) Listing {
	return Listing{
		Enabled: listingGraph.Enabled,
		Sold:    listingGraph.Sold,
		Price:   ParseBigFloat(listingGraph.Price),
	}
}

// ----------------------------------------
// Query functions
// ----------------------------------------
func (client *HoppersGraphClient) FetchAllHoppers() ([]Hopper, error) {
	hoppers := make([]Hopper, 0)

	for i := 0; i <= constants.HOPPERS_TOTAL_SUPPLY; i += 1000 {
		req := graphql.NewRequest(GET_HOPPERS_QUERY)
		req.Var("skip", i)

		res := &HoppersResponse{}
		if err := client.Graph.Run(context.Background(), req, res); err != nil {
			return []Hopper{}, err
		}

		for _, hopper := range res.HopperNFTs {
			hoppers = append(hoppers, parseHopper(hopper))
		}
	}

	return hoppers, nil
}
