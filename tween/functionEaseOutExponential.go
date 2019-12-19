package tween

import "math"

// en: exponential easing out - decelerating to zero velocity
var KEaseOutExponential = func(interactionCurrent, interactionTotal, startValue, delta float64) float64 {
	return delta*(-1*math.Pow(2, -10*interactionCurrent/interactionTotal)+1) + startValue
}
