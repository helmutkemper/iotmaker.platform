package mathTween

import "math"

// en: circular easing out - decelerating to zero velocity
var KEaseOutCircular = func(interactionCurrent, interactionTotal, startValue, delta float64) float64 {
	interactionCurrent = interactionCurrent / interactionTotal
	interactionCurrent--
	return delta*math.Sqrt(1-interactionCurrent*interactionCurrent) + startValue
}
