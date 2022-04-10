package tween

import "math"

// en: circular easing in - accelerating from zero velocity

// KEaseInCircular
//
// English:
//
//  Circular easing in, accelerating from zero velocity
//
// Português:
//
//  Circular easing in, acelerando do zero até a velocidade
var KEaseInCircular = func(interactionCurrent, interactionTotal, startValue, delta float64) float64 {
	interactionCurrent = interactionCurrent / interactionTotal
	return -delta*(math.Sqrt(math.Abs(1-interactionCurrent*interactionCurrent))-1) + startValue
}
