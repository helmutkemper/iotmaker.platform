package tween

import "math"

// en: exponential easing in/out - accelerating until halfway, then decelerating
var KEaseInSineOutExponential = func(interactionCurrent, interactionTotal, startValue, endValue, delta float64) float64 {
	interactionCurrentToCalc := interactionCurrent / interactionTotal * 2
	if interactionCurrentToCalc < 1 {
		return -1*delta/2*(math.Cos(math.Pi*interactionCurrent/interactionTotal)-1) + startValue
	}
	interactionCurrentToCalc--
	return delta/2*(-1*math.Pow(2, -10*interactionCurrentToCalc)+2) + startValue
}
