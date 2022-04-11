package tween

// en: Quadratic easing in/out - acceleration until halfway, then deceleration
var KEaseInQuadraticOutQuintic = func(interactionCurrent, interactionTotal, startValue, endValue, delta float64) float64 {
	interactionCurrent /= interactionTotal / 2
	if interactionCurrent < 1 {
		return delta/2*interactionCurrent*interactionCurrent + startValue
	}
	interactionCurrent -= 2
	return delta/2*(interactionCurrent*interactionCurrent*interactionCurrent*interactionCurrent*interactionCurrent+2) + startValue
}
