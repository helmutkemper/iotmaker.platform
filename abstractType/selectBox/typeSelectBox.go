package selectBox

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
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

	Density  interface{}
	IDensity iotmaker_platform_coordinate.IDensity

	fatherOutBoxDimensions genericTypes.Dimensions

	imageDataMethod           genericTypes.ImageDataCaptureMethod
	imageDataComplete         map[int]map[int]color.RGBA
	imageDataAlphaChannel     map[int]map[int]uint8
	imageDataBooleanCollision map[int]map[int]bool

	prepareShadowFilterFunctionPointer   func(iotmaker_platform_IDraw.ICanvasShadow)
	prepareGradientFilterFunctionPointer func(iotmaker_platform_IDraw.ICanvasGradient)

	BoxA *basicBox.BasicBox
	BoxB *basicBox.BasicBox
	BoxC *basicBox.BasicBox
	BoxD *basicBox.BasicBox
	BoxE *basicBox.BasicBox
	BoxF *basicBox.BasicBox
	BoxG *basicBox.BasicBox
	BoxH *basicBox.BasicBox

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

	el.createCornerA()
	el.createCornerB()
	el.createCornerC()
	el.createCornerD()
	el.createCornerE()
	el.createCornerF()
	el.createCornerG()
	el.createCornerH()
}

func (el *ResizeBoxes) createCornerA() {
	// a corner
	el.boxAX = el.fatherOutBoxDimensions.X
	el.boxAY = el.fatherOutBoxDimensions.Y

	// a corner space
	el.boxAX -= el.Dimensions.X
	el.boxAY -= el.Dimensions.Y

	// a corner create
	el.BoxA = basicBox.NewBasicBox(el.Platform, el.ScratchPad, el.Id+"_cornerA", el.boxAX, el.boxAY, el.Dimensions.Width, el.Dimensions.Height, el.Dimensions.Border, 1, nil, nil, el.Density, el.IDensity)
}

func (el *ResizeBoxes) createCornerB() {
	// b corner
	el.boxBX = el.fatherOutBoxDimensions.X + el.fatherOutBoxDimensions.Width/2
	el.boxBY = el.fatherOutBoxDimensions.Y

	// b corner space
	el.boxBY -= el.Dimensions.Y

	// b corner create
	el.BoxB = basicBox.NewBasicBox(el.Platform, el.ScratchPad, el.Id+"_cornerB", el.boxBX, el.boxBY, el.Dimensions.Width, el.Dimensions.Height, el.Dimensions.Border, 1, nil, nil, el.Density, el.IDensity)
}

func (el *ResizeBoxes) createCornerC() {
	// c corner
	el.boxCX = el.fatherOutBoxDimensions.X + el.fatherOutBoxDimensions.Width
	el.boxCY = el.fatherOutBoxDimensions.Y

	// c corner space
	el.boxCX += el.Dimensions.X
	el.boxCY += el.Dimensions.Y

	// c corner create
	el.BoxC = basicBox.NewBasicBox(el.Platform, el.ScratchPad, el.Id+"_cornerC", el.boxCX, el.boxCY, el.Dimensions.Width, el.Dimensions.Height, el.Dimensions.Border, 1, nil, nil, el.Density, el.IDensity)
}

func (el *ResizeBoxes) createCornerD() {
	// d corner
	el.boxDX = el.fatherOutBoxDimensions.X + el.fatherOutBoxDimensions.Width
	el.boxDY = el.fatherOutBoxDimensions.Y + el.fatherOutBoxDimensions.Height/2

	// d corner space
	el.boxDX += el.Dimensions.X

	// d corner create
	el.BoxD = basicBox.NewBasicBox(el.Platform, el.ScratchPad, el.Id+"_cornerD", el.boxDX, el.boxDY, el.Dimensions.Width, el.Dimensions.Height, el.Dimensions.Border, 1, nil, nil, el.Density, el.IDensity)
}

func (el *ResizeBoxes) createCornerE() {
	// e corner
	el.boxEX = el.fatherOutBoxDimensions.X + el.fatherOutBoxDimensions.Width
	el.boxEY = el.fatherOutBoxDimensions.Y + el.fatherOutBoxDimensions.Height

	// e corner space
	el.boxEX += el.Dimensions.X
	el.boxEY += el.Dimensions.Y

	// e corner create
	el.BoxE = basicBox.NewBasicBox(el.Platform, el.ScratchPad, el.Id+"_cornerE", el.boxEX, el.boxEY, el.Dimensions.Width, el.Dimensions.Height, el.Dimensions.Border, 1, nil, nil, el.Density, el.IDensity)
}

func (el *ResizeBoxes) createCornerF() {
	// f corner
	el.boxFX = el.fatherOutBoxDimensions.X + el.fatherOutBoxDimensions.Width/2
	el.boxFY = el.fatherOutBoxDimensions.Y + el.fatherOutBoxDimensions.Height

	// f corner space
	el.boxFY += el.Dimensions.Y

	// f corner create
	el.BoxF = basicBox.NewBasicBox(el.Platform, el.ScratchPad, el.Id+"_cornerF", el.boxFX, el.boxFY, el.Dimensions.Width, el.Dimensions.Height, el.Dimensions.Border, 1, nil, nil, el.Density, el.IDensity)
}

func (el *ResizeBoxes) createCornerG() {
	// g corner
	el.boxGX = el.fatherOutBoxDimensions.X
	el.boxGY = el.fatherOutBoxDimensions.Y + el.fatherOutBoxDimensions.Height

	// g corner space
	el.boxGX -= el.Dimensions.X
	el.boxGY += el.Dimensions.Y

	// g corner create
	el.BoxG = basicBox.NewBasicBox(el.Platform, el.ScratchPad, el.Id+"_cornerG", el.boxGX, el.boxGY, el.Dimensions.Width, el.Dimensions.Height, el.Dimensions.Border, 1, nil, nil, el.Density, el.IDensity)
}

func (el *ResizeBoxes) createCornerH() {
	// h corner
	el.boxHX = el.fatherOutBoxDimensions.X
	el.boxHY = el.fatherOutBoxDimensions.Y + el.fatherOutBoxDimensions.Height/2

	// h corner space
	el.boxHX -= el.Dimensions.X

	// h corner create
	el.BoxH = basicBox.NewBasicBox(el.Platform, el.ScratchPad, el.Id+"_cornerH", el.boxHX, el.boxHY, el.Dimensions.Width, el.Dimensions.Height, el.Dimensions.Border, 1, nil, nil, el.Density, el.IDensity)
}
