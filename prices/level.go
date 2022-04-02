package prices

import "math"

func CalculateLevelCost(forLevel uint8) float64 {
	if forLevel <= 1 || forLevel > 100 {
		return 0
	}

	if forLevel < 21 {
		return float64(forLevel / 2)
	}

	return math.Pow(float64(forLevel), 1.43522) / 7.5
}

func CalculateCumulatedLevelCost(forLevel uint8) float64 {
	totalCosts := 0.0

	for i := uint8(2); i <= forLevel; i++ {
		totalCosts += CalculateLevelCost(i)
	}

	return totalCosts
}
