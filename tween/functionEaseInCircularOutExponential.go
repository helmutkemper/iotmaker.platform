package tween

import "math"

// en: exponential easing in/out - accelerating until halfway, then decelerating
var KEaseInCircularOutExponential = func(interactionCurrent, interactionTotal, startValue, endValue, delta float64) float64 {
	interactionCurrent /= interactionTotal / 2
	if interactionCurrent < 1 {
		return -1*delta/2*(math.Sqrt(1-interactionCurrent*interactionCurrent)-1) + startValue
	}
	interactionCurrent--
	return delta/2*(-1*math.Pow(2, -10*interactionCurrent)+2) + startValue
}
