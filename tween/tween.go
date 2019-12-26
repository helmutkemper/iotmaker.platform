package tween

import (
	"github.com/helmutkemper/iotmaker.platform/fps"
	"time"
)

type Tween struct {
	FramesPerSecond int
	StartValue      float64
	EndValue        float64
	Arguments       []interface{}
	ticker          *time.Ticker
	startTime       time.Time
	Duration        time.Duration
	Func            func(currentTime, duration, startValue, changeInValue float64) float64
	Interaction     func(value, percentToComplete float64, arguments []interface{})
	OnStart         func(value float64)
	OnEnd           func(value float64)
	invert          bool
	Repeat          int
}

func (el *Tween) Start() {
	if el.FramesPerSecond == 0 {
		el.FramesPerSecond = fps.Get()
	}

	el.ticker = time.NewTicker(time.Second / time.Duration(el.FramesPerSecond))
	el.startTime = time.Now()
	el.invert = true
	go el.tickerRunner(el.StartValue, el.EndValue)
}

func (el *Tween) tickerRunner(startValue, endValue float64) {
	for {
		if el.Func == nil {
			return
		}

		if el.OnStart != nil {
			el.OnStart(el.StartValue)
		}

		select {
		case <-el.ticker.C:
			elapsed := time.Since(el.startTime)
			value := el.Func(elapsed.Seconds(), el.Duration.Seconds(), startValue, endValue-startValue)
			percent := elapsed.Seconds() / el.Duration.Seconds()

			if el.Interaction != nil {
				el.Interaction(value, percent, el.Arguments)
			}

			if elapsed >= el.Duration {

				if el.OnEnd != nil {
					el.OnEnd(value)
				}

				if el.Repeat != 0 {
					el.startTime = time.Now()

					if el.invert == true {
						go el.tickerRunner(el.EndValue, el.StartValue)
					} else {
						go el.tickerRunner(el.StartValue, el.EndValue)
					}

					el.invert = !el.invert
					el.Repeat -= 1
				}

				return
			}
		}
	}
}
