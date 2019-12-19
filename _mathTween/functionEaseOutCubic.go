package _mathTween

// en: cubic easing out - decelerating to zero velocity
var KEaseOutCubic = func(interactionCurrent, interactionTotal, startValue, delta float64) float64 {
	interactionCurrent = interactionCurrent / interactionTotal
	interactionCurrent--
	return delta*(interactionCurrent*interactionCurrent*interactionCurrent+1) + startValue
}
