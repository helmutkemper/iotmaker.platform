package tween

import "math"

// en: exponential easing in/out - accelerating until halfway, then decelerating
var KEaseInExponentialOutCircular = func(interactionCurrent, interactionTotal, startValue, delta float64) float64 {
	interactionCurrent /= interactionTotal / 2
	if interactionCurrent < 1 {
		return delta/2*math.Pow(2, 10*(interactionCurrent-1)) + startValue
	}
	interactionCurrent -= 2
	return delta/2*(math.Sqrt(1-interactionCurrent*interactionCurrent)+1) + startValue
}
