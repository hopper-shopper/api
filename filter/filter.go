package filter

import "github.com/steschwa/hopper-analytics-api/models"

type (
	Key      int
	Operator string

	FilterFn interface {
		Match(value interface{}) bool
	}

	Query struct {
		Key Key
		Fn  FilterFn
	}

	Filter interface {
		Get(hoppers []models.Hopper) []models.Hopper
	}

	HoppersFilter struct {
		Queries []Query
	}
)

const (
	KeyStrength Key = iota
	KeyAgility
	KeyVitality
	KeyIntelligence
	KeyLevel
	KeyMarket
	KeyAdventure
)

func NewHoppersFilter() *HoppersFilter {
	return &HoppersFilter{
		Queries: make([]Query, 0),
	}
}

func (filter *HoppersFilter) Where(key Key, fn FilterFn) *HoppersFilter {
	query := Query{
		Key: key,
		Fn:  fn,
	}
	filter.Queries = append(filter.Queries, query)

	return filter
}

func (filter *HoppersFilter) Get(hoppers []models.Hopper) []models.Hopper {
	if len(filter.Queries) == 0 {
		return hoppers
	}

	result := make([]models.Hopper, 0)

	for _, hopper := range hoppers {
		if filter.Match(hopper) {
			result = append(result, hopper)
		}
	}

	return result
}

func (filter *HoppersFilter) Match(hopper models.Hopper) bool {
	for _, query := range filter.Queries {
		if !checkMatch(hopper, query) {
			return false
		}
	}

	return true
}

func checkMatch(hopper models.Hopper, query Query) bool {
	var filterVal interface{}

	switch query.Key {
	case KeyStrength:
		filterVal = hopper.Strength
	case KeyAgility:
		filterVal = hopper.Agility
	case KeyVitality:
		filterVal = hopper.Vitality
	case KeyIntelligence:
		filterVal = hopper.Intelligence
	case KeyLevel:
		filterVal = hopper.Level
	case KeyMarket:
		filterVal = hopper.Market
	case KeyAdventure:
		filterVal = hopper.Adventure
	}

	return query.Fn.Match(filterVal)
}
