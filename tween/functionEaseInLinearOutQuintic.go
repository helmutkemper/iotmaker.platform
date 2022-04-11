package tween

// en: sinusoidal easing in/out - accelerating until halfway, then decelerating
var KEaseInLinearOutQuintic = func(interactionCurrent, interactionTotal, startValue, endValue, delta float64) float64 {
	interactionCurrentToCalc := interactionCurrent / interactionTotal * 2
	if interactionCurrentToCalc < 1 {
		return delta*interactionCurrent/interactionTotal + startValue
	}
	interactionCurrentToCalc -= 2
	return delta/2*(interactionCurrentToCalc*interactionCurrentToCalc*interactionCurrentToCalc*interactionCurrentToCalc*interactionCurrentToCalc+2) + startValue
}
