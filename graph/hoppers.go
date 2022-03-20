package graph

import (
	"context"

	"github.com/machinebox/graphql"
	"github.com/steschwa/hopper-analytics-api/constants"
	"github.com/steschwa/hopper-analytics-api/models"
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
		image
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
		Image        string         `json:"image"`
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

func parseHopper(hopperGraph HopperGraph) models.Hopper {
	listings := []models.Listing{}

	for _, listing := range hopperGraph.Listings {
		listings = append(listings, parseListing(listing))
	}

	return models.Hopper{
		TokenId:      hopperGraph.TokenId,
		Strength:     ParseInt(hopperGraph.Strength),
		Agility:      ParseInt(hopperGraph.Agility),
		Vitality:     ParseInt(hopperGraph.Vitality),
		Intelligence: ParseInt(hopperGraph.Intelligence),
		Market:       hopperGraph.Market,
		Level:        ParseInt(hopperGraph.Level),
		Adventure:    hopperGraph.Adventure,
		Image:        hopperGraph.Image,
		Listings:     listings,
	}
}

func parseListing(listingGraph ListingGraph) models.Listing {
	return models.Listing{
		Enabled: listingGraph.Enabled,
		Sold:    listingGraph.Sold,
		Price:   ParseBigFloat(listingGraph.Price),
	}
}

// ----------------------------------------
// Query functions
// ----------------------------------------

func (client *HoppersGraphClient) FetchAllHoppers() ([]models.Hopper, error) {
	hoppers := make([]models.Hopper, 0)

	for i := 0; i <= constants.HOPPERS_TOTAL_SUPPLY; i += 1000 {
		req := graphql.NewRequest(GET_HOPPERS_QUERY)
		req.Var("skip", i)

		res := &HoppersResponse{}
		if err := client.Graph.Run(context.Background(), req, res); err != nil {
			return []models.Hopper{}, err
		}

		for _, hopper := range res.HopperNFTs {
			hoppers = append(hoppers, parseHopper(hopper))
		}
	}

	return hoppers, nil
}
