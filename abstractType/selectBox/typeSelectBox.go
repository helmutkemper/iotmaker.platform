package selectBox

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	"github.com/helmutkemper/iotmaker.platform/abstractType/basicBox"
	"github.com/helmutkemper/iotmaker.platform/abstractType/genericTypes"
	"image/color"
)

type ResizeBoxes struct {
	Platform   iotmaker_platform_IDraw.IDraw
	ScratchPad iotmaker_platform_IDraw.IDraw
	Id         string
	Dimensions genericTypes.Dimensions
	Ink        genericTypes.Ink

	fatherOutBoxDimensions genericTypes.Dimensions

	imageDataMethod           genericTypes.ImageDataCaptureMethod
	imageDataComplete         map[int]map[int]color.RGBA
	imageDataAlphaChannel     map[int]map[int]uint8
	imageDataBooleanCollision map[int]map[int]bool

	prepareShadowFilterFunctionPointer   func(iotmaker_platform_IDraw.ICanvasShadow)
	prepareGradientFilterFunctionPointer func(iotmaker_platform_IDraw.ICanvasGradient)

	//
	//   a          b           c
	//      +-------+--------+
	//      |                |
	//      |                |
	//      |                |
	//   h  +                +  d
	//      |                |
	//      |                |
	//      |                |
	//      +-------+--------+
	//   g          f           e
	//

	BoxA basicBox.BasicBox
	BoxB basicBox.BasicBox
	BoxC basicBox.BasicBox
	BoxD basicBox.BasicBox
	BoxE basicBox.BasicBox
	BoxF basicBox.BasicBox
	BoxG basicBox.BasicBox
	BoxH basicBox.BasicBox

	boxAX int
	boxAY int

	boxBX int
	boxBY int

	boxCX int
	boxCY int

	boxDX int
	boxDY int

	boxEX int
	boxEY int

	boxFX int
	boxFY int

	boxGX int
	boxGY int

	boxHX int
	boxHY int
}

func (el *ResizeBoxes) calculateCoordinates() {
	el.boxAX = el.fatherOutBoxDimensions.X
	el.boxAY = el.fatherOutBoxDimensions.Y

	el.boxBX = el.fatherOutBoxDimensions.X + el.fatherOutBoxDimensions.Width/2
	el.boxBY = el.fatherOutBoxDimensions.Y

	el.boxCX = el.fatherOutBoxDimensions.X + el.fatherOutBoxDimensions.Width
	el.boxCY = el.fatherOutBoxDimensions.Y

	el.boxDX = el.fatherOutBoxDimensions.X + el.fatherOutBoxDimensions.Width
	el.boxDY = el.fatherOutBoxDimensions.Y + el.fatherOutBoxDimensions.Height/2

	el.boxEX = el.fatherOutBoxDimensions.X + el.fatherOutBoxDimensions.Width
	el.boxEY = el.fatherOutBoxDimensions.Y + el.fatherOutBoxDimensions.Height

	el.boxFX = el.fatherOutBoxDimensions.X + el.fatherOutBoxDimensions.Width/2
	el.boxFY = el.fatherOutBoxDimensions.Y + el.fatherOutBoxDimensions.Height

	el.boxGX = el.fatherOutBoxDimensions.X
	el.boxGY = el.fatherOutBoxDimensions.Y + el.fatherOutBoxDimensions.Height

	el.boxHX = el.fatherOutBoxDimensions.X
	el.boxHY = el.fatherOutBoxDimensions.Y + el.fatherOutBoxDimensions.Height/2
}
