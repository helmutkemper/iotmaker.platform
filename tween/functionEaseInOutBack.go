package tween

import "math"

var KEaseInOutBack = func(interactionCurrent, interactionTotal, startValue, endValue, delta float64) float64 {
	value := interactionCurrent / interactionTotal / 0.5

	if value < 1 {
		return 0.5*(math.Pow(value, 2)*(3.5949095*value-2.5949095))*delta + startValue
	}
	value -= 2
	return 0.5*(value*value*(3.5949095*value+2.5949095)+2)*delta + startValue
}
