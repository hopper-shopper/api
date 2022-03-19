package helpers

import (
	"github.com/steschwa/hopper-analytics-api/constants"
	"github.com/steschwa/hopper-analytics-api/graph"
)

func CanEnter(adventure constants.Adventure, hopper graph.Hopper) bool {
	switch adventure {
	case constants.AdventurePond, constants.AdventureStream, constants.AdventureSwamp:
		return true
	case constants.AdventureRiver:
		return CanEnterRiver(hopper)
	case constants.AdventureForest:
		return CanEnterForest(hopper)
	case constants.AdventureGreatLake:
		return CanEnterGreatLake(hopper)
	default:
		return false
	}
}

func CanEnterRiver(hopper graph.Hopper) bool {
	return hopper.Strength >= 5 && hopper.Intelligence >= 5
}

func CanEnterForest(hopper graph.Hopper) bool {
	return hopper.Agility >= 5 && hopper.Vitality >= 5 && hopper.Intelligence >= 5
}

func CanEnterGreatLake(hopper graph.Hopper) bool {
	return hopper.Strength >= 5 && hopper.Agility >= 5 && hopper.Vitality >= 5 && hopper.Intelligence >= 5
}

func getAttributes(adventure constants.Adventure, hopper graph.Hopper) []int {
	switch adventure {
	case constants.AdventurePond:
		return []int{hopper.Strength}
	case constants.AdventureStream:
		return []int{hopper.Agility}
	case constants.AdventureSwamp:
		return []int{hopper.Vitality}
	case constants.AdventureRiver:
		return []int{hopper.Strength, hopper.Intelligence}
	case constants.AdventureForest:
		return []int{hopper.Agility, hopper.Vitality, hopper.Intelligence}
	case constants.AdventureGreatLake:
		return []int{hopper.Strength, hopper.Agility, hopper.Vitality, hopper.Intelligence}
	default:
		return []int{}
	}
}

func CalculateBaseShare(adventure constants.Adventure, hopper graph.Hopper) int {
	attributes := getAttributes(adventure, hopper)

	attributesTotal := 1
	for _, attribute := range attributes {
		attributesTotal *= attribute
	}

	return attributesTotal * hopper.Level
}
