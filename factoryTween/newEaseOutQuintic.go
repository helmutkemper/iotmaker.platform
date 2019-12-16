package factoryTween

import (
	"github.com/helmutkemper/iotmaker.platform/tween"
	"time"
)

// en: quintic easing out - decelerating to zero velocity
func NewEaseOutQuintic(duration time.Duration, startValue, endValue float64, interactionFunc, doneFunc func(float64)) *tween.Tween {
	t := &tween.Tween{
		Duration:    duration,
		StartValue:  startValue,
		EndValue:    endValue,
		Func:        tween.KEaseOutQuintic,
		Interaction: interactionFunc,
		Done:        doneFunc,
	}
	t.Start()

	return t
}
