package tween

import "math"

// sinusoidal easing in/out - accelerating until halfway, then decelerating
var EaseInOutSine = func(currentTime, duration, startValue, changeInValue float64) float64 {
	return -1*changeInValue/2*(math.Cos(math.Pi*currentTime/duration)-1) + startValue
}
