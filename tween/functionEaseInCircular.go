package tween

import "math"

// en: circular easing in - accelerating from zero velocity
var KEaseInCircular = func(currentTime, duration, startValue, changeInValue float64) float64 {
	currentTime = currentTime / duration
	return -changeInValue*(math.Sqrt(1-currentTime*currentTime)-1) + startValue
}
