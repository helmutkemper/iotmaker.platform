package tween

// en: Cubic easing in - accelerating from zero velocity
var KEaseInCubic = func(interactionCurrent, interactionTotal, startValue, endValue, delta float64) float64 {
	interactionCurrent = interactionCurrent / interactionTotal
	return delta*interactionCurrent*interactionCurrent*interactionCurrent + startValue
}
