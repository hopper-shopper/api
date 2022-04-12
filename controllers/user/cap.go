package user

import (
	"math/big"

	"github.com/steschwa/hopper-analytics-collector/constants"
	"github.com/steschwa/hopper-analytics-collector/contracts"
)

type (
	UserFlyGenerationCalculator struct {
		OnChainClient *contracts.OnChainClient
	}

	UserAdventureFlyGeneration struct {
		Cap     *big.Int
		Current *big.Int
	}
)

func NewUserFlyGenerationCalculator(onChainClient *contracts.OnChainClient) *UserFlyGenerationCalculator {
	return &UserFlyGenerationCalculator{
		OnChainClient: onChainClient,
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

	return UserAdventureFlyGeneration{
		Cap:     userCap,
		Current: sumGenerated,
	}, nil
}
