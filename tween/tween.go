package tween

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/engine"
	"time"
)

type Tween struct {
	Engine             engine.IEngine
	StartValue         float64
	EndValue           float64
	Arguments          []interface{}
	startTime          time.Time
	Duration           time.Duration
	Func               func(currentTime, duration, startValue, changeInValue float64) float64
	Interaction        func(value, percentToComplete float64, arguments ...interface{})
	OnCycleStart       func(value float64, arguments ...interface{})
	OnCycleEnd         func(value float64, arguments ...interface{})
	OnStart            func(value float64, arguments ...interface{})
	OnEnd              func(value float64, arguments ...interface{})
	OnInvert           func(value float64, arguments ...interface{})
	DoNotReverseMotion bool
	invert             bool
	Repeat             int
	fpsUId             string
	loopStartValue     float64
	loopEndValue       float64
}

func (el *Tween) Start() {
	el.startTime = time.Now()
	el.invert = true

	if el.Func == nil {
		return
	}

	if el.OnStart != nil {
		el.OnStart(el.StartValue, el.Arguments)
	}

	go el.tickerRunnerPrepare(el.StartValue, el.EndValue)
}

func (el *Tween) tickerRunnerPrepare(startValue, endValue float64) {
	if el.OnCycleStart != nil {
		el.OnCycleStart(el.StartValue, el.Arguments)
	}

	el.loopStartValue = startValue
	el.loopEndValue = endValue

	// fixme: _ é um index do array? tem uso aqui?
	el.fpsUId, _ = el.Engine.MathAddToFunctions(el.tickerRunnerRun)
}

func (el *Tween) tickerRunnerRun() {
	elapsed := time.Since(el.startTime)
	value := el.Func(elapsed.Seconds(), el.Duration.Seconds(), el.loopStartValue, el.loopEndValue-el.loopStartValue)
	percent := elapsed.Seconds() / el.Duration.Seconds()

	if el.Interaction != nil {
		el.Interaction(value, percent, el.Arguments)
	}

	if elapsed >= el.Duration {

		el.Stop()

		if el.Repeat == 0 {
			if el.OnEnd != nil {
				el.OnEnd(value)
			}
		}

		if el.Repeat != 0 {
			el.startTime = time.Now()

			if el.OnInvert != nil {
				el.OnInvert(value)
			}

			if el.DoNotReverseMotion == true {
				el.tickerRunnerPrepare(el.StartValue, el.EndValue)
			} else {
				if el.invert == true {
					el.tickerRunnerPrepare(el.EndValue, el.StartValue)
				} else {
					el.tickerRunnerPrepare(el.StartValue, el.EndValue)
				}
				el.invert = !el.invert
			}

			el.Repeat -= 1
		}
	}
}

func (el *Tween) End() {
	el.Engine.MathDeleteFromFunctions(el.fpsUId)
}

func (el *Tween) Stop() {
	el.Engine.MathDeleteFromFunctions(el.fpsUId)
	if el.OnCycleEnd != nil {
		el.OnCycleEnd(el.EndValue, el.Arguments)
	}
}
