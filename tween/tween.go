package tween

import (
	"github.com/helmutkemper/iotmaker.platform/fps"
	"time"
)

type Tween struct {
	StartValue     float64
	EndValue       float64
	Arguments      []interface{}
	startTime      time.Time
	Duration       time.Duration
	Func           func(currentTime, duration, startValue, changeInValue float64) float64
	Interaction    func(value, percentToComplete float64, arguments []interface{})
	OnStart        func(value float64)
	OnEnd          func(value float64)
	OnInvert       func(value float64)
	invert         bool
	Repeat         int
	fpsUId         string
	loopStartValue float64
	loopEndValue   float64
}

func (el *Tween) Start() {
	el.startTime = time.Now()
	el.invert = true
	go el.tickerRunnerPrepare(el.StartValue, el.EndValue)
}

func (el *Tween) tickerRunnerPrepare(startValue, endValue float64) {
	if el.Func == nil {
		return
	}

	if el.OnStart != nil {
		el.OnStart(el.StartValue)
	}

	el.loopStartValue = startValue
	el.loopEndValue = endValue

	el.fpsUId = fps.AddToRunnerPriorityFunc(el.tickerRunnerRun, true)
}

func (el *Tween) tickerRunnerRun() {
	elapsed := time.Since(el.startTime)
	value := el.Func(elapsed.Seconds(), el.Duration.Seconds(), el.loopStartValue, el.loopEndValue-el.loopStartValue)
	percent := elapsed.Seconds() / el.Duration.Seconds()

	if el.Interaction != nil {
		el.Interaction(value, percent, el.Arguments)
	}

	if elapsed >= el.Duration {

		if el.Repeat == 0 {
			if el.OnEnd != nil {
				el.OnEnd(value)
			}
		}

		el.tickerRunnerDelete()

		if el.Repeat != 0 {
			el.startTime = time.Now()

			if el.OnInvert != nil {
				el.OnInvert(value)
			}

			if el.invert == true {
				el.tickerRunnerPrepare(el.EndValue, el.StartValue)
			} else {
				el.tickerRunnerPrepare(el.StartValue, el.EndValue)
			}

			el.invert = !el.invert
			el.Repeat -= 1
		}
	}
}

func (el *Tween) tickerRunnerDelete() {
	fps.DeleteFromRunnerPriorityFunc(el.fpsUId)
}
