package tween

var KEaseOutBounce = func(interactionCurrent, interactionTotal, startValue, endValue, delta float64) float64 {
	value := interactionCurrent / interactionTotal

	if value < 0.36363636363 {
		formula := 7.5625 * value * value
		return formula*delta + startValue
	} else if value < 0.72727272727 {
		value -= 0.54545454545
		formula := 7.5625*value*value + 0.75
		return formula*delta + startValue
	} else if value < 0.90909090909 {
		value -= 0.81818181818
		formula := 7.5625*value*value + 0.9375
		return formula*delta + startValue
	} else {
		value -= 0.95454545454
		formula := 7.5625*value*value + 0.984375
		return formula*delta + startValue
	}
}
