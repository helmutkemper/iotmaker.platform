package factoryTween

import (
	"github.com/helmutkemper/iotmaker.platform/tween"
	"time"
)

// en: circular easing in/out - acceleration until halfway, then deceleration
func NewEaseInOutCircular(duration time.Duration, startValue, endValue float64, interactionFunc, doneFunc func(float64)) *tween.Tween {
	t := &tween.Tween{
		Duration:    duration,
		StartValue:  startValue,
		EndValue:    endValue,
		Func:        tween.KEaseInOutCircular,
		Interaction: interactionFunc,
		Done:        doneFunc,
	}
	t.Start()

	return t
}
