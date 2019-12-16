package tween

// quartic easing out - decelerating to zero velocity
var KEaseOutQuartic = func(currentTime, duration, startValue, changeInValue float64) float64 {
	currentTime = currentTime / duration
	currentTime--
	return -1*changeInValue*(currentTime*currentTime*currentTime*currentTime-1) + startValue
}
