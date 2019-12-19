package tween

// en: quintic easing out - decelerating to zero velocity
var KEaseOutQuintic = func(currentTime, duration, startValue, changeInValue float64) float64 {
	currentTime /= duration
	currentTime--
	return changeInValue*(currentTime*currentTime*currentTime*currentTime*currentTime+1) + startValue
}
