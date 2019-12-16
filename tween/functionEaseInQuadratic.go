package tween

// en: quadratic easing in - accelerating from zero velocity
var KEaseInQuadratic = func(currentTime, duration, startValue, changeInValue float64) float64 {
	currentTime = currentTime / duration
	return changeInValue*currentTime*currentTime + startValue
}
