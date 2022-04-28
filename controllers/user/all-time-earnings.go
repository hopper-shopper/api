package user

import (
	"fmt"
	"math/big"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/steschwa/hopper-analytics-collector/graph"
)

type (
	UserEarningsCalculator struct {
		Graph *graph.TransfersGraphClient
		Cache *cache.Cache
	}

	GroupedTransfers struct {
		Key       string
		Transfers []graph.Transfer
	}
)

func NewUserAllTimeEarningsCalculator() *UserEarningsCalculator {
	cache := cache.New(time.Minute, time.Second*30)

	return &UserEarningsCalculator{
		Graph: graph.NewTransfersGraphClient(),
		Cache: cache,
	}
}

func (calculator *UserEarningsCalculator) FetchTransfers(filter UserEarningsFilter) ([]graph.Transfer, error) {
	cacheKey := fmt.Sprintf("%s.%s.claimed-transfers", filter.Adventure.String(), filter.User)

	if v, ok := calculator.Cache.Get(cacheKey); ok {
		return v.([]graph.Transfer), nil
	}

	transfers, err := calculator.Graph.FetchClaimedTransfers(filter.Adventure, filter.User)
	if err != nil {
		return []graph.Transfer{}, err
	}

	calculator.Cache.Set(cacheKey, transfers, 0)
	return transfers, nil
}

func (calculator *UserEarningsCalculator) GetAllTimeEarnings(filter UserEarningsFilter) (*big.Int, error) {
	transfers, err := calculator.FetchTransfers(filter)
	if err != nil {
		return big.NewInt(0), err
	}

	total := big.NewInt(0)
	for _, transfer := range transfers {
		total = big.NewInt(0).Add(total, transfer.Amount)
	}

	return total, nil
}

func (calculator *UserEarningsCalculator) GetTransfersHistory(filter UserEarningsHistoryFilter) (map[string]*GroupedTransfers, error) {
	transfers, err := calculator.FetchTransfers(UserEarningsFilter{
		User:      filter.User,
		Adventure: filter.Adventure,
	})
	if err != nil {
		return map[string]*GroupedTransfers{}, err
	}

	groups := map[string]*GroupedTransfers{}
	for _, transfer := range transfers {
		groupKey := getTransferGroupKey(filter.GroupBy, transfer)
		if groupKey == "" {
			continue
		}

		group, found := groups[groupKey]
		if found {
			group.Transfers = append(group.Transfers, transfer)
		} else {
			groups[groupKey] = &GroupedTransfers{
				Key:       groupKey,
				Transfers: []graph.Transfer{transfer},
			}
		}
	}

	return groups, nil
}

func getTransferGroupKey(groupBy GroupTransfersBy, transfer graph.Transfer) string {
	switch groupBy {
	case GroupTransfersByDay:
		return transfer.Timestamp.Format("2006-01-02")
	case GroupTransfersByMonth:
		return transfer.Timestamp.Format("2006-01")
	}

	return ""
}
