package fps

import (
	"math/rand"
	"time"
)

var kUIdCharList []string

const kUIdSize = 10

var fps = 60
var done = make(chan struct{}, 0)
var ticker *time.Ticker
var funcListToRunner map[string]func()

func Set(value int) {
	fps = value
	<-done
}

func Get() int {
	return fps
}

func AddRunner(runnerFunc func()) string {
	UId := getUId()
	funcListToRunner[UId] = runnerFunc
	return UId
}

func DeleteRunner(UId string) {
	delete(funcListToRunner, UId)
}

func init() {
	kUIdCharList = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s",
		"t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P",
		"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "_", "!", "@",
		"#", "$", "%", "&", "*", "(", ")", "-", "_", "+", "=", "[", "{", "}", "]", "/", "?", ";", ":", ".", ",", "<", ">",
		"|"}
	funcListToRunner = make(map[string]func())
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
	ticker = time.NewTicker(time.Second / time.Duration(fps))
	go func() { tickerRunner() }()
}

func tickerRunner() {
	for {
		select {
		case <-ticker.C:
			for _, runnerFunc := range funcListToRunner {
				if runnerFunc != nil {
					go runnerFunc()
				}
			}

		case <-done:
			go func() { tickerStart() }()
			return
		}
	}
}
