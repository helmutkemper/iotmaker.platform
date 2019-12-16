package tween

import "math"

// exponential easing in/out - accelerating until halfway, then decelerating
var EaseInOutExponential = func(currentTime, duration, startValue, changeInValue float64) float64 {
	currentTime /= duration / 2
	if currentTime < 1 {
		return changeInValue/2*math.Pow(2, 10*(currentTime-1)) + startValue
	}
	currentTime--
	return changeInValue/2*(-1*math.Pow(2, -10*currentTime)+2) + startValue
}
