package basic

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/genericTypes"
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
	Id                                   string
	Platform                             iotmaker_platform_IDraw.IDraw
	ScratchPad                           iotmaker_platform_IDraw.IDraw
	Dimensions                           genericTypes.Dimensions
	OutBoxDimensions                     genericTypes.Dimensions
	Ink                                  genericTypes.Ink
	prepareShadowFilterFunctionPointer   func(iotmaker_platform_IDraw.ICanvasShadow)
	prepareGradientFilterFunctionPointer func(iotmaker_platform_IDraw.ICanvasGradient)
	Drag
}

func (el *Sprite) SetDragMode(mode DragMode) {
	el.dragMode = mode
}

func (el *Sprite) SetPlatform(platform iotmaker_platform_IDraw.IDraw) {
	el.Platform = platform
}

func (el *Sprite) GetPlatform() iotmaker_platform_IDraw.IDraw {
	return el.Platform
}

func (el *Sprite) DragStart() {
	el.setDraggable(true)
}

func (el *Sprite) DragStop() {
	el.setDraggable(false)
}

func (el *Sprite) Move(x, y float64) {
	el.OutBoxDimensions.X = x
	el.Dimensions.X = x
	el.OutBoxDimensions.Y = y
	el.Dimensions.Y = y
}

func (el *Sprite) SetDraggableToDesktop() {
	platformMouse.AddFunctionPointer(el.Id+"DraggableImage", el.GetCollisionBox, func(x, y float64, collision bool, event platformMouse.EventMouse) {

		switch event {
		case platformMouse.KMouseDown:

			if x >= el.OutBoxDimensions.X &&
				x <= el.OutBoxDimensions.X+el.OutBoxDimensions.Width &&
				y >= el.OutBoxDimensions.Y &&
				y <= el.OutBoxDimensions.Y+el.OutBoxDimensions.Height {

				el.isMouseDown = true
				el.xDelta = x - el.Dimensions.X
				el.yDelta = y - el.Dimensions.Y
			}

		case platformMouse.KMouseUp:
			el.isMouseDown = false
		}

		if el.dragMode == KDragModeAlways {
			el.Move(x, y)
		} else if el.dragMode == KDragModeDesktop && el.isMouseDown == true {
			el.Move(x-el.xDelta, y-el.yDelta)
		} else if el.dragMode == KDragModeMobile && event == platformMouse.KClick {
			el.Move(x-el.xDelta, y-el.yDelta)
		}

	})
}

func (el *Sprite) RemoveDraggableToDesktop() {
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
