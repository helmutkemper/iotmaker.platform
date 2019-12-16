package tween

// en: Quadratic easing in/out - acceleration until halfway, then deceleration
var KEaseInOutQuadratic = func(currentTime, duration, startValue, changeInValue float64) float64 {
	currentTime /= duration / 2
	if currentTime < 1 {
		return changeInValue/2*currentTime*currentTime + startValue
	}
	currentTime--
	return -1*changeInValue/2*(currentTime*(currentTime-2)-1) + startValue
}
