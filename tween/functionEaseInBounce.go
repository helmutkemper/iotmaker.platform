package tween

import "math"

var KEaseInBounce = func(interactionCurrent, interactionTotal, startValue, endValue, delta float64) float64 {
	value := interactionCurrent / interactionTotal
	formula := 0.0
	value = 1.0 - value

	if value < 0.36363636363 {
		formula = 7.5625 * math.Pow(value, 2.0)
	} else if value < 0.72727272727 {
		value -= 0.54545454545
		formula = 7.5625*value*value + 0.75
	} else if value < 0.90909090909 {
		value -= 0.81818181818
		formula = 7.5625*value*value + 0.9375
	} else {
		value -= 0.95454545454
		formula = 7.5625*value*value + 0.984375
	}

	if startValue < endValue {
		return endValue - formula*delta + startValue
	} else {
		return endValue - formula*delta
	}
}
