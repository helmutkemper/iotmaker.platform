package engine

// fixme: projeto a parte?
type IEngine interface {
	Init()
	SetSleepFrame(value int)
	GetSleepFrame() int
	Set(value int)
	Get() int
	AddCursorDrawFunc(runnerFunc func()) string
	RemoveCursorDrawFunc(id string)
	AddToHighLatency(runnerFunc func()) string
	DeleteFromHighLatency(UId string)
	AddToSystem(runnerFunc func()) string
	DeleteFromSystem(UId string)
	AddToAfterSystem(runnerFunc func()) string
	DeleteFromAfterSystem(UId string)
	AddToCalculate(runnerFunc func()) string
	DeleteFromCalculate(UId string)
	AddToDraw(runnerFunc func()) string
	DeleteFromDraw(UId string)
	SetZIndex(UId string, index int) int
	SetAsLastFunctionToRun(UId string) int
	ToBack(UId string) int
	GetZIndex(UId string) int
}
