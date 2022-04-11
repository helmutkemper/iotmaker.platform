package tween

// en: quadratic easing in - accelerating from zero velocity
var KEaseInQuadratic = func(interactionCurrent, interactionTotal, startValue, endValue, delta float64) float64 {
	interactionCurrent = interactionCurrent / interactionTotal
	return delta*interactionCurrent*interactionCurrent + startValue
}
