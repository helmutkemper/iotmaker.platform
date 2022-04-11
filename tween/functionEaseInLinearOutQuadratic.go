package tween

// en: Quadratic easing in/out - acceleration until halfway, then deceleration
var KEaseInLinearOutQuadratic = func(interactionCurrent, interactionTotal, startValue, endValue, delta float64) float64 {
	interactionCurrentToCalc := interactionCurrent / interactionTotal * 2
	if interactionCurrentToCalc < 1 {
		return delta*interactionCurrent/interactionTotal + startValue
	}
	interactionCurrentToCalc--
	return -1*delta/2*(interactionCurrentToCalc*(interactionCurrentToCalc-2)-1) + startValue
}
