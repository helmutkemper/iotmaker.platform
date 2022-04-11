package tween

// en: quartic easing out - decelerating to zero velocity
var KEaseOutQuartic = func(interactionCurrent, interactionTotal, startValue, endValue, delta float64) float64 {
	interactionCurrent = interactionCurrent / interactionTotal
	interactionCurrent--
	return -1*delta*(interactionCurrent*interactionCurrent*interactionCurrent*interactionCurrent-1) + startValue
}
