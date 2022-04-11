package tween

import "math"

var KEaseInElastic = func(interactionCurrent, interactionTotal, startValue, endValue, delta float64) float64 {

	value := interactionCurrent / interactionTotal
	if value == 0 {
		return startValue
	}

	if value == 1.0 {
		return 1.0*delta + startValue
	}

	value -= 1

	return -(math.Pow(2, 10*value)*math.Sin((value-0.075)*20.9435102))*delta + startValue
}
