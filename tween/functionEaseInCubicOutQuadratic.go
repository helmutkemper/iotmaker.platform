package tween

// en: cubic easing in/out - acceleration until halfway, then deceleration
var KEaseInCubicOutQuadratic = func(interactionCurrent, interactionTotal, startValue, endValue, delta float64) float64 {
	interactionCurrent = interactionCurrent / interactionTotal * 2
	if interactionCurrent < 1 {
		return delta/2*interactionCurrent*interactionCurrent*interactionCurrent + startValue
	}
	interactionCurrent--
	return -1*delta/2*(interactionCurrent*(interactionCurrent-2)-1) + startValue
}
