package factoryTween

import (
	"github.com/helmutkemper/iotmaker.platform/tween"
	"time"
)

// en: Circular easing in - accelerating from zero velocity
func NewEaseInCircular(duration time.Duration, startValue, endValue float64, interactionFunc, doneFunc func(float64)) *tween.Tween {
	t := &tween.Tween{
		Duration:    duration,
		StartValue:  startValue,
		EndValue:    endValue,
		Func:        tween.KEaseInCircular,
		Interaction: interactionFunc,
		Done:        doneFunc,
	}
	t.Start()

	return t
}
