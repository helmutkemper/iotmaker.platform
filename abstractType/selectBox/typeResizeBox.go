package selectBox

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	webBrowserMouse "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/mouse"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/genericTypes"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/dimensions"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/independentDraw"
	"image/color"
)

type ResizeBoxes struct {
	Platform   iotmaker_platform_IDraw.IDraw
	ScratchPad iotmaker_platform_IDraw.IDraw

	Dimensions             dimensions.Dimensions
	FatherOutBoxDimensions dimensions.Dimensions
	outBoxDimensions       dimensions.Dimensions

	CornerFillColor  interface{}
	CornerLineWidth  float64
	ContourLineWidth float64
	ContourColor     color.RGBA

	imageDataMethod           genericTypes.ImageDataCaptureMethod
	imageDataComplete         map[int]map[int]color.RGBA
	imageDataAlphaChannel     map[int]map[int]uint8
	imageDataBooleanCollision map[int]map[int]bool

	prepareShadowFilterFunctionPointer   func(iotmaker_platform_IDraw.ICanvasShadow)
	prepareGradientFilterFunctionPointer func(iotmaker_platform_IDraw.ICanvasGradient)

	//MouseFunc     mouse.SetCursorFunc
	MouseGeneric webBrowserMouse.CursorType
	MouseCornerA webBrowserMouse.CursorType
	MouseCornerB webBrowserMouse.CursorType
	MouseCornerC webBrowserMouse.CursorType
	MouseCornerD webBrowserMouse.CursorType
	MouseCornerE webBrowserMouse.CursorType
	MouseCornerF webBrowserMouse.CursorType
	MouseCornerG webBrowserMouse.CursorType
	MouseCornerH webBrowserMouse.CursorType

	boxAX float64
	boxAY float64

	boxBX float64
	boxBY float64

	boxCX float64
	boxCY float64

	boxDX float64
	boxDY float64

	boxEX float64
	boxEY float64

	boxFX float64
	boxFY float64

	boxGX float64
	boxGY float64

	boxHX float64
	boxHY float64
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

	//el.Platform.SetLineWidth(1)
	//independentDraw.DrawBoxWithRoundedCornersIntoThePath(el.Platform, el.FatherOutBoxDimensions.X, el.FatherOutBoxDimensions.Y, el.FatherOutBoxDimensions.Width, el.FatherOutBoxDimensions.Height, 0)
	//el.Platform.Stroke()

	el.createCornerA()
	el.createCornerB()
	el.createCornerC()
	el.createCornerD()
	el.createCornerE()
	el.createCornerF()
	el.createCornerG()
	el.createCornerH()

	el.lineFromCornerAToCornerB()
	el.lineFromCornerBToCornerC()
	el.lineFromCornerCToCornerD()
	el.lineFromCornerDToCornerE()
	el.lineFromCornerEToCornerF()
	el.lineFromCornerFToCornerG()
	el.lineFromCornerGToCornerH()
	el.lineFromCornerHToCornerA()

	el.calculateOutBoxDimensions()
}

func (el *ResizeBoxes) calculateOutBoxDimensions() {
	el.outBoxDimensions.X = int(el.boxAX - el.CornerLineWidth/2)
	el.outBoxDimensions.Y = int(el.boxAY - el.CornerLineWidth/2)
	el.outBoxDimensions.Width = int(el.boxEX) + el.Dimensions.Width - int(el.boxAX+el.CornerLineWidth)
	el.outBoxDimensions.Height = int(el.boxEY) + el.Dimensions.Height - int(el.boxAY+el.CornerLineWidth)

	/*el.Platform.SetLineWidth(1)
	  el.Platform.MoveTo( el.outBoxDimensions.X, el.outBoxDimensions.Y )
	  el.Platform.LineTo( el.outBoxDimensions.X, el.outBoxDimensions.Y+el.outBoxDimensions.Height )
	  el.Platform.LineTo( el.outBoxDimensions.X+el.outBoxDimensions.Width, el.outBoxDimensions.Y+el.outBoxDimensions.Height )
	  el.Platform.LineTo( el.outBoxDimensions.X+el.outBoxDimensions.Width, el.outBoxDimensions.Y )
	  el.Platform.LineTo( el.outBoxDimensions.X, el.outBoxDimensions.Y )
	  el.Platform.Stroke()*/
}

func (el *ResizeBoxes) createCornerA() {
	// a corner
	el.boxAX = float64(el.FatherOutBoxDimensions.X - el.Dimensions.Width)
	el.boxAY = float64(el.FatherOutBoxDimensions.Y - el.Dimensions.Height)

	// a corner space
	el.boxAX -= float64(el.Dimensions.X)
	el.boxAY -= float64(el.Dimensions.Y)

	// a corner create
	el.Platform.SetLineWidth(el.CornerLineWidth)
	independentDraw.DrawBoxWithRoundedCornersIntoThePath(
		el.Platform,
		el.boxAX,
		el.boxAY,
		float64(el.Dimensions.Width),
		float64(el.Dimensions.Height),
		float64(el.Dimensions.Border),
	)
	el.Platform.SetFillStyle(el.CornerFillColor)

	if el.CornerFillColor != nil {
		el.Platform.Fill()
	}

	if el.CornerLineWidth != 0 {
		el.Platform.Stroke()
	}
}

func (el *ResizeBoxes) createCornerB() {
	// b corner
	el.boxBX = float64(el.FatherOutBoxDimensions.X + el.FatherOutBoxDimensions.Width/2)
	el.boxBY = float64(el.FatherOutBoxDimensions.Y - el.Dimensions.Height)

	// b corner space
	el.boxBY -= float64(el.Dimensions.Y)

	// b corner create
	el.Platform.SetLineWidth(el.CornerLineWidth)
	independentDraw.DrawBoxWithRoundedCornersIntoThePath(
		el.Platform,
		el.boxBX,
		el.boxBY,
		float64(el.Dimensions.Width),
		float64(el.Dimensions.Height),
		float64(el.Dimensions.Border),
	)
	el.Platform.SetFillStyle(el.CornerFillColor)

	if el.CornerFillColor != nil {
		el.Platform.Fill()
	}

	if el.CornerLineWidth != 0 {
		el.Platform.Stroke()
	}
}

func (el *ResizeBoxes) createCornerC() {
	// c corner
	el.boxCX = float64(el.FatherOutBoxDimensions.X + el.FatherOutBoxDimensions.Width)
	el.boxCY = float64(el.FatherOutBoxDimensions.Y - el.Dimensions.Height)

	// c corner space
	el.boxCX += float64(el.Dimensions.X)
	el.boxCY -= float64(el.Dimensions.Y)

	// c corner create
	el.Platform.SetLineWidth(el.CornerLineWidth)
	independentDraw.DrawBoxWithRoundedCornersIntoThePath(
		el.Platform,
		el.boxCX,
		el.boxCY,
		float64(el.Dimensions.Width),
		float64(el.Dimensions.Height),
		float64(el.Dimensions.Border),
	)
	el.Platform.SetFillStyle(el.CornerFillColor)

	if el.CornerFillColor != nil {
		el.Platform.Fill()
	}

	if el.CornerLineWidth != 0 {
		el.Platform.Stroke()
	}
}

func (el *ResizeBoxes) createCornerD() {
	// d corner
	el.boxDX = float64(el.FatherOutBoxDimensions.X + el.FatherOutBoxDimensions.Width)
	el.boxDY = float64(el.FatherOutBoxDimensions.Y + el.FatherOutBoxDimensions.Height/2)

	// d corner space
	el.boxDX += float64(el.Dimensions.X)

	// d corner create
	el.Platform.SetLineWidth(el.CornerLineWidth)
	independentDraw.DrawBoxWithRoundedCornersIntoThePath(
		el.Platform,
		el.boxDX,
		el.boxDY,
		float64(el.Dimensions.Width),
		float64(el.Dimensions.Height),
		float64(el.Dimensions.Border),
	)
	el.Platform.SetFillStyle(el.CornerFillColor)

	if el.CornerFillColor != nil {
		el.Platform.Fill()
	}

	if el.CornerLineWidth != 0 {
		el.Platform.Stroke()
	}
}

func (el *ResizeBoxes) createCornerE() {
	// e corner
	el.boxEX = float64(el.FatherOutBoxDimensions.X + el.FatherOutBoxDimensions.Width)
	el.boxEY = float64(el.FatherOutBoxDimensions.Y + el.FatherOutBoxDimensions.Height)

	// e corner space
	el.boxEX += float64(el.Dimensions.X)
	el.boxEY += float64(el.Dimensions.Y)

	// e corner create
	el.Platform.SetLineWidth(el.CornerLineWidth)
	independentDraw.DrawBoxWithRoundedCornersIntoThePath(
		el.Platform,
		el.boxEX,
		el.boxEY,
		float64(el.Dimensions.Width),
		float64(el.Dimensions.Height),
		float64(el.Dimensions.Border),
	)
	el.Platform.SetFillStyle(el.CornerFillColor)

	if el.CornerFillColor != nil {
		el.Platform.Fill()
	}

	if el.CornerLineWidth != 0 {
		el.Platform.Stroke()
	}
}

func (el *ResizeBoxes) createCornerF() {
	// f corner
	el.boxFX = float64(el.FatherOutBoxDimensions.X + el.FatherOutBoxDimensions.Width/2)
	el.boxFY = float64(el.FatherOutBoxDimensions.Y + el.FatherOutBoxDimensions.Height)

	// f corner space
	el.boxFY += float64(el.Dimensions.Y)

	// f corner create
	el.Platform.SetLineWidth(el.CornerLineWidth)
	independentDraw.DrawBoxWithRoundedCornersIntoThePath(
		el.Platform,
		el.boxFX,
		el.boxFY,
		float64(el.Dimensions.Width),
		float64(el.Dimensions.Height),
		float64(el.Dimensions.Border),
	)
	el.Platform.SetFillStyle(el.CornerFillColor)

	if el.CornerFillColor != nil {
		el.Platform.Fill()
	}

	if el.CornerLineWidth != 0 {
		el.Platform.Stroke()
	}
}

func (el *ResizeBoxes) createCornerG() {
	// g corner
	el.boxGX = float64(el.FatherOutBoxDimensions.X - el.Dimensions.Width)
	el.boxGY = float64(el.FatherOutBoxDimensions.Y + el.FatherOutBoxDimensions.Height)

	// g corner space
	el.boxGX -= float64(el.Dimensions.X)
	el.boxGY += float64(el.Dimensions.Y)

	// g corner create
	el.Platform.SetLineWidth(el.CornerLineWidth)
	independentDraw.DrawBoxWithRoundedCornersIntoThePath(
		el.Platform,
		el.boxGX,
		el.boxGY,
		float64(el.Dimensions.Width),
		float64(el.Dimensions.Height),
		float64(el.Dimensions.Border),
	)
	el.Platform.SetFillStyle(el.CornerFillColor)

	if el.CornerFillColor != nil {
		el.Platform.Fill()
	}

	if el.CornerLineWidth != 0 {
		el.Platform.Stroke()
	}
}

func (el *ResizeBoxes) createCornerH() {
	// h corner
	el.boxHX = float64(el.FatherOutBoxDimensions.X - el.Dimensions.Width)
	el.boxHY = float64(el.FatherOutBoxDimensions.Y + el.FatherOutBoxDimensions.Height/2)

	// h corner space
	el.boxHX -= float64(el.Dimensions.X)

	// h corner create
	el.Platform.SetLineWidth(el.CornerLineWidth)
	independentDraw.DrawBoxWithRoundedCornersIntoThePath(
		el.Platform,
		el.boxHX,
		el.boxHY,
		float64(el.Dimensions.Width),
		float64(el.Dimensions.Height),
		float64(el.Dimensions.Border),
	)
	el.Platform.SetFillStyle(el.CornerFillColor)

	if el.CornerFillColor != nil {
		el.Platform.Fill()
	}

	if el.CornerLineWidth != 0 {
		el.Platform.Stroke()
	}
}

func (el *ResizeBoxes) lineFromCornerAToCornerB() {
	x0 := el.boxAX + float64(el.Dimensions.Width)
	y0 := el.boxAY + float64(el.Dimensions.Height/2)

	x1 := el.boxBX

	el.Platform.SetLineWidth(el.ContourLineWidth)
	el.Platform.SetStrokeStyle(el.ContourColor)
	independentDraw.DrawSimpleLineIntoThePath(el.Platform, x0, y0, x1, y0)
	el.Platform.Stroke()
}

func (el *ResizeBoxes) lineFromCornerBToCornerC() {
	x0 := el.boxBX + float64(el.Dimensions.Width)
	y0 := el.boxBY + float64(el.Dimensions.Height/2)

	x1 := el.boxCX

	el.Platform.SetLineWidth(el.ContourLineWidth)
	el.Platform.SetStrokeStyle(el.ContourColor)
	independentDraw.DrawSimpleLineIntoThePath(el.Platform, x0, y0, x1, y0)
	el.Platform.Stroke()
}

func (el *ResizeBoxes) lineFromCornerCToCornerD() {
	x0 := el.boxCX + float64(el.Dimensions.Width/2)
	y0 := el.boxCY + float64(el.Dimensions.Height)

	y1 := el.boxDY

	el.Platform.SetLineWidth(el.ContourLineWidth)
	el.Platform.SetStrokeStyle(el.ContourColor)
	independentDraw.DrawSimpleLineIntoThePath(el.Platform, x0, y0, x0, y1)
	el.Platform.Stroke()
}

func (el *ResizeBoxes) lineFromCornerDToCornerE() {
	x0 := el.boxDX + float64(el.Dimensions.Width/2)
	y0 := el.boxDY + float64(el.Dimensions.Height)

	y1 := el.boxEY

	el.Platform.SetLineWidth(el.ContourLineWidth)
	el.Platform.SetStrokeStyle(el.ContourColor)
	independentDraw.DrawSimpleLineIntoThePath(el.Platform, x0, y0, x0, y1)
	el.Platform.Stroke()
}

func (el *ResizeBoxes) lineFromCornerEToCornerF() {
	x0 := el.boxEX
	y0 := el.boxEY + float64(el.Dimensions.Height/2)

	x1 := el.boxFX + float64(el.Dimensions.Width)

	el.Platform.SetLineWidth(el.ContourLineWidth)
	el.Platform.SetStrokeStyle(el.ContourColor)
	independentDraw.DrawSimpleLineIntoThePath(el.Platform, x0, y0, x1, y0)
	el.Platform.Stroke()
}

func (el *ResizeBoxes) lineFromCornerFToCornerG() {
	x0 := el.boxFX
	y0 := el.boxFY + float64(el.Dimensions.Height/2)

	x1 := el.boxGX + float64(el.Dimensions.Width)

	el.Platform.SetLineWidth(el.ContourLineWidth)
	el.Platform.SetStrokeStyle(el.ContourColor)
	independentDraw.DrawSimpleLineIntoThePath(el.Platform, x0, y0, x1, y0)
	el.Platform.Stroke()
}

func (el *ResizeBoxes) lineFromCornerGToCornerH() {
	x0 := el.boxGX + float64(el.Dimensions.Width/2)
	y0 := el.boxGY

	y1 := el.boxHY + float64(el.Dimensions.Width)

	el.Platform.SetLineWidth(el.ContourLineWidth)
	el.Platform.SetStrokeStyle(el.ContourColor)
	independentDraw.DrawSimpleLineIntoThePath(el.Platform, x0, y0, x0, y1)
	el.Platform.Stroke()
}

func (el *ResizeBoxes) lineFromCornerHToCornerA() {
	x0 := el.boxHX + float64(el.Dimensions.Width/2)
	y0 := el.boxHY

	y1 := el.boxAY + float64(el.Dimensions.Width)

	el.Platform.SetLineWidth(el.ContourLineWidth)
	el.Platform.SetStrokeStyle(el.ContourColor)
	independentDraw.DrawSimpleLineIntoThePath(el.Platform, x0, y0, x0, y1)
	el.Platform.Stroke()
}

func (el *ResizeBoxes) getEventOverCorner(boxX, boxY, x, y float64) bool {
	if x < boxX || x > boxX+float64(el.Dimensions.Width) {
		return false
	}

	if y < boxY || y > boxY+float64(el.Dimensions.Height) {
		return false
	}

	return true
}

func (el *ResizeBoxes) GetCollisionBox(xEvent, yEvent float64) bool {
	return float64(el.outBoxDimensions.X) <= xEvent && float64(el.outBoxDimensions.X+el.outBoxDimensions.Width) >= xEvent &&
		float64(el.outBoxDimensions.Y) <= yEvent && float64(el.outBoxDimensions.Y+el.outBoxDimensions.Height) >= yEvent
}

func (el *ResizeBoxes) ProcessMousePosition(x, y float64, collision bool) {
	if el.getEventOverCorner(el.boxAX, el.boxAY, x, y) {
		el.Platform.SetMouseCursor(el.MouseCornerA)
	} else if el.getEventOverCorner(el.boxBX, el.boxBY, x, y) {
		el.Platform.SetMouseCursor(el.MouseCornerB)
	} else if el.getEventOverCorner(el.boxCX, el.boxCY, x, y) {
		el.Platform.SetMouseCursor(el.MouseCornerC)
	} else if el.getEventOverCorner(el.boxDX, el.boxDY, x, y) {
		el.Platform.SetMouseCursor(el.MouseCornerD)
	} else if el.getEventOverCorner(el.boxEX, el.boxEY, x, y) {
		el.Platform.SetMouseCursor(el.MouseCornerE)
	} else if el.getEventOverCorner(el.boxFX, el.boxFY, x, y) {
		el.Platform.SetMouseCursor(el.MouseCornerF)
	} else if el.getEventOverCorner(el.boxGX, el.boxGY, x, y) {
		el.Platform.SetMouseCursor(el.MouseCornerG)
	} else if el.getEventOverCorner(el.boxHX, el.boxHY, x, y) {
		el.Platform.SetMouseCursor(el.MouseCornerH)
	} else {
		el.Platform.SetMouseCursor(el.MouseGeneric)
	}
}
