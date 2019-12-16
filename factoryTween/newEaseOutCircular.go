package factoryTween

import (
	"github.com/helmutkemper/iotmaker.platform/tween"
	"time"
)

// en: circular easing out - decelerating to zero velocity
func NewEaseOutCircular(duration time.Duration, startValue, endValue float64, interactionFunc, doneFunc func(float64)) *tween.Tween {
	t := &tween.Tween{
		Duration:    duration,
		StartValue:  startValue,
		EndValue:    endValue,
		Func:        tween.KEaseOutCircular,
		Interaction: interactionFunc,
		Done:        doneFunc,
	}
	t.Start()

	return t
}
