package tween

import "math"

var KEaseInBounce = func(interactionCurrent, interactionTotal, currentPercentage, startValue, endValue, delta float64) float64 {
	formula := 0.0
	currentPercentage = 1.0 - currentPercentage

	if currentPercentage < 0.36363636363 {
		formula = 7.5625 * math.Pow(currentPercentage, 2.0)
	} else if currentPercentage < 0.72727272727 {
		currentPercentage -= 0.54545454545
		formula = 7.5625*math.Pow(currentPercentage, 2.0) + 0.75
	} else if currentPercentage < 0.90909090909 {
		currentPercentage -= 0.81818181818
		formula = 7.5625*math.Pow(currentPercentage, 2.0) + 0.9375
	} else {
		currentPercentage -= 0.95454545454
		formula = 7.5625*math.Pow(currentPercentage, 2.0) + 0.984375
	}

	if startValue < endValue {
		return endValue - formula*delta + startValue
	} else {
		return endValue - formula*delta
	}
}
