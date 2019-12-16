package factoryTween

import (
	"github.com/helmutkemper/iotmaker.platform/tween"
	"time"
)

// en: exponential easing in/out - accelerating until halfway, then decelerating
func NewEaseInOutExponential(duration time.Duration, startValue, endValue float64, interactionFunc, doneFunc func(float64)) *tween.Tween {
	t := &tween.Tween{
		Duration:    duration,
		StartValue:  startValue,
		EndValue:    endValue,
		Func:        tween.KEaseInOutExponential,
		Interaction: interactionFunc,
		Done:        doneFunc,
	}
	t.Start()

	return t
}
