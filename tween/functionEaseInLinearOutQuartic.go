package tween

// en: quartic easing in/out - acceleration until halfway, then deceleration
var KEaseInLinearOutQuartic = func(interactionCurrent, interactionTotal, startValue, endValue, delta float64) float64 {
	interactionCurrentToCalc := interactionCurrent / interactionTotal * 2
	if interactionCurrentToCalc < 1 {
		return delta*interactionCurrent/interactionTotal + startValue
	}
	interactionCurrentToCalc -= 2
	return -1*delta/2*(interactionCurrentToCalc*interactionCurrentToCalc*interactionCurrentToCalc*interactionCurrentToCalc-2) + startValue
}
