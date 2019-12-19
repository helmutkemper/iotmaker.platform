package tween

// en: Cubic easing in - accelerating from zero velocity
var KEaseInCubic = func(currentTime, duration, startValue, changeInValue float64) float64 {
	currentTime = currentTime / duration
	return changeInValue*currentTime*currentTime*currentTime + startValue
}
