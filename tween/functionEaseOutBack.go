package tween

import "math"

var KEaseOutBack = func(interactionCurrent, interactionTotal, startValue, endValue, delta float64) float64 {
	value := interactionCurrent / interactionTotal
	value = value - 1
	return (math.Pow(value, 2.0)*(2.70158*value+1.70158)+1)*delta + startValue
}
