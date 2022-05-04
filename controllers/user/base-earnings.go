package user

import (
	"context"

	"github.com/shopspring/decimal"
	"github.com/steschwa/hopper-analytics-collector/constants"
	"github.com/steschwa/hopper-analytics-collector/contracts"
	"github.com/steschwa/hopper-analytics-collector/helpers"
	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	UserBaseFlyCalculator struct {
		OnChainClient *contracts.OnChainClient
		DbClient      *db.MongoDbClient
	}
)

func NewUserBaseFlyCalculator(onChainClient *contracts.OnChainClient, dbClient *db.MongoDbClient) *UserBaseFlyCalculator {
	return &UserBaseFlyCalculator{
		OnChainClient: onChainClient,
		DbClient:      dbClient,
	}
}

func (calculator *UserBaseFlyCalculator) CalculateBaseFly(adventure constants.Adventure, user string) (float64, error) {
	totalBaseShare, err := calculator.loadLatestTotalBaseShareByAdventure(adventure)
	if err != nil {
		return 0, err
	}

	userBaseShare, err := calculator.loadUserBaseShareByAdventure(adventure, user)
	if err != nil {
		return 0, err
	}

	share, _ := userBaseShare.Div(totalBaseShare).Float64()

	return helpers.GetAdventureRewards(adventure) * share, nil
}

func (calculator *UserBaseFlyCalculator) loadLatestTotalBaseShareByAdventure(adventure constants.Adventure) (decimal.Decimal, error) {
	baseSharesCollection := &db.BaseSharesCollection{
		Client: calculator.DbClient,
	}
	result := baseSharesCollection.GetCollection().FindOne(
		context.Background(),
		getBaseSharesFilter(adventure),
		&options.FindOneOptions{
			Sort: bson.D{{
				Key:   "updated",
				Value: -1,
			}},
		},
	)

	if err := result.Err(); err != nil {
		return decimal.NewFromInt(0), err
	}

	baseShareDocument := models.BaseSharesDocument{}
	if err := result.Decode(&baseShareDocument); err != nil {
		return decimal.NewFromInt(0), err
	}

	return decimal.NewFromString(baseShareDocument.TotalBaseShares.Int().String())
}

func (calculator *UserBaseFlyCalculator) loadUserBaseShareByAdventure(adventure constants.Adventure, user string) (decimal.Decimal, error) {
	userBaseShare, err := calculator.OnChainClient.GetUserBaseSharesByAdventure(adventure, user)
	if err != nil {
		return decimal.NewFromInt(0), err
	}

	return decimal.NewFromString(userBaseShare.String())
}
