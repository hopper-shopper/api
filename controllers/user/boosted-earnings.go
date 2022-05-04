package user

import (
	"context"
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/shopspring/decimal"
	"github.com/steschwa/hopper-analytics-collector/constants"
	"github.com/steschwa/hopper-analytics-collector/contracts"
	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	UserBoostedFlyCalculator struct {
		OnChainClient *contracts.OnChainClient
		DbClient      *db.MongoDbClient
	}
)

func NewUserBoostedFlyCalculator(onChainClient *contracts.OnChainClient, dbClient *db.MongoDbClient) *UserBoostedFlyCalculator {
	return &UserBoostedFlyCalculator{
		OnChainClient: onChainClient,
		DbClient:      dbClient,
	}
}

func (calculator *UserBoostedFlyCalculator) CalculateBoostedFly(adventure constants.Adventure, user string) (float64, error) {
	boostedReward, err := calculator.calculateBoostedReward(adventure)
	if err != nil {
		return 0, err
	}

	userVeSharePercent, err := calculator.calculateVeSharePercent(adventure, user)
	if err != nil {
		return 0, err
	}

	return boostedReward * userVeSharePercent, nil
}

func (calculator *UserBoostedFlyCalculator) calculateBoostedReward(adventure constants.Adventure) (float64, error) {
	votesInAdventure, err := calculator.loadLatestTotalVotesByAdventure(adventure)
	if err != nil {
		return 0, err
	}
	votesInAdventureDec, err := decimal.NewFromString(votesInAdventure.String())
	if err != nil {
		return 0, err
	}

	totalVotes, err := calculator.calculateTotalVotes()
	if err != nil {
		return 0, err
	}

	boostedReward, _ := votesInAdventureDec.Div(totalVotes).Float64()
	return boostedReward * constants.BONUS_FLY_EMISSIONS, nil
}

func (calculator *UserBoostedFlyCalculator) calculateTotalVotes() (decimal.Decimal, error) {
	adventures := []constants.Adventure{
		constants.AdventurePond,
		constants.AdventureStream,
		constants.AdventureSwamp,
		constants.AdventureRiver,
		constants.AdventureForest,
		constants.AdventureGreatLake,
	}

	totalVotes := decimal.NewFromInt(0)
	for _, adventure := range adventures {
		votesInAdventure, err := calculator.loadLatestTotalVotesByAdventure(adventure)
		if err != nil {
			log.Println(err)
			sentry.CaptureException(err)
			continue
		}

		totalVotes = totalVotes.Add(votesInAdventure)
	}

	return totalVotes, nil
}

func (calculator *UserBoostedFlyCalculator) calculateVeSharePercent(adventure constants.Adventure, user string) (float64, error) {
	userVeShareInAdventure, err := calculator.OnChainClient.GetUserVeShareByAdventure(adventure, user)
	if err != nil {
		return 0, err
	}
	userVeShareInAdventureDec, err := decimal.NewFromString(userVeShareInAdventure.String())
	if err != nil {
		return 0, err
	}

	totalVeShareInAdventure, err := calculator.loadLatestTotalVeShareByAdventure(adventure)
	if err != nil {
		return 0, err
	}

	veSharePercent, _ := userVeShareInAdventureDec.Div(totalVeShareInAdventure).Float64()
	return veSharePercent, nil
}

func (calculator *UserBoostedFlyCalculator) loadLatestVoteDocument(adventure constants.Adventure) (models.VoteDocument, error) {
	votesCollection := &db.VotesCollection{
		Client: calculator.DbClient,
	}
	result := votesCollection.GetCollection().FindOne(
		context.Background(),
		getVotesFilter(adventure),
		&options.FindOneOptions{
			Sort: bson.D{{
				Key:   "updated",
				Value: -1,
			}},
		},
	)

	if err := result.Err(); err != nil {
		return models.VoteDocument{}, err
	}

	voteDocument := models.VoteDocument{}
	if err := result.Decode(&voteDocument); err != nil {
		return models.VoteDocument{}, err
	}

	return voteDocument, nil
}

func (calculator *UserBoostedFlyCalculator) loadLatestTotalVotesByAdventure(adventure constants.Adventure) (decimal.Decimal, error) {
	voteDocument, err := calculator.loadLatestVoteDocument(adventure)
	if err != nil {
		return decimal.NewFromInt(0), err
	}

	return decimal.NewFromString(voteDocument.Votes.Int().String())
}

func (calculator *UserBoostedFlyCalculator) loadLatestTotalVeShareByAdventure(adventure constants.Adventure) (decimal.Decimal, error) {
	voteDocument, err := calculator.loadLatestVoteDocument(adventure)
	if err != nil {
		return decimal.NewFromInt(0), err
	}

	return decimal.NewFromString(voteDocument.VeShare.Int().String())
}
