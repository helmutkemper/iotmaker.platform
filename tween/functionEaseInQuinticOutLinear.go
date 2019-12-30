package tween

// en: sinusoidal easing in/out - accelerating until halfway, then decelerating
var KEaseInQuinticOutLinear = func(interactionCurrent, interactionTotal, startValue, delta float64) float64 {
	interactionCurrentToCalc := interactionCurrent / interactionTotal * 2
	if interactionCurrentToCalc < 1 {
		return delta/2*interactionCurrentToCalc*interactionCurrentToCalc*interactionCurrentToCalc*interactionCurrentToCalc*interactionCurrentToCalc + startValue
	}

	return delta*interactionCurrent/interactionTotal + startValue
}
