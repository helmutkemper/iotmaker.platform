package tween

// simple linear tweening - no easing, no acceleration
var KLinear = func(currentTime, duration, startValue, changeInValue float64) float64 {
	return changeInValue*currentTime/duration + startValue
}
