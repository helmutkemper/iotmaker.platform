package tween

import "math"

// en: cubic easing in/out - acceleration until halfway, then deceleration
var KEaseInSineOutCubic = func(interactionCurrent, interactionTotal, startValue, delta float64) float64 {
	interactionCurrentToCalc := interactionCurrent / interactionTotal * 2
	if interactionCurrentToCalc < 1 {
		return -1*delta/2*(math.Cos(math.Pi*interactionCurrent/interactionTotal)-1) + startValue
	}
	interactionCurrentToCalc -= 2
	return delta/2*(interactionCurrentToCalc*interactionCurrentToCalc*interactionCurrentToCalc+2) + startValue
}
