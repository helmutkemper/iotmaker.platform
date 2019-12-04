package basicBox

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	iotmaker_threadsafe "github.com/helmutkemper/iotmaker.threadsafe"
	"image/color"
)

type ImageDataCaptureMethod uint8

const (
	KImageDataCaptureMethodBooleanSensibility ImageDataCaptureMethod = iota
	KImageDataCaptureMethodCompleteData
	KImageDataCaptureMethodAlphaChannelOnly
)

type BasicBox struct {
	Platform   iotmaker_platform_IDraw.IDraw
	ScratchPad iotmaker_platform_IDraw.IDraw
	Id         string
	Dimensions Dimensions
	Ink        Ink

	imageDataMethod           ImageDataCaptureMethod
	imageDataComplete         map[int]map[int]color.RGBA
	imageDataAlphaChannel     map[int]map[int]uint8
	imageDataBooleanCollision map[int]map[int]bool

	prepareShadowFilterFunctionPointer   func(iotmaker_platform_IDraw.ICanvasShadow)
	prepareGradientFilterFunctionPointer func(iotmaker_platform_IDraw.ICanvasGradient)

	// see calculateCoordinates() - start
	x1 int
	y1 int

	x2 int
	y2 int

	x3 int
	y3 int

	x4 int
	y4 int
	// see calculateCoordinates() - end

	// see clearRectangle() and getCompleteImageData() - start
	x      int
	y      int
	width  int
	height int
	// see clearRectangle() and getCompleteImageData() - end

	alphaChannelSensibility  uint8
	enableDataImageCalculate bool
}

func (el *BasicBox) prepareShadowFilter() {
	if el.prepareShadowFilterFunctionPointer != nil {
		el.prepareShadowFilterFunctionPointer(el.Platform)
	}
}

func (el *BasicBox) prepareGradientFilter(platform iotmaker_platform_IDraw.IDraw) {
	if el.prepareGradientFilterFunctionPointer != nil {
		el.prepareGradientFilterFunctionPointer(platform)
	}
}

func (el *BasicBox) configShadowPlatformAndFilter() {
	if el.Ink.Shadow == nil {
		return
	}

	el.setShadowFilter(el.Ink.Shadow.PrepareFilter)
}

func (el *BasicBox) configGradientPlatformAndFilter() {
	if el.Ink.Gradient == nil {
		return
	}

	el.setGradientAndMountColorListFilter(el.Ink.Gradient.PrepareFilter)
}

func (el *BasicBox) setGradientAndMountColorListFilter(f func(iotmaker_platform_IDraw.ICanvasGradient)) {
	el.prepareGradientFilterFunctionPointer = f
}

func (el *BasicBox) setShadowFilter(f func(iotmaker_platform_IDraw.ICanvasShadow)) {
	el.prepareShadowFilterFunctionPointer = f
}

func (el *BasicBox) calculateCoordinates() {
	//  draw_1:
	//              border        border
	//             x1  x2         x3 x4
	//           l     a          b     c
	//        y1    +--|----------|--+    y1
	//  border      |                |      border
	//        y2 k ---              --- d y2
	//              |                |
	//              |                |
	//        y3 j ---              --- e y3
	//  border      |                |      border
	//        y4    +--|----------|--+    y4
	//           i     h          g     f
	//             x1  x2         x3 x4
	//              border        border

	// correct the line width size
	el.Dimensions.X -= el.Ink.LineWidth / 2
	el.Dimensions.Y -= el.Ink.LineWidth / 2
	el.Dimensions.Width -= el.Ink.LineWidth
	el.Dimensions.Height -= el.Ink.LineWidth

	// set coordinates from de box in draw_1
	el.x1 = el.Dimensions.X
	el.x2 = el.x1 + el.Dimensions.Border
	el.x3 = el.x2 + el.Dimensions.Width - 2*el.Dimensions.Border
	el.x4 = el.x3 + el.Dimensions.Border

	el.y1 = el.Dimensions.Y
	el.y2 = el.y1 + el.Dimensions.Border
	el.y3 = el.y2 + el.Dimensions.Height - 2*el.Dimensions.Border
	el.y4 = el.y3 + el.Dimensions.Border

	// calculate outline from the box
	el.x = el.Dimensions.X - el.Ink.LineWidth/2
	el.y = el.Dimensions.Y - el.Ink.LineWidth/2
	el.width = el.Dimensions.Width + el.Ink.LineWidth
	el.height = el.Dimensions.Height + el.Ink.LineWidth
}

func (el *BasicBox) getCompleteImageData(platform iotmaker_platform_IDraw.IDraw) {
	el.imageDataComplete = platform.GetImageData(el.x, el.y, el.width, el.height)
}

func (el *BasicBox) getImageDataAlphaChannelOnly(platform iotmaker_platform_IDraw.IDraw) {
	el.imageDataAlphaChannel = platform.GetImageDataAlphaChannelOnly(el.x, el.y, el.width, el.height)
}
func (el *BasicBox) getImageDataCollisionByAlphaChannelValue(platform iotmaker_platform_IDraw.IDraw) {
	el.imageDataBooleanCollision = platform.GetImageDataCollisionByAlphaChannelValue(el.x, el.y, el.width, el.height, el.alphaChannelSensibility)
}

func (el *BasicBox) clearRectangle(platform iotmaker_platform_IDraw.IDraw) {
	platform.ClearRect(el.x, el.y, el.width, el.height)
}

func (el *BasicBox) drawInvisible(platform iotmaker_platform_IDraw.IDraw) {
	platform.SetLineWidth(el.Ink.LineWidth)
	platform.BeginPath()
	platform.MoveTo(el.x2, el.y1)                                    // a
	platform.LineTo(el.x3, el.y1)                                    // a->b
	platform.ArcTo(el.x4, el.y1, el.x4, el.y2, el.Dimensions.Border) // c->d
	platform.LineTo(el.x4, el.y3)                                    // d->e
	platform.ArcTo(el.x4, el.y4, el.x3, el.y4, el.Dimensions.Border) // f->g
	platform.LineTo(el.x2, el.y4)                                    // g->h
	platform.ArcTo(el.x1, el.y4, el.x1, el.y3, el.Dimensions.Border) // i->j
	platform.LineTo(el.x1, el.y2)                                    // j->k
	platform.ArcTo(el.x1, el.y1, el.x2, el.y1, el.Dimensions.Border) // i->j
	platform.ClosePath(el.x2, el.y1)                                 // a
}

func (el *BasicBox) drawVisible() {
	el.drawInvisible(el.Platform)
	el.prepareGradientFilter(el.Platform)
	el.prepareShadowFilter()
	el.Platform.Stroke()
}

func (el *BasicBox) Create() {
	el.calculateCoordinates()

	if el.enableDataImageCalculate == true {
		el.CalculateImageData()
	}

	el.drawVisible()

	if el.alphaChannelSensibility == 0 {
		el.alphaChannelSensibility = 255 * 0.1
	}
}

// see SetEnableDataImageCalculate()
// see SetAlphaChannelSensibility()
// see SetImageDataCalculateMethod()
func (el *BasicBox) CalculateImageData() {

	switch el.imageDataMethod {
	case KImageDataCaptureMethodCompleteData:
		iotmaker_threadsafe.ScratchPad(
			el.ScratchPad,
			el.prepareGradientFilter,
			el.drawInvisible,
			el.getCompleteImageData,
			el.clearRectangle,
		)
	case KImageDataCaptureMethodAlphaChannelOnly:
		iotmaker_threadsafe.ScratchPad(
			el.ScratchPad,
			el.prepareGradientFilter,
			el.drawInvisible,
			el.getImageDataAlphaChannelOnly,
			el.clearRectangle,
		)
	case KImageDataCaptureMethodBooleanSensibility:
		iotmaker_threadsafe.ScratchPad(
			el.ScratchPad,
			el.prepareGradientFilter,
			el.drawInvisible,
			el.getImageDataCollisionByAlphaChannelValue,
			el.clearRectangle,
		)
	}
}

// pt_br: Define o método usado para calcular os dados da imagem usado para detectar colisão
//     KImageDataCaptureMethodBooleanSensibility: Arquiva um mapa de booleanos no formato mapa[x][y]bool
//     KImageDataCaptureMethodCompleteData: Arquiva um mapa de RGBA no formato mapa[x][y]RGBA
//     KImageDataCaptureMethodAlphaChannelOnly: Arquiva um mapa de uint8 no formato mapa[x][y]uint8
//     Veja também SetAlphaChannelSensibility() e SetEnableDataImageCalculate()
func (el *BasicBox) SetImageDataCalculateMethod(value ImageDataCaptureMethod) {
	el.imageDataMethod = value
}

func (el *BasicBox) SetAlphaChannelSensibility(value uint8) {
	el.alphaChannelSensibility = value
}

func (el *BasicBox) SetEnableDataImageCalculate(value bool) {
	el.enableDataImageCalculate = value
}

func (el *BasicBox) GetCollisionByAlphaChannel(x, y int) bool {
	return el.imageDataBooleanCollision[x][y]
}

func (el *BasicBox) GetCollisionBySimpleBox(x, y int) bool {
	return el.x <= x && el.x+el.width >= x && el.y <= y && el.y+el.height >= y
}

func (el *BasicBox) GetPixelAlphaChannel(x, y int) uint8 {
	if el.enableDataImageCalculate == false {
		return 0
	}

	return el.imageDataComplete[x][y].A
}

func (el *BasicBox) GetPixelColor(x, y int) color.RGBA {
	if el.enableDataImageCalculate == false {
		return color.RGBA{}
	}

	return el.imageDataComplete[x][y]
}
