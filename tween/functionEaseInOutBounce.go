package tween

import "math"

var KEaseInOutBounce = func(interactionCurrent, interactionTotal, startValue, endValue, delta float64) float64 {
	value := interactionCurrent / interactionTotal

	if value < 0.5 {
		return (1.0-bounceOut(1-2*value))/2*delta + startValue
	} else {
		return (1.0+bounceOut(2*value-1))/2*delta + startValue
	}
}

func bounceOut(value float64) float64 {
	if value < 0.36363636363 {
		return 7.5625 * math.Pow(value, 2.0)
	} else if value < 0.72727272727 {
		value -= 0.54545454545
		return 7.5625*value*value + 0.75
	} else if value < 0.90909090909 {
		value -= 0.81818181818
		return 7.5625*value*value + 0.9375
	} else {
		value -= 0.95454545454
		return 7.5625*value*value + 0.984375
	}
}
