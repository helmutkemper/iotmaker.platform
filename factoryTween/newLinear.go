package factoryTween

import (
	"github.com/helmutkemper/iotmaker.platform/tween"
	"time"
)

// en: simple linear tweening - no easing, no acceleration
func NewLinear(duration time.Duration, startValue, endValue float64, interactionFunc, doneFunc func(float64)) *tween.Tween {
	t := &tween.Tween{
		Duration:    duration,
		StartValue:  startValue,
		EndValue:    endValue,
		Func:        tween.KLinear,
		Interaction: interactionFunc,
		Done:        doneFunc,
	}
	t.Start()

	return t
}
