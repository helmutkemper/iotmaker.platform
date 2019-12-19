package tween

import "math"

// en: circular easing out - decelerating to zero velocity
var KEaseOutCircular = func(currentTime, duration, startValue, changeInValue float64) float64 {
	currentTime = currentTime / duration
	currentTime--
	return changeInValue*math.Sqrt(1-currentTime*currentTime) + startValue
}
