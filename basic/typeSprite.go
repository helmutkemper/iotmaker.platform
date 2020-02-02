package basic

import (
	"fmt"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.interfaces/iStage"
	iotmakerPlatformIDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/dimensions"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/ink"
	platformMouse "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/mouse"
	"strings"
)

type DragMode int

func (el DragMode) String() string {
	return dragModeString[el]
}

func (el DragMode) FromString(value string) DragMode {
	return dragModeFromStringMap[strings.ToLower(value)]
}

var dragModeFromStringMap = map[string]DragMode{
	"desktop": KDragModeDesktop,
	"mobile":  KDragModeMobile,
	"always":  KDragModeAlways,
}

var dragModeString = [...]string{
	"desktop",
	"mobile",
	"always",
}

const (
	KDragModeDesktop DragMode = iota
	KDragModeMobile
	KDragModeAlways
)

type Drag struct {
	dragId       string
	isDraggable  bool
	isMouseDown  bool
	dragMode     DragMode
	xDelta       int
	yDelta       int
	onMoveXStart int
	onMoveYStart int
	onMoveXDelta int
	onMoveYDelta int

	onStartFun func(int, int)
	onEndFun   func(int, int)
}

type Sprite struct {
	Stage                                iStage.IStage
	Id                                   string
	Platform                             iotmakerPlatformIDraw.IDraw
	ScratchPad                           iotmakerPlatformIDraw.IDraw
	Dimensions                           dimensions.Dimensions
	OutBoxDimensions                     dimensions.Dimensions
	Ink                                  ink.Interface
	prepareShadowFilterFunctionPointer   func(iotmakerPlatformIDraw.ICanvasShadow)
	prepareGradientFilterFunctionPointer func(iotmakerPlatformIDraw.ICanvasGradient)
	Drag

	MovieDeltaX int
	MovieDeltaY int

	idMouseChannelOnMoveEvent  int
	idMouseChannelOnDownEvent  int
	idMouseChannelOnUpEvent    int
	idMouseChannelOnClickEvent int

	mouseChannelOnMoveEvent  chan platformMouse.Coordinate
	mouseChannelOnDownEvent  chan platformMouse.Coordinate
	mouseChannelOnUpEvent    chan platformMouse.Coordinate
	mouseChannelOnClickEvent chan platformMouse.Coordinate
}

// en: Get a unique element id
//
// pt_br: Retorna o id único do elemento
func (el *Sprite) GetId() string {
	return el.Id
}

// en: Draw function to re draw element at every frame
//     note: Shadowing function Draw must be implemented into parent struct
//
// pt_br: Função de desenho para redesenhar o elemento a cada frame
//     nota: Função a ser implementada no struct pai
func (el *Sprite) Draw() {

}

func (el *Sprite) SetOnDragStartFunc(f func(int, int)) {
	el.onStartFun = f
}

func (el *Sprite) SetOnDragEndFunc(f func(int, int)) {
	el.onEndFun = f
}

func (el *Sprite) SetDragMode(mode DragMode) {
	el.dragMode = mode
}

func (el *Sprite) SetPlatform(platform iotmakerPlatformIDraw.IDraw) {
	el.Platform = platform
}

func (el *Sprite) GetPlatform() iotmakerPlatformIDraw.IDraw {
	return el.Platform
}

func (el *Sprite) DragIsDragging() bool {
	return el.isMouseDown
}

func (el *Sprite) DragStart() {
	el.setDraggable(true)
}

func (el *Sprite) DragStop() {
	el.setDraggable(false)
}

func (el *Sprite) GetDragDelta() (int, int) {
	return el.onMoveXDelta, el.onMoveYDelta
}

func (el *Sprite) Move(x, y int) {
	el.OutBoxDimensions.X = x + el.MovieDeltaX
	el.Dimensions.X = x + el.MovieDeltaX
	el.OutBoxDimensions.Y = y + el.MovieDeltaY
	el.Dimensions.Y = y + el.MovieDeltaY
}

func (el *Sprite) MoveX(x int) {
	el.OutBoxDimensions.X = x + el.MovieDeltaX
	el.Dimensions.X = x + el.MovieDeltaX
}

func (el *Sprite) MoveY(x, y int) {
	el.OutBoxDimensions.Y = y + el.MovieDeltaY
	el.Dimensions.Y = y + el.MovieDeltaY
}

func (el *Sprite) ColorFiltersStart(platform iotmakerPlatformIDraw.IDraw) {
	if el.Ink == nil {
		return
	}

	el.Ink.ShadowPrepareFilter(platform)
	el.Ink.GradientPrepareFilter(platform)
}

func (el *Sprite) dragOnMouseMove() {
	var x int
	var y int

	select {
	case coordinate := <-el.mouseChannelOnMoveEvent:
		if coordinate.HasInit == false {
			return
		}

		x = coordinate.X
		y = coordinate.Y

		if el.dragMode == KDragModeAlways {
			el.Move(x, y)
		} else if el.dragMode == KDragModeDesktop && el.isMouseDown == true {
			el.Move(x-el.xDelta, y-el.yDelta)
		}

	case coordinate := <-el.mouseChannelOnClickEvent:

		x = coordinate.X
		y = coordinate.Y

		if el.dragMode == KDragModeMobile {
			el.Move(x-el.xDelta, y-el.yDelta)
		}

	case coordinate := <-el.mouseChannelOnDownEvent:

		x = coordinate.X
		y = coordinate.Y

		if el.Stage.IsDraggable(x, y) != el.Id {
			return
		}

		if el.GetCollisionBox(x, y) {

			el.isMouseDown = true
			el.xDelta = x - el.Dimensions.X
			el.yDelta = y - el.Dimensions.Y

			el.onMoveXStart = x
			el.onMoveYStart = y

			if el.onStartFun != nil {
				el.onStartFun(x, y)
			}
		}

	case coordinate := <-el.mouseChannelOnUpEvent:

		x = coordinate.X
		y = coordinate.Y

		el.onMoveXDelta = x - el.onMoveXStart
		el.onMoveYDelta = y - el.onMoveYStart

		if el.isMouseDown == true && el.onEndFun != nil {
			el.isMouseDown = false
			el.onEndFun(x, y)
		}

		el.isMouseDown = false

	default:
	}
}

func (el *Sprite) SetDraggableToDesktop() {

	el.mouseChannelOnMoveEvent = make(chan platformMouse.Coordinate, 1)
	el.mouseChannelOnDownEvent = make(chan platformMouse.Coordinate, 1)
	el.mouseChannelOnUpEvent = make(chan platformMouse.Coordinate, 1)
	el.mouseChannelOnClickEvent = make(chan platformMouse.Coordinate, 1)

	el.idMouseChannelOnMoveEvent = platformMouse.Move.Add(el.mouseChannelOnMoveEvent)
	el.idMouseChannelOnDownEvent = platformMouse.Down.Add(el.mouseChannelOnDownEvent)
	el.idMouseChannelOnUpEvent = platformMouse.Up.Add(el.mouseChannelOnUpEvent)
	el.idMouseChannelOnClickEvent = platformMouse.Click.Add(el.mouseChannelOnClickEvent)

	engine := el.Stage.GetEngine()
	engine.DrawAddToFunctions(el.dragOnMouseMove)
}

func (el *Sprite) RemoveDraggableToDesktop() {
	var err error
	err = platformMouse.Move.Remove(el.idMouseChannelOnMoveEvent)
	if err != nil {
		fmt.Printf("typeSprite.RemoveDraggableToDesktop().Move.error: %v\n", err)
	}

	err = platformMouse.Down.Remove(el.idMouseChannelOnDownEvent)
	if err != nil {
		fmt.Printf("typeSprite.RemoveDraggableToDesktop().Down.error: %v\n", err)
	}

	err = platformMouse.Up.Remove(el.idMouseChannelOnUpEvent)
	if err != nil {
		fmt.Printf("typeSprite.RemoveDraggableToDesktop().Up.error: %v\n", err)
	}

	err = platformMouse.Click.Remove(el.idMouseChannelOnClickEvent)
	if err != nil {
		fmt.Printf("typeSprite.RemoveDraggableToDesktop().Click.error: %v\n", err)
	}

	platformMouse.RemoveFunctionPointer(el.Id + "DraggableImage")
}

func (el *Sprite) setDraggable(enable bool) {
	el.isDraggable = enable
	if enable == true {
		el.SetDraggableToDesktop()
	} else {
		el.RemoveDraggableToDesktop()
	}
}

// en: Get an information about (x, y) is in element boxe
//
// pt_br: Retorna a informação se (x, y) está dentro da caxa do elemento
func (el *Sprite) GetCollisionBox(xEvent, yEvent int) bool {
	return el.OutBoxDimensions.X <= xEvent && el.OutBoxDimensions.X+el.OutBoxDimensions.Width >= xEvent &&
		el.OutBoxDimensions.Y <= yEvent && el.OutBoxDimensions.Y+el.OutBoxDimensions.Height >= yEvent
}
