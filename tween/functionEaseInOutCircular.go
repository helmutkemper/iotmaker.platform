package tween

import "math"

// en: circular easing in/out - acceleration until halfway, then deceleration
var KEaseInOutCircular = func(currentTime, duration, startValue, changeInValue float64) float64 {
	currentTime = currentTime / duration / 2
	if currentTime < 1 {
		return -1*changeInValue/2*(math.Sqrt(1-currentTime*currentTime)-1) + startValue
	}
	currentTime -= 2
	return changeInValue/2*(math.Sqrt(1-currentTime*currentTime)+1) + startValue
}
