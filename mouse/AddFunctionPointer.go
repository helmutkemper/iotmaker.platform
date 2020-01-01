package mouse

var listTestEventOverFunctions map[string][]PointerCollisionFunction
var listOnEventFunctions map[string][]PointerEventFunction

type platformCursorPointer float64

func (el *platformCursorPointer) Set(value platformCursorPointer) {
	*el = value
}

const (
	KPlatformWebBrowser platformCursorPointer = iota
)

func ManagerMouseMove(x, y float64) {
	var isOver bool

	for id, listFuncTestOver := range listTestEventOverFunctions {

		for k, funcTestOver := range listFuncTestOver {

			if funcTestOver != nil {

				isOver = funcTestOver(x, y)
				listOnEventFunctions[id][k](x, y, isOver)

			}
		}
	}
}

func AddFunctionPointer(id string, collisionFunction PointerCollisionFunction, positiveEventFunction PointerEventFunction) {
	if len(listTestEventOverFunctions) == 0 {
		listTestEventOverFunctions = make(map[string][]PointerCollisionFunction)
		listOnEventFunctions = make(map[string][]PointerEventFunction)
	}

	if len(listTestEventOverFunctions[id]) == 0 {
		listTestEventOverFunctions[id] = make([]PointerCollisionFunction, 0)
		listOnEventFunctions[id] = make([]PointerEventFunction, 0)
	}

	listTestEventOverFunctions[id] = append(listTestEventOverFunctions[id], collisionFunction)
	listOnEventFunctions[id] = append(listOnEventFunctions[id], positiveEventFunction)
}

func RemoveFunctionPointer(id string) {
	delete(listTestEventOverFunctions, id)
	delete(listOnEventFunctions, id)
}

func AddFunctionPointerList(id string, collisionFunction []PointerCollisionFunction, positiveEventFunction []PointerEventFunction) {
	for key := range collisionFunction {
		AddFunctionPointer(id, collisionFunction[key], positiveEventFunction[key])
	}
}

func DeleteFunctionPointer(id string) {
	delete(listTestEventOverFunctions, id)
	delete(listOnEventFunctions, id)
}
