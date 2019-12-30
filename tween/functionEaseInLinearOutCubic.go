package tween

// en: cubic easing in/out - acceleration until halfway, then deceleration
var KEaseInLinearOutCubic = func(interactionCurrent, interactionTotal, startValue, delta float64) float64 {
	interactionCurrentToCalc := interactionCurrent / interactionTotal * 2
	if interactionCurrentToCalc < 1 {
		return delta*interactionCurrent/interactionTotal + startValue
	}
	interactionCurrentToCalc -= 2
	return delta/2*(interactionCurrentToCalc*interactionCurrentToCalc*interactionCurrentToCalc+2) + startValue
}
