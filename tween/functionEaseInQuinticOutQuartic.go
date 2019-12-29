package tween

// en: quartic easing in/out - acceleration until halfway, then deceleration
var KEaseInQuinticOutQuartic = func(interactionCurrent, interactionTotal, startValue, delta float64) float64 {
	interactionCurrent = interactionCurrent / interactionTotal * 2
	if interactionCurrent < 1 {
		return delta/2*interactionCurrent*interactionCurrent*interactionCurrent*interactionCurrent*interactionCurrent + startValue
	}
	interactionCurrent -= 2
	return -1*delta/2*(interactionCurrent*interactionCurrent*interactionCurrent*interactionCurrent-2) + startValue
}
