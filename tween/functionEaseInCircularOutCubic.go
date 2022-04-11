package tween

import "math"

// en: cubic easing in/out - acceleration until halfway, then deceleration
var KEaseInCircularOutCubic = func(interactionCurrent, interactionTotal, startValue, endValue, delta float64) float64 {
	interactionCurrent = interactionCurrent / interactionTotal * 2
	if interactionCurrent < 1 {
		return -1*delta/2*(math.Sqrt(1-interactionCurrent*interactionCurrent)-1) + startValue
	}
	interactionCurrent -= 2
	return delta/2*(interactionCurrent*interactionCurrent*interactionCurrent+2) + startValue
}
