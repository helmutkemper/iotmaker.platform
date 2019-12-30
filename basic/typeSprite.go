package basic

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	"github.com/helmutkemper/iotmaker.platform/abstractType/genericTypes"
	platformMouse "github.com/helmutkemper/iotmaker.platform/mouse"
)

type Drag struct {
	IsDraggable bool
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

func (el *Sprite) SetDraggable(enable bool) {
	el.IsDraggable = enable
	if enable == true {
		platformMouse.AddFunctionPointer("size", el.GetCollisionBox, func(x, y float64, collision bool) {
			//if collision == true {
			el.OutBoxDimensions.X = x - el.OutBoxDimensions.Width/2
			el.Dimensions.X = x - el.Dimensions.Width/2
			el.OutBoxDimensions.Y = y - el.OutBoxDimensions.Height/2
			el.Dimensions.Y = y - el.Dimensions.Height/2
			//}
		})
	}
}

func (el *Sprite) GetCollisionBox(xEvent, yEvent float64) bool {
	return el.OutBoxDimensions.X <= xEvent && el.OutBoxDimensions.X+el.OutBoxDimensions.Width >= xEvent &&
		el.OutBoxDimensions.Y <= yEvent && el.OutBoxDimensions.Y+el.OutBoxDimensions.Height >= yEvent
}
