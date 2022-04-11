package tween

import "math"

// en: sinusoidal easing in/out - accelerating until halfway, then decelerating
var KEaseInQuinticOutSine = func(interactionCurrent, interactionTotal, startValue, endValue, delta float64) float64 {
	interactionCurrentToCalc := interactionCurrent / interactionTotal * 2
	if interactionCurrentToCalc < 1 {
		return delta/2*interactionCurrentToCalc*interactionCurrentToCalc*interactionCurrentToCalc*interactionCurrentToCalc*interactionCurrentToCalc + startValue
	}

	return -1*delta/2*(math.Cos(math.Pi*interactionCurrent/interactionTotal)-1) + startValue
}
