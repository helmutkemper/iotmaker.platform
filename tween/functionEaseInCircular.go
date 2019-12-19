package tween

import "math"

// en: circular easing in - accelerating from zero velocity
var KEaseInCircular = func(interactionCurrent, interactionTotal, startValue, delta float64) float64 {
	interactionCurrent = interactionCurrent / interactionTotal
	return -delta*(math.Sqrt(math.Abs(1-interactionCurrent*interactionCurrent))-1) + startValue
}
