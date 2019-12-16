package tween

import "time"

var fps = 60.0

func SetFPS(value float64) {
	fps = value
}

func GetFPS() float64 {
	return fps
}

type Tween struct {
	StartValue  float64
	EndValue    float64
	ticker      *time.Ticker
	startTime   time.Time
	Duration    time.Duration
	Func        func(currentTime, duration time.Duration, startValue, changeInValue float64) float64
	Interaction func(value float64)
	Done        func(value float64)
	abortLoop   float64
	abortStep   float64
}

func (el *Tween) Start() {
	el.ticker = time.NewTicker(time.Second / time.Duration(fps))
	el.startTime = time.Now()
	el.abortStep = el.Duration.Seconds() / GetFPS()
	go el.tickerRunner()
}

func (el *Tween) tickerRunner() {
	for {
		if el.Func == nil || el.Interaction == nil {
			return
		}

		select {
		case <-el.ticker.C:
			elapsed := time.Since(el.startTime)
			value := el.Func(elapsed, el.Duration, el.StartValue, el.EndValue-el.StartValue)
			el.Interaction(value)
			el.abortLoop += el.abortStep
			if el.abortLoop >= el.Duration.Seconds() {

				if el.Done != nil {
					el.Done(value)
				}

				return
			}
		}
	}
}

func (el *Tween) SetTime(value time.Duration) {
	el.Duration = value
}

var Linear = func(currentTime, duration time.Duration, startValue, changeInValue float64) float64 {
	return changeInValue*currentTime.Seconds()/duration.Seconds() + startValue
}
