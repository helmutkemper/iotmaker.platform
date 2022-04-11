package tween

import (
	"math"
)

// en: exponential easing in - accelerating from zero velocity
var KEaseInExponential = func(interactionCurrent, interactionTotal, startValue, endValue, delta float64) float64 {
	return delta*math.Pow(2, 10*(interactionCurrent/interactionTotal-1)) + startValue
}
