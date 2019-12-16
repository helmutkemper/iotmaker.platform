package factoryTween

import (
	"github.com/helmutkemper/iotmaker.platform/tween"
	"time"
)

func NewTweenEaseInQuadratic(duration time.Duration, startValue, endValue float64, interactionFunc, doneFunc func(float64)) *tween.Tween {
	t := &tween.Tween{
		Duration:    duration,
		StartValue:  startValue,
		EndValue:    endValue,
		Func:        tween.KEaseInQuadratic,
		Interaction: interactionFunc,
		Done:        doneFunc,
	}
	t.Start()

	return t
}
