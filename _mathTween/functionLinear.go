package _mathTween

// en: simple linear tweening - no easing, no acceleration
var KLinear = func(interactionCurrent, interactionTotal, startValue, delta float64) float64 {
	return delta*interactionCurrent/interactionTotal + startValue
}
