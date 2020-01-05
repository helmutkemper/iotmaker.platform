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
}

var dragModeString = [...]string{
	"desktop",
	"mobile",
}

const (
	KDragModeDesktop DragMode = iota
	KDragModeMobile
)

type Drag struct {
	dragId      string
	isDraggable bool
	dragMode    DragMode
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

func (el *Sprite) SetDraggableToDesktop() {
	platformMouse.AddFunctionPointer(el.Id+"DraggableImage", el.GetCollisionBox, func(x, y float64, collision bool) {
		el.OutBoxDimensions.X = x - el.OutBoxDimensions.Width/2
		el.Dimensions.X = x - el.Dimensions.Width/2
		el.OutBoxDimensions.Y = y - el.OutBoxDimensions.Height/2
		el.Dimensions.Y = y - el.Dimensions.Height/2
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
