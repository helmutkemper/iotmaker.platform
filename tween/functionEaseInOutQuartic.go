package tween

// quartic easing in/out - acceleration until halfway, then deceleration
var KEaseInOutQuartic = func(currentTime, duration, startValue, changeInValue float64) float64 {
	currentTime = currentTime / duration / 2
	if currentTime < 1 {
		return changeInValue/2*currentTime*currentTime*currentTime*currentTime + startValue
	}
	currentTime -= 2
	return -1*changeInValue/2*(currentTime*currentTime*currentTime*currentTime-2) + startValue
}
