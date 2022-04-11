package tween

import "math"

// en: exponential easing in/out - accelerating until halfway, then decelerating
var KEaseInExponentialOutQuadratic = func(interactionCurrent, interactionTotal, startValue, endValue, delta float64) float64 {
	interactionCurrent /= interactionTotal / 2
	if interactionCurrent < 1 {
		return delta/2*math.Pow(2, 10*(interactionCurrent-1)) + startValue
	}
	interactionCurrent--
	return -1*delta/2*(interactionCurrent*(interactionCurrent-2)-1) + startValue
}
