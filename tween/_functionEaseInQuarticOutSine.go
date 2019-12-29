package tween

import "math"

// en: quartic easing in/out - acceleration until halfway, then deceleration
var KEaseInQuarticOutSine = func(interactionCurrent, interactionTotal, startValue, delta float64) float64 {
	interactionCurrent = interactionCurrent / interactionTotal * 2
	if interactionCurrent < 1 {
		return delta/2*interactionCurrent*interactionCurrent*interactionCurrent*interactionCurrent + startValue
	}
	return -1*delta/2*(math.Cos(math.Pi*interactionCurrent/interactionTotal)-1) + startValue
}
