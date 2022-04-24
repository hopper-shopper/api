package user

import (
	"math"
	"math/big"

	"github.com/steschwa/hopper-analytics-api/utils"
	"github.com/steschwa/hopper-analytics-collector/constants"
	"github.com/steschwa/hopper-analytics-collector/contracts"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	UserFlyGenerationCalculator struct {
		OnChainClient *contracts.OnChainClient
		MongoClient   *mongo.Client
	}

	UserAdventureFlyGeneration struct {
		Cap     float64
		Current float64
		Time    uint
	}
)

func NewUserFlyGenerationCalculator(onChainClient *contracts.OnChainClient, mongoClient *mongo.Client) *UserFlyGenerationCalculator {
	return &UserFlyGenerationCalculator{
		OnChainClient: onChainClient,
		MongoClient:   mongoClient,
	}
}

func (calculator *UserFlyGenerationCalculator) CalculateFlyGeneration(adventure constants.Adventure, user string) (UserAdventureFlyGeneration, error) {
	userCap, err := calculator.OnChainClient.GetUserMaxFlyGeneration(adventure, user)
	if err != nil {
		return UserAdventureFlyGeneration{}, err
	}

	userBaseGenerated, _, err := calculator.OnChainClient.GetUserGeneratedFly(adventure, user)
	if err != nil {
		return UserAdventureFlyGeneration{}, err
	}
	userBonusGenerated, _, err := calculator.OnChainClient.GetUserBoostedGeneratedFly(adventure, user)
	if err != nil {
		return UserAdventureFlyGeneration{}, err
	}

	sumGenerated := big.NewInt(0).Add(userBaseGenerated, userBonusGenerated)

	baseFlyCalculator := NewUserBaseFlyCalculator(calculator.OnChainClient, calculator.MongoClient)
	baseFly, err := baseFlyCalculator.CalculateBaseFly(adventure, user)
	if err != nil {
		return UserAdventureFlyGeneration{}, err
	}

	boostedFlyCalculator := NewUserBoostedFlyCalculator(calculator.OnChainClient, calculator.MongoClient)
	boostedFly, err := boostedFlyCalculator.CalculateBoostedFly(adventure, user)
	if err != nil {
		return UserAdventureFlyGeneration{}, err
	}

	flyPerDay := baseFly + boostedFly
	flyPerSecond := flyPerDay / 24 / 60 / 60

	capTotal, _ := utils.ToDecimal(userCap.String(), 30).Float64()
	capCurrent, _ := utils.ToDecimal(sumGenerated.String(), 18).Float64()

	clampedTotalCap := math.Max(0, capTotal)
	clampedCurrentCap := utils.Clamp(0, clampedTotalCap, capCurrent)

	capLeft := clampedTotalCap - clampedCurrentCap

	var timeBeforeCap uint = 0
	if flyPerSecond > 0 {
		timeBeforeCap = uint(capLeft / flyPerSecond)
	}

	return UserAdventureFlyGeneration{
		Cap:     clampedTotalCap,
		Current: clampedCurrentCap,
		Time:    timeBeforeCap,
	}, nil
}
