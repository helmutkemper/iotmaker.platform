package tween

// quintic easing in - accelerating from zero velocity
var EaseInQuintic = func(currentTime, duration, startValue, changeInValue float64) float64 {
	currentTime = currentTime / duration
	return changeInValue*currentTime*currentTime*currentTime*currentTime*currentTime + startValue
}
