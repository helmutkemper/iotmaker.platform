package tween

// en: quintic easing in - accelerating from zero velocity
var KEaseInQuintic = func(currentTime, duration, startValue, changeInValue float64) float64 {
	currentTime = currentTime / duration
	return changeInValue*currentTime*currentTime*currentTime*currentTime*currentTime + startValue
}
