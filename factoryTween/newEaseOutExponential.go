package factoryTween

import (
	"github.com/helmutkemper/iotmaker.platform/tween"
	"time"
)

// en: exponential easing out - decelerating to zero velocity
func NewEaseOutExponential(duration time.Duration, startValue, endValue float64, interactionFunc, doneFunc func(float64)) *tween.Tween {
	t := &tween.Tween{
		Duration:    duration,
		StartValue:  startValue,
		EndValue:    endValue,
		Func:        tween.KEaseOutExponential,
		Interaction: interactionFunc,
		Done:        doneFunc,
	}
	t.Start()

	return t
}
