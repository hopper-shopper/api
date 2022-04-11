package utils

import "math"

func Clamp(min float64, max float64, value float64) float64 {
	return math.Min(max, math.Max(min, value))
}
