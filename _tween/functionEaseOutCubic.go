package tween

// en: cubic easing out - decelerating to zero velocity
var KEaseOutCubic = func(currentTime, duration, startValue, changeInValue float64) float64 {
	currentTime = currentTime / duration
	currentTime--
	return changeInValue*(currentTime*currentTime*currentTime+1) + startValue
}
