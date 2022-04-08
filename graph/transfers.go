package graph

import (
	"context"
	"strings"

	"github.com/machinebox/graphql"
	"github.com/steschwa/hopper-analytics-api/models"
)

const (
	TRANSFERS_GRAPH_URL = "https://api.thegraph.com/subgraphs/name/steschwa/hoppers-fly"
)

type (
	TransfersGraphClient struct {
		Graph *graphql.Client
	}
)

func NewTransfersGraphClient() *TransfersGraphClient {
	return &TransfersGraphClient{
		Graph: graphql.NewClient(TRANSFERS_GRAPH_URL),
	}
}

// ----------------------------------------
// Queries
// ----------------------------------------

const TRANSFERS_FROM_USER_QUERY = `
query($user: Bytes!) {
	transfers(
		where: { from: $user },
		orderBy: timestamp,
		orderDirection: desc
	) {
		from
		to
		methodId
		contract
		amount
		timestamp
		transaction
	}
}`

const TRANSFERS_TO_USER_QUERY = `
query($user: Bytes!) {
	transfers(
		where: { to: $user },
		orderBy: timestamp,
		orderDirection: desc
	) {
		from
		to
		methodId
		contract
		amount
		timestamp
		transaction
	}
}`

// ----------------------------------------
// Graph responses
// ----------------------------------------

type (
	TransferGraph struct {
		From        string `json:"from"`
		To          string `json:"to"`
		MethodId    string `json:"methodId"`
		Contract    string `json:"contract"`
		Amount      string `json:"amount"`
		Timestamp   string `json:"timestamp"`
		Transaction string `json:"transaction"`
	}

	TransfersResponse struct {
		Transfers []TransferGraph `json:"transfers"`
	}
)

// ----------------------------------------
// Graph response converters
// ----------------------------------------

func parseTransfer(transferGraph TransferGraph) models.Transfer {
	return models.Transfer{
		From:        transferGraph.From,
		To:          transferGraph.To,
		MethodId:    transferGraph.MethodId,
		Contract:    transferGraph.Contract,
		Amount:      ParseBigInt(transferGraph.Amount),
		Timestamp:   ParseUnixTime(transferGraph.Timestamp),
		Transaction: transferGraph.Transaction,
	}
}

// ----------------------------------------
// Query functions
// ----------------------------------------

func (client *TransfersGraphClient) FetchTransfersFromUser(user string) ([]models.Transfer, error) {
	req := graphql.NewRequest(TRANSFERS_FROM_USER_QUERY)
	req.Var("user", strings.ToLower(user))

	res := &TransfersResponse{}
	if err := client.Graph.Run(context.Background(), req, res); err != nil {
		return []models.Transfer{}, err
	}

	transfers := make([]models.Transfer, len(res.Transfers))
	for i, transfer := range res.Transfers {
		transfers[i] = parseTransfer(transfer)
	}

	return transfers, nil
}

func (client *TransfersGraphClient) FetchTransfersToUser(user string) ([]models.Transfer, error) {
	req := graphql.NewRequest(TRANSFERS_TO_USER_QUERY)
	req.Var("user", strings.ToLower(user))

	res := &TransfersResponse{}
	if err := client.Graph.Run(context.Background(), req, res); err != nil {
		return []models.Transfer{}, err
	}

	transfers := make([]models.Transfer, len(res.Transfers))
	for i, transfer := range res.Transfers {
		transfers[i] = parseTransfer(transfer)
	}

	return transfers, nil
}
