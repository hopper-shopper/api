package helpers

import (
	"github.com/steschwa/hopper-analytics-collector/constants"
	"github.com/steschwa/hopper-analytics-collector/helpers"
	"github.com/steschwa/hopper-analytics-collector/models"
)

func getAttributes(adventure constants.Adventure, hopper models.HopperDocument) []int {
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

func CalculateBaseShare(adventure constants.Adventure, hopper models.HopperDocument) int {
	if !helpers.CanEnter(adventure, helpers.HopperDocumentToHopper(hopper)) {
		return 0
	}

	attributes := getAttributes(adventure, hopper)

	attributesTotal := 1
	for _, attribute := range attributes {
		attributesTotal *= attribute
	}

	return attributesTotal * hopper.Level
}
