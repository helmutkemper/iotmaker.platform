package tween

import "math"

// en: sinusoidal easing in/out - accelerating until halfway, then decelerating
var KEaseInOutSine = func(interactionCurrent, interactionTotal, startValue, delta float64) float64 {
	return -1*delta/2*(math.Cos(math.Pi*interactionCurrent/interactionTotal)-1) + startValue
}
