package tween

import "math"

// en: exponential easing in/out - accelerating until halfway, then decelerating
var KEaseInExponentialOutLinear = func(interactionCurrent, interactionTotal, startValue, endValue, delta float64) float64 {
	interactionCurrentToCalc := interactionCurrent / interactionTotal * 2
	if interactionCurrentToCalc < 1 {
		return delta/2*math.Pow(2, 10*(interactionCurrentToCalc-1)) + startValue
	}
	return delta*interactionCurrent/interactionTotal + startValue
}
