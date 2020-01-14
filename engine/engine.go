package engine

import (
	"math"
	"math/rand"
	"time"
)

const kUIdSize = 10

type funcList struct {
	id string
	f  func()
}

type Engine struct {
	sleepFrame    int
	fps           int
	fpsLowLatency int
	kUIdCharList  []string

	// en: Warning! stopTicker should be a channel, however, conflict with webassembly <-done channer
	// pt_br: Cuidado! stopTicker deveria ser um channel, porém, deu conflito com o webassembly <-done channer
	stopTicker bool

	ticker           *time.Ticker
	tickerLowLatency *time.Ticker

	funcListToHighLatency []funcList
	funcListToSystem      []funcList
	funcListToCalculate   []funcList
	funcListToDraw        []funcList

	funcCursorDraw funcList

	// pt_br: impede que o loop ocorra em intervalos muitos próximos e trave o
	// processamento do browser para outras tarefas
	slipFrame          int
	slipFrameTimeAlarm time.Duration
}

func (el *Engine) Init() {
	// fixme: must be a interval of time
	el.sleepFrame = 2
	el.fps = 60
	el.fpsLowLatency = 1

	el.kUIdCharList = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s",
		"t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P",
		"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "_", "!", "@",
		"#", "$", "%", "&", "*", "(", ")", "-", "_", "+", "=", "[", "{", "}", "]", "/", "?", ";", ":", ".", ",", "<", ">",
		"|"}
	el.funcListToSystem = make([]funcList, 0)
	el.funcListToCalculate = make([]funcList, 0)
	el.funcListToDraw = make([]funcList, 0)
	el.tickerStart()
}

func (el *Engine) SetSleepFrame(value int) {
	el.sleepFrame = value
}

func (el *Engine) GetSleepFrame() int {
	return el.sleepFrame
}

func (el *Engine) Set(value int) {
	el.fps = value
	el.stopTicker = true
}

func (el *Engine) Get() int {
	return el.fps
}

func (el *Engine) AddCursorDrawFunc(runnerFunc func()) string {
	UId := el.getUId()
	el.funcCursorDraw = funcList{id: UId, f: runnerFunc}

	return UId
}

func (el *Engine) RemoveCursorDrawFunc(id string) {
	el.funcCursorDraw = funcList{}
}

func (el *Engine) AddToHighLatency(runnerFunc func()) string {
	UId := el.getUId()
	el.funcListToHighLatency = append(el.funcListToHighLatency, funcList{id: UId, f: runnerFunc})

	return UId
}

func (el *Engine) DeleteFromHighLatency(UId string) {
	for k, runner := range el.funcListToHighLatency {
		if runner.id == UId {
			el.funcListToHighLatency = append(el.funcListToHighLatency[:k], el.funcListToHighLatency[k+1:]...)
			break
		}
	}
}

func (el *Engine) AddToSystem(runnerFunc func()) string {
	UId := el.getUId()
	el.funcListToSystem = append(el.funcListToSystem, funcList{id: UId, f: runnerFunc})

	return UId
}

func (el *Engine) DeleteFromSystem(UId string) {
	for k, runner := range el.funcListToSystem {
		if runner.id == UId {
			el.funcListToSystem = append(el.funcListToSystem[:k], el.funcListToSystem[k+1:]...)
			break
		}
	}
}

func (el *Engine) AddToCalculate(runnerFunc func()) string {
	UId := el.getUId()
	el.funcListToCalculate = append(el.funcListToCalculate, funcList{id: UId, f: runnerFunc})

	return UId
}

func (el *Engine) DeleteFromCalculate(UId string) {
	for k, runner := range el.funcListToCalculate {
		if runner.id == UId {
			el.funcListToCalculate = append(el.funcListToCalculate[:k], el.funcListToCalculate[k+1:]...)
			break
		}
	}
}

func (el *Engine) AddToDraw(runnerFunc func()) string {
	UId := el.getUId()
	el.funcListToDraw = append(el.funcListToDraw, funcList{id: UId, f: runnerFunc})

	return UId
}

func (el *Engine) DeleteFromDraw(UId string) {
	for k, runner := range el.funcListToDraw {
		if runner.id == UId {
			el.funcListToDraw = append(el.funcListToDraw[:k], el.funcListToDraw[k+1:]...)
			break
		}
	}
}

func (el *Engine) SetZIndex(UId string, index int) int {
	var function funcList
	var pass = false
	var length = len(el.funcListToDraw)

	if index < 0 || index > length-1 {
		return math.MaxInt32
	}

	for k, runner := range el.funcListToDraw {
		if runner.id == UId {
			pass = true
			function = runner
			el.funcListToDraw = append(el.funcListToDraw[:k], el.funcListToDraw[k+1:]...)
			break
		}
	}

	if pass == false {
		return math.MaxInt32
	}

	if index == 0 {

		el.funcListToDraw = append([]funcList{function}, el.funcListToDraw...)

	} else if index == length-1 {

		el.funcListToDraw = append(el.funcListToDraw, function)

	} else {

		firstPart := make([]funcList, len(el.funcListToDraw[:index]))
		lastPart := make([]funcList, len(el.funcListToDraw[index:]))

		copy(firstPart, el.funcListToDraw[:index])
		copy(lastPart, el.funcListToDraw[index:])

		firstPart = append(firstPart, function)

		el.funcListToDraw = make([]funcList, 0)
		el.funcListToDraw = append(firstPart, lastPart...)
	}

	return index
}

func (el *Engine) ToFront(UId string) int {
	var function funcList
	var pass = false
	for k, runner := range el.funcListToDraw {
		if runner.id == UId {
			pass = true
			function = runner
			el.funcListToDraw = append(el.funcListToDraw[:k], el.funcListToDraw[k+1:]...)
			break
		}
	}

	if pass == false {
		return math.MaxInt32
	}

	el.funcListToDraw = append(el.funcListToDraw, function)

	return 0
}

func (el *Engine) ToBack(UId string) int {
	var function funcList
	var pass = false
	for k, runner := range el.funcListToDraw {
		if runner.id == UId {
			pass = true
			function = runner
			el.funcListToDraw = append(el.funcListToDraw[:k], el.funcListToDraw[k+1:]...)
			break
		}
	}

	if pass == false {
		return math.MaxInt32
	}

	el.funcListToDraw = append([]funcList{function}, el.funcListToDraw...)

	return len(el.funcListToDraw) - 1
}

func (el *Engine) GetZIndex(UId string) int {
	for k, runner := range el.funcListToDraw {
		if runner.id == UId {
			return k
		}
	}

	return math.MaxInt32
}

func (el *Engine) getUId() string {
	var UId = ""
	for i := 0; i != kUIdSize; i += 1 {
		UId += el.kUIdCharList[rand.Intn(len(el.kUIdCharList)-1)]
	}

	return UId
}

func (el *Engine) tickerStart() {
	el.ticker = time.NewTicker(time.Second / time.Duration(el.fps))
	el.tickerLowLatency = time.NewTicker(time.Second / time.Duration(el.fpsLowLatency))
	el.slipFrameTimeAlarm = time.Second / time.Duration(el.fps)
	go func() { el.tickerRunner() }()
}

func (el *Engine) tickerRunner() {
	defer el.tickerStart()
	for {
		select {
		case <-el.tickerLowLatency.C:

			for _, runnerFunc := range el.funcListToHighLatency {
				if runnerFunc.f != nil {
					runnerFunc.f()
				}
			}

		case <-el.ticker.C:

			if el.slipFrame != 0 {
				el.slipFrame -= 1
				continue
			}

			if el.stopTicker == true {
				el.stopTicker = false
				return
			}

			start := time.Now()

			for _, runnerFunc := range el.funcListToSystem {
				if runnerFunc.f != nil {
					runnerFunc.f()
				}
			}

			for _, runnerFunc := range el.funcListToCalculate {
				if runnerFunc.f != nil {
					runnerFunc.f()
				}
			}

			for _, runnerFunc := range el.funcListToDraw {
				if runnerFunc.f != nil {
					runnerFunc.f()
				}
			}

			if el.funcCursorDraw.f != nil {
				el.funcCursorDraw.f()
			}

			elapsed := time.Since(start)
			if elapsed > el.slipFrameTimeAlarm {
				el.slipFrame = el.sleepFrame
			}
		}
	}
}
