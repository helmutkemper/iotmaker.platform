package tween

// quadratic easing out - decelerating to zero velocity
var KEaseOutQuadratic = func(currentTime, duration, startValue, changeInValue float64) float64 {
	currentTime = currentTime / duration
	return -1*changeInValue*currentTime*(currentTime-2) + startValue
}
