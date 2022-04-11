package tween

import "math"

// en: Quadratic easing in/out - acceleration until halfway, then deceleration
var KEaseInSineOutQuadratic = func(interactionCurrent, interactionTotal, startValue, endValue, delta float64) float64 {
	interactionCurrentToCalc := interactionCurrent / interactionTotal * 2
	if interactionCurrentToCalc < 1 {
		return -1*delta/2*(math.Cos(math.Pi*interactionCurrent/interactionTotal)-1) + startValue
	}
	interactionCurrentToCalc--
	return -1*delta/2*(interactionCurrentToCalc*(interactionCurrentToCalc-2)-1) + startValue
}
