package tween

// en: quintic easing in/out - acceleration until halfway, then deceleration
var KEaseInQuarticOutQuintic = func(interactionCurrent, interactionTotal, startValue, endValue, delta float64) float64 {
	interactionCurrent = interactionCurrent / interactionTotal * 2
	if interactionCurrent < 1 {
		return delta/2*interactionCurrent*interactionCurrent*interactionCurrent*interactionCurrent + startValue
	}
	interactionCurrent -= 2
	return delta/2*(interactionCurrent*interactionCurrent*interactionCurrent*interactionCurrent*interactionCurrent+2) + startValue
}
