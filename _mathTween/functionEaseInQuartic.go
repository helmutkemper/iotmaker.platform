package _mathTween

// en: quartic easing in - accelerating from zero velocity
var KEaseInQuartic = func(interactionCurrent, interactionTotal, startValue, delta float64) float64 {
	interactionCurrent /= interactionTotal
	return delta*interactionCurrent*interactionCurrent*interactionCurrent*interactionCurrent + startValue
}
