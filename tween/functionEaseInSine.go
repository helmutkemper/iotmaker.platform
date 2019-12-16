package tween

import "math"

// sinusoidal easing in - accelerating from zero velocity
var KEaseInSine = func(currentTime, duration, startValue, changeInValue float64) float64 {
	return -1*changeInValue*math.Cos(currentTime/duration*(math.Pi/2)) + changeInValue + startValue
}
