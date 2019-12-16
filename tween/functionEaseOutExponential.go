package tween

import "math"

// exponential easing out - decelerating to zero velocity
var EaseOutExponential = func(currentTime, duration, startValue, changeInValue float64) float64 {
	return changeInValue*(-1*math.Pow(2, -10*currentTime/duration)+1) + startValue
}
