package _mathTween

// en: quintic easing out - decelerating to zero velocity
var KEaseOutQuintic = func(interactionCurrent, interactionTotal, startValue, delta float64) float64 {
	interactionCurrent /= interactionTotal
	interactionCurrent--
	return delta*(interactionCurrent*interactionCurrent*interactionCurrent*interactionCurrent*interactionCurrent+1) + startValue
}
