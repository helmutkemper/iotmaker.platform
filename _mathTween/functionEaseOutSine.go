package _mathTween

import "math"

// en: sinusoidal easing out - decelerating to zero velocity
var KEaseOutSine = func(interactionCurrent, interactionTotal, startValue, delta float64) float64 {
	return delta*math.Sin(interactionCurrent/interactionTotal*(math.Pi/2)) + startValue
}
