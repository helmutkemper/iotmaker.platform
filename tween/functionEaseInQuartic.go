package tween

// quartic easing in - accelerating from zero velocity
var KEaseInQuartic = func(currentTime, duration, startValue, changeInValue float64) float64 {
	currentTime /= duration
	return changeInValue*currentTime*currentTime*currentTime*currentTime + startValue
}
