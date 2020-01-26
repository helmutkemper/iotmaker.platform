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
	dragId      string
	isDraggable bool
	isMouseDown bool
	dragMode    DragMode
	xDelta      float64
	yDelta      float64
}

type Sprite struct {
	Stage                                iStage.IStage
	Id                                   string
	Platform                             iotmakerPlatformIDraw.IDraw
	ScratchPad                           iotmakerPlatformIDraw.IDraw
	Dimensions                           dimensions.Dimensions
	OutBoxDimensions                     dimensions.Dimensions
	Ink                                  ink.Ink
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

func (el *Sprite) SetDragMode(mode DragMode) {
	el.dragMode = mode
}

func (el *Sprite) SetPlatform(platform iotmakerPlatformIDraw.IDraw) {
	el.Platform = platform
}

func (el *Sprite) GetPlatform() iotmakerPlatformIDraw.IDraw {
	return el.Platform
}

func (el *Sprite) DragStart() {
	el.setDraggable(true)
}

func (el *Sprite) DragStop() {
	el.setDraggable(false)
}

func (el *Sprite) Move(x, y float64) {
	el.OutBoxDimensions.X = x + float64(el.MovieDeltaX)
	el.Dimensions.X = x + float64(el.MovieDeltaX)
	el.OutBoxDimensions.Y = y + float64(el.MovieDeltaY)
	el.Dimensions.Y = y + float64(el.MovieDeltaY)
}

func (el *Sprite) dragOnMouseMove() {
	var x float64
	var y float64

	select {
	case coordinate := <-el.mouseChannelOnMoveEvent:
		if coordinate.HasInit == false {
			return
		}

		x = float64(coordinate.X)
		y = float64(coordinate.Y)

		if el.dragMode == KDragModeAlways {
			el.Move(x, y)
		} else if el.dragMode == KDragModeDesktop && el.isMouseDown == true {
			el.Move(x-el.xDelta, y-el.yDelta)
		}

	case coordinate := <-el.mouseChannelOnClickEvent:

		x = float64(coordinate.X)
		y = float64(coordinate.Y)

		if el.dragMode == KDragModeMobile {
			el.Move(x-el.xDelta, y-el.yDelta)
		}

	case coordinate := <-el.mouseChannelOnDownEvent:

		x = float64(coordinate.X)
		y = float64(coordinate.Y)

		if x >= el.OutBoxDimensions.X &&
			x <= el.OutBoxDimensions.X+el.OutBoxDimensions.Width &&
			y >= el.OutBoxDimensions.Y &&
			y <= el.OutBoxDimensions.Y+el.OutBoxDimensions.Height {

			el.isMouseDown = true
			el.xDelta = x - el.Dimensions.X
			el.yDelta = y - el.Dimensions.Y
		}

	case <-el.mouseChannelOnUpEvent:

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

func (el *Sprite) GetCollisionBox(xEvent, yEvent float64) bool {
	return el.OutBoxDimensions.X <= xEvent && el.OutBoxDimensions.X+el.OutBoxDimensions.Width >= xEvent &&
		el.OutBoxDimensions.Y <= yEvent && el.OutBoxDimensions.Y+el.OutBoxDimensions.Height >= yEvent
}
