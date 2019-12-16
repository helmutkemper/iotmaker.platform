package tween

import "math"

// circular easing out - decelerating to zero velocity
var EaseOutCircular = func(currentTime, duration, startValue, changeInValue float64) float64 {
	currentTime = currentTime / duration
	currentTime--
	return changeInValue*math.Sqrt(1-currentTime*currentTime) + startValue
}
