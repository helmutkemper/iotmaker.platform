package tween

import "math"

// exponential easing in - accelerating from zero velocity
var EaseInExponential = func(currentTime, duration, startValue, changeInValue float64) float64 {
	return changeInValue*math.Pow(2, 10*(currentTime/duration-1)) + startValue
}
