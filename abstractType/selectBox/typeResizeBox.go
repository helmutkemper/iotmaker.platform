package selectBox

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	"github.com/helmutkemper/iotmaker.platform/abstractType/genericTypes"
	"github.com/helmutkemper/iotmaker.platform/independentDraw"
	"image/color"
)

type ResizeBoxes struct {
	Platform   iotmaker_platform_IDraw.IDraw
	ScratchPad iotmaker_platform_IDraw.IDraw
	Dimensions genericTypes.Dimensions

	FatherOutBoxDimensions genericTypes.Dimensions

	imageDataMethod           genericTypes.ImageDataCaptureMethod
	imageDataComplete         map[int]map[int]color.RGBA
	imageDataAlphaChannel     map[int]map[int]uint8
	imageDataBooleanCollision map[int]map[int]bool

	prepareShadowFilterFunctionPointer   func(iotmaker_platform_IDraw.ICanvasShadow)
	prepareGradientFilterFunctionPointer func(iotmaker_platform_IDraw.ICanvasGradient)

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

func (el *ResizeBoxes) Create() {

	//   a  _       _       _  c
	//     |_|-----|_|-----|_|
	//      |       b       |
	//      |   +-------+   |
	//      |   |       |   |
	//     |_|h |       | d|_|
	//      |   |       |   |
	//      |   +-------+   |
	//      |       f       |
	//     |_|-----|_|-----|_|
	//   g                     e

	el.Platform.SetLineWidth(1)
	independentDraw.DrawBoxWithRoundedCornersIntoThePath(el.Platform, el.FatherOutBoxDimensions.X, el.FatherOutBoxDimensions.Y, el.FatherOutBoxDimensions.Width, el.FatherOutBoxDimensions.Height, 0)
	el.Platform.Stroke()

	el.Platform.SetLineWidth(1)
	el.createCornerA()
	el.Platform.Stroke()

	el.Platform.SetLineWidth(1)
	el.createCornerB()
	el.Platform.Stroke()

	el.Platform.SetLineWidth(1)
	el.createCornerC()
	el.Platform.Stroke()

	el.Platform.SetLineWidth(1)
	el.createCornerD()
	el.Platform.Stroke()

	el.Platform.SetLineWidth(1)
	el.createCornerE()
	el.Platform.Stroke()

	el.Platform.SetLineWidth(1)
	el.createCornerF()
	el.Platform.Stroke()

	el.Platform.SetLineWidth(1)
	el.createCornerG()
	el.Platform.Stroke()

	el.Platform.SetLineWidth(1)
	el.createCornerH()
	el.Platform.Stroke()

}

func (el *ResizeBoxes) createCornerA() {
	// a corner
	x := el.FatherOutBoxDimensions.X - el.Dimensions.Width
	y := el.FatherOutBoxDimensions.Y - el.Dimensions.Height

	// a corner space
	x -= el.Dimensions.X
	y -= el.Dimensions.Y

	// a corner create
	independentDraw.DrawBoxWithRoundedCornersIntoThePath(el.Platform, x, y, el.Dimensions.Width, el.Dimensions.Height, el.Dimensions.Border)
}

func (el *ResizeBoxes) createCornerB() {
	// b corner
	el.boxBX = el.FatherOutBoxDimensions.X + el.FatherOutBoxDimensions.Width/2
	el.boxBY = el.FatherOutBoxDimensions.Y - el.Dimensions.Height

	// b corner space
	el.boxBY -= el.Dimensions.Y

	// b corner create
	independentDraw.DrawBoxWithRoundedCornersIntoThePath(el.Platform, el.boxBX, el.boxBY, el.Dimensions.Width, el.Dimensions.Height, el.Dimensions.Border)
}

func (el *ResizeBoxes) createCornerC() {
	// c corner
	el.boxCX = el.FatherOutBoxDimensions.X + el.FatherOutBoxDimensions.Width
	el.boxCY = el.FatherOutBoxDimensions.Y - el.Dimensions.Height

	// c corner space
	el.boxCX += el.Dimensions.X
	el.boxCY -= el.Dimensions.Y

	// c corner create
	independentDraw.DrawBoxWithRoundedCornersIntoThePath(el.Platform, el.boxCX, el.boxCY, el.Dimensions.Width, el.Dimensions.Height, el.Dimensions.Border)
}

func (el *ResizeBoxes) createCornerD() {
	// d corner
	el.boxDX = el.FatherOutBoxDimensions.X + el.FatherOutBoxDimensions.Width
	el.boxDY = el.FatherOutBoxDimensions.Y + el.FatherOutBoxDimensions.Height/2

	// d corner space
	el.boxDX += el.Dimensions.X

	// d corner create
	independentDraw.DrawBoxWithRoundedCornersIntoThePath(el.Platform, el.boxDX, el.boxDY, el.Dimensions.Width, el.Dimensions.Height, el.Dimensions.Border)
}

func (el *ResizeBoxes) createCornerE() {
	// e corner
	el.boxEX = el.FatherOutBoxDimensions.X + el.FatherOutBoxDimensions.Width
	el.boxEY = el.FatherOutBoxDimensions.Y + el.FatherOutBoxDimensions.Height

	// e corner space
	el.boxEX += el.Dimensions.X
	el.boxEY += el.Dimensions.Y

	// e corner create
	independentDraw.DrawBoxWithRoundedCornersIntoThePath(el.Platform, el.boxEX, el.boxEY, el.Dimensions.Width, el.Dimensions.Height, el.Dimensions.Border)
}

func (el *ResizeBoxes) createCornerF() {
	// f corner
	el.boxFX = el.FatherOutBoxDimensions.X + el.FatherOutBoxDimensions.Width/2
	el.boxFY = el.FatherOutBoxDimensions.Y + el.FatherOutBoxDimensions.Height

	// f corner space
	el.boxFY += el.Dimensions.Y

	// f corner create
	independentDraw.DrawBoxWithRoundedCornersIntoThePath(el.Platform, el.boxFX, el.boxFY, el.Dimensions.Width, el.Dimensions.Height, el.Dimensions.Border)
}

func (el *ResizeBoxes) createCornerG() {
	// g corner
	el.boxGX = el.FatherOutBoxDimensions.X - el.Dimensions.Width
	el.boxGY = el.FatherOutBoxDimensions.Y + el.FatherOutBoxDimensions.Height

	// g corner space
	el.boxGX -= el.Dimensions.X
	el.boxGY += el.Dimensions.Y

	// g corner create
	independentDraw.DrawBoxWithRoundedCornersIntoThePath(el.Platform, el.boxGX, el.boxGY, el.Dimensions.Width, el.Dimensions.Height, el.Dimensions.Border)
}

func (el *ResizeBoxes) createCornerH() {
	// h corner
	el.boxHX = el.FatherOutBoxDimensions.X - el.Dimensions.Width
	el.boxHY = el.FatherOutBoxDimensions.Y + el.FatherOutBoxDimensions.Height/2

	// h corner space
	el.boxHX -= el.Dimensions.X

	// h corner create
	independentDraw.DrawBoxWithRoundedCornersIntoThePath(el.Platform, el.boxHX, el.boxHY, el.Dimensions.Width, el.Dimensions.Height, el.Dimensions.Border)
}
