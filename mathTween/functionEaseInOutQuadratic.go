package mathTween

// en: Quadratic easing in/out - acceleration until halfway, then deceleration
var KEaseInOutQuadratic = func(interactionCurrent, interactionTotal, startValue, delta float64) float64 {
	interactionCurrent /= interactionTotal / 2
	if interactionCurrent < 1 {
		return delta/2*interactionCurrent*interactionCurrent + startValue
	}
	interactionCurrent--
	return -1*delta/2*(interactionCurrent*(interactionCurrent-2)-1) + startValue
}
