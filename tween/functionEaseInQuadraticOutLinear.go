package tween

// en: Quadratic easing in/out - acceleration until halfway, then deceleration
var KEaseInQuadraticOutLinear = func(interactionCurrent, interactionTotal, startValue, delta float64) float64 {
	interactionCurrentToCalc := interactionCurrent / interactionTotal * 2
	if interactionCurrentToCalc < 1 {
		return delta/2*interactionCurrentToCalc*interactionCurrentToCalc + startValue
	}
	return delta*interactionCurrent/interactionTotal + startValue
}
