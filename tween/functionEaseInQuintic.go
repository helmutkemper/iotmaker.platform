package tween

// en: quintic easing in - accelerating from zero velocity
var KEaseInQuintic = func(interactionCurrent, interactionTotal, startValue, endValue, delta float64) float64 {
	interactionCurrent = interactionCurrent / interactionTotal
	return delta*interactionCurrent*interactionCurrent*interactionCurrent*interactionCurrent*interactionCurrent + startValue
}
