package basic

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/genericTypes"
	platformMouse "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/mouse"
)

type Drag struct {
	isDraggable bool
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

func (el *Sprite) setDraggable(enable bool) {
	id := el.Id + "DraggableImage"
	el.isDraggable = enable
	if enable == true {
		platformMouse.AddFunctionPointer(id, el.GetCollisionBox, func(x, y float64, collision bool) {
			el.OutBoxDimensions.X = x - el.OutBoxDimensions.Width/2
			el.Dimensions.X = x - el.Dimensions.Width/2
			el.OutBoxDimensions.Y = y - el.OutBoxDimensions.Height/2
			el.Dimensions.Y = y - el.Dimensions.Height/2
		})
	} else {
		platformMouse.RemoveFunctionPointer(id)
	}
}

func (el *Sprite) GetCollisionBox(xEvent, yEvent float64) bool {
	return el.OutBoxDimensions.X <= xEvent && el.OutBoxDimensions.X+el.OutBoxDimensions.Width >= xEvent &&
		el.OutBoxDimensions.Y <= yEvent && el.OutBoxDimensions.Y+el.OutBoxDimensions.Height >= yEvent
}
