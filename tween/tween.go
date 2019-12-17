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
	Done            func(value float64)
}

func (el *Tween) Start() {
	if el.FramesPerSecond == 0 {
		el.FramesPerSecond = fps.Get()
	}

	el.ticker = time.NewTicker(time.Second / time.Duration(el.FramesPerSecond))
	el.startTime = time.Now()
	go el.tickerRunner()
}

func (el *Tween) tickerRunner() {
	for {
		if el.Func == nil {
			return
		}

		select {
		case <-el.ticker.C:
			elapsed := time.Since(el.startTime)
			value := el.Func(elapsed.Seconds(), el.Duration.Seconds(), el.StartValue, el.EndValue-el.StartValue)
			percent := elapsed.Seconds() / el.Duration.Seconds()

			if el.Interaction != nil {
				el.Interaction(value, percent, el.Arguments)
			}

			if elapsed >= el.Duration {

				if el.Done != nil {
					el.Done(value)
				}

				return
			}
		}
	}
}
