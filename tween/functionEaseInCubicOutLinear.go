package tween

// en: cubic easing in/out - acceleration until halfway, then deceleration
var KEaseInCubicOutLinear = func(interactionCurrent, interactionTotal, startValue, endValue, delta float64) float64 {
	interactionCurrentToCalc := interactionCurrent / interactionTotal * 2
	if interactionCurrentToCalc < 1 {
		return delta/2*interactionCurrentToCalc*interactionCurrentToCalc*interactionCurrentToCalc + startValue
	}
	return delta*interactionCurrent/interactionTotal + startValue
}
