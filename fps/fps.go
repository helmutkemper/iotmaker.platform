package fps

import (
	"math/rand"
	"time"
)

const kUIdSize = 10

var fps = 120
var fpsCache = 10
var kUIdCharList []string

// en: Warning! stopTicker should be a channel, however, conflict with webassembly <-done channer
// pt_br: Cuidado! stopTicker deveria ser um channel, porém, deu conflito com o webassembly <-done channer
var stopTicker bool
var ticker *time.Ticker
var tickerCache *time.Ticker
var funcListToRunner map[string]func()
var funcListToCacheRunner map[string]func()
var funcListPriorityToRunner map[string]func()

// pt_br: impede que o loop ocorra em intervalos muitos próximos e trave o
// processamento do browser para outras tarefas
var slipFrame int = 0
var slipFrameTimeAlarm time.Duration

func Set(value int) {
	fps = value
	stopTicker = true
}

func SetCacheUpdate(value int) {
	fpsCache = value
	stopTicker = true
}

func Get() int {
	return fps
}

func GetCacheUpdate() int {
	return fpsCache
}

func AddToRunner(runnerFunc func()) string {
	UId := getUId()
	funcListToRunner[UId] = runnerFunc

	return UId
}

func AddToCacheRunner(runnerFunc func()) string {
	UId := getUId()
	funcListToCacheRunner[UId] = runnerFunc

	return UId
}

func DeleteFromRunner(UId string) {
	delete(funcListToRunner, UId)
}

func DeleteFromCacheRunner(UId string) {
	delete(funcListToCacheRunner, UId)
}

func AddToRunnerPriorityFunc(runnerFunc func()) string {
	UId := getUId()
	funcListPriorityToRunner[UId] = runnerFunc

	return UId
}

func DeleteFromRunnerPriorityFunc(UId string) {
	delete(funcListPriorityToRunner, UId)
}

func init() {
	kUIdCharList = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s",
		"t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P",
		"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "_", "!", "@",
		"#", "$", "%", "&", "*", "(", ")", "-", "_", "+", "=", "[", "{", "}", "]", "/", "?", ";", ":", ".", ",", "<", ">",
		"|"}
	funcListToRunner = make(map[string]func())
	funcListToCacheRunner = make(map[string]func())
	funcListPriorityToRunner = make(map[string]func())
	tickerStart()
}

func getUId() string {
	var UId = ""
	for i := 0; i != kUIdSize; i += 1 {
		UId += kUIdCharList[rand.Intn(len(kUIdCharList)-1)]
	}

	return UId
}

func tickerStart() {
	tickerCache = time.NewTicker(time.Second / time.Duration(fpsCache))
	ticker = time.NewTicker(time.Second / time.Duration(fps))
	slipFrameTimeAlarm = time.Second / time.Duration(fps)
	go func() { tickerRunner() }()
}

func tickerRunner() {
	defer func() { tickerStart() }()

	for {
		select {
		case <-tickerCache.C:

			if stopTicker == true {
				stopTicker = false
				return
			}

			for _, runnerFunc := range funcListToCacheRunner {
				if runnerFunc != nil {
					runnerFunc()
				}
			}

		case <-ticker.C:

			if slipFrame != 0 {
				slipFrame -= 1
				continue
			}

			if stopTicker == true {
				stopTicker = false
				return
			}

			start := time.Now()

			for _, runnerFunc := range funcListPriorityToRunner {
				if runnerFunc != nil {
					runnerFunc()
				}
			}

			for _, runnerFunc := range funcListToRunner {
				if runnerFunc != nil {
					runnerFunc()
				}
			}

			elapsed := time.Since(start)
			if elapsed > slipFrameTimeAlarm {
				slipFrame = 2
			}
		}
	}
}
