package tween

import "math"

var KEaseInBack = func(interactionCurrent, interactionTotal, startValue, endValue, delta float64) float64 {
	value := interactionCurrent / interactionTotal
	return math.Pow(value, 2.0)*(1.70158+1*value-1.70158)*delta + startValue
}
