package factoryTween

import (
	"github.com/helmutkemper/iotmaker.platform/tween"
	"time"
)

// en: sinusoidal easing in/out - accelerating until halfway, then decelerating
func NewEaseInOutSine(duration time.Duration, startValue, endValue float64, interactionFunc, doneFunc func(float64)) *tween.Tween {
	t := &tween.Tween{
		Duration:    duration,
		StartValue:  startValue,
		EndValue:    endValue,
		Func:        tween.KEaseInOutSine,
		Interaction: interactionFunc,
		Done:        doneFunc,
	}
	t.Start()

	return t
}
