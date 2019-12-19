package tween

// en: quadratic easing out - decelerating to zero velocity
var KEaseOutQuadratic = func(interactionCurrent, interactionTotal, startValue, delta float64) float64 {
	interactionCurrent = interactionCurrent / interactionTotal
	return -1*delta*interactionCurrent*(interactionCurrent-2) + startValue
}
