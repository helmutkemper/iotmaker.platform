package tween

import "math"

// en: sinusoidal easing in - accelerating from zero velocity
var KEaseInSine = func(interactionCurrent, interactionTotal, startValue, endValue, delta float64) float64 {
	return -1*delta*math.Cos(interactionCurrent/interactionTotal*(math.Pi/2)) + delta + startValue
}
