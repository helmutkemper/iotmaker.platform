package abstractType

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"image/color"
	"reflect"
)

type Shadow struct {
	Color                 color.RGBA
	colorHasInitialized   bool
	Blur                  int
	blurHasInitialized    bool
	OffsetX               int
	offsetXHasInitialized bool
	OffsetY               int
	offsetYHasInitialized bool
}

func (el *Shadow) SetColor(color color.RGBA) {
	el.Color = color
	el.colorHasInitialized = true
}

func (el *Shadow) ClearColor() {
	el.colorHasInitialized = false
}

func (el *Shadow) SetBlur(blur int) {
	el.Blur = blur
	el.blurHasInitialized = true
}

func (el *Shadow) ClearBlur() {
	el.blurHasInitialized = false
}

func (el *Shadow) SetOffsetX(x int) {
	el.OffsetX = x
	el.offsetXHasInitialized = true
}

func (el *Shadow) ClearOffsetX() {
	el.offsetXHasInitialized = false
}

func (el *Shadow) SetOffsetY(y int) {
	el.OffsetY = y
	el.offsetYHasInitialized = true
}

func (el *Shadow) ClearOffsetY() {
	el.offsetYHasInitialized = false
}

type DimensionsBasicBox struct {
	Density   float64
	X         int
	Y         int
	Width     int
	Height    int
	Border    int
	LineWidth int
}

type BasicBox struct {
	Platform   iotmaker_platform_IDraw.IDraw
	Id         string
	Dimensions DimensionsBasicBox
	Shadow     Shadow
	Gradient   Gradient
}

type CoordinateBox struct {
	X0 int
	Y0 int
	X1 int
	Y1 int
}

type Gradient struct {
	Type       int
	Coordinate CoordinateBox
	ColorList  []ColorStop
}

type ColorStop struct {
	Stop  float64
	Color color.RGBA
}

func NewBasicBox(config BasicBox) BasicBox {

	coordinate := iotmaker_platform_coordinate.Coordinate{}
	coordinate.SetDensityFactor(config.Dimensions.Density)
	coordinate.Set(config.Dimensions.X)
	config.Dimensions.X = coordinate.Int()

	coordinate.Set(config.Dimensions.Y)
	config.Dimensions.Y = coordinate.Int()

	coordinate.Set(config.Dimensions.Width)
	config.Dimensions.Width = coordinate.Int()

	coordinate.Set(config.Dimensions.Height)
	config.Dimensions.Height = coordinate.Int()

	coordinate.Set(config.Dimensions.Border)
	config.Dimensions.Border = coordinate.Int()

	coordinate.Set(config.Dimensions.LineWidth)
	config.Dimensions.LineWidth = coordinate.Int()

	config.Create()

	return config
}

func (el *BasicBox) prepareGradientCoordinateFilter() {
	x0 := el.Gradient.Coordinate.X0
	x1 := el.Gradient.Coordinate.X1

	y0 := el.Gradient.Coordinate.Y0
	y1 := el.Gradient.Coordinate.Y1

	if x0 == 0 && x1 == 0 && y0 == 0 && y1 == 0 {
		el.Gradient.Coordinate.X0 = el.Dimensions.X
		el.Gradient.Coordinate.X1 = x0 + el.Dimensions.Width

		el.Gradient.Coordinate.Y0 = el.Dimensions.Y
		el.Gradient.Coordinate.Y1 = y0 + el.Dimensions.Height
	}
}

func (el *BasicBox) prepareGradientMountColorListFilter() {
	x0 := el.Gradient.Coordinate.X0
	x1 := el.Gradient.Coordinate.X1

	y0 := el.Gradient.Coordinate.Y0
	y1 := el.Gradient.Coordinate.Y1

	var gradient interface{}

	//TODO: radial gradient
	if el.Gradient.Type == 0 {
		gradient = el.Platform.CreateLinearGradient(x0, y0, x1, y1)
	} else {
		//FIXME: mudar para radial
		gradient = el.Platform.CreateLinearGradient(x0, y0, x1, y1)
	}
}

func (el *BasicBox) prepareGradientFilter() {

	el.prepareGradientCoordinateFilter()
}

func (el *BasicBox) prepareShadowFilter() {
	if el.Shadow.colorHasInitialized == true {
		el.Platform.ShadowColor(el.Shadow.Color)
	}

	if el.Shadow.blurHasInitialized == true {
		el.Platform.ShadowBlur(el.Shadow.Blur)
	}

	if el.Shadow.offsetXHasInitialized == true {
		el.Platform.ShadowOffsetX(el.Shadow.OffsetX)
	}

	if el.Shadow.offsetYHasInitialized == true {
		el.Platform.ShadowOffsetY(el.Shadow.OffsetY)
	}
}

func (el *BasicBox) prepareToDrawCanvas() {
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

	var x1, x2, x3, x4 iotmaker_platform_coordinate.Coordinate
	var y1, y2, y3, y4 iotmaker_platform_coordinate.Coordinate

	// set screen density of pixels
	x1.SetDensityFactor(el.Dimensions.Density)
	x2.SetDensityFactor(el.Dimensions.Density)
	x3.SetDensityFactor(el.Dimensions.Density)
	x4.SetDensityFactor(el.Dimensions.Density)

	y1.SetDensityFactor(el.Dimensions.Density)
	y2.SetDensityFactor(el.Dimensions.Density)
	y3.SetDensityFactor(el.Dimensions.Density)
	y4.SetDensityFactor(el.Dimensions.Density)

	// correct the line width size
	el.Dimensions.X += el.Dimensions.LineWidth / 2
	el.Dimensions.Y += el.Dimensions.LineWidth / 2
	el.Dimensions.Width -= el.Dimensions.LineWidth / 2
	el.Dimensions.Height -= el.Dimensions.LineWidth / 2

	// set coordinates from de box in draw_1
	x1.Set(el.Dimensions.X)
	x2.Set(x1.Int() + el.Dimensions.Border)
	x3.Set(x2.Int() + el.Dimensions.Width - 2*el.Dimensions.Border)
	x4.Set(x3.Int() + el.Dimensions.Border)

	y1.Set(el.Dimensions.Y)
	y2.Set(y1.Int() + el.Dimensions.Border)
	y3.Set(y2.Int() + el.Dimensions.Height - 2*el.Dimensions.Border)
	y4.Set(y3.Int() + el.Dimensions.Border)

	el.Platform.LineWidth(el.Dimensions.LineWidth)

	el.Platform.BeginPath()

	if !reflect.DeepEqual(el.Gradient, Gradient{}) {
		//todo: continua aqui
	}

	el.Platform.MoveTo(x2.Int(), y1.Int())                                          // a
	el.Platform.LineTo(x3.Int(), y1.Int())                                          // a->b
	el.Platform.ArcTo(x4.Int(), y1.Int(), x4.Int(), y2.Int(), el.Dimensions.Border) // c->d
	el.Platform.LineTo(x4.Int(), y3.Int())                                          // d->e
	el.Platform.ArcTo(x4.Int(), y4.Int(), x3.Int(), y4.Int(), el.Dimensions.Border) // f->g
	el.Platform.LineTo(x2.Int(), y4.Int())                                          // g->h
	el.Platform.ArcTo(x1.Int(), y4.Int(), x1.Int(), y3.Int(), el.Dimensions.Border) // i->j
	el.Platform.LineTo(x1.Int(), y2.Int())                                          // j->k
	el.Platform.ArcTo(x1.Int(), y1.Int(), x2.Int(), y1.Int(), el.Dimensions.Border) // i->j
	el.Platform.ClosePath(x2.Int(), y1.Int())                                       // a

	// pegar dados da imagem aqui

	el.Platform.Stroke()
}

func (el *BasicBox) Create() {
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

	var x1, x2, x3, x4 iotmaker_platform_coordinate.Coordinate
	var y1, y2, y3, y4 iotmaker_platform_coordinate.Coordinate

	// set screen density of pixels
	x1.SetDensityFactor(el.Dimensions.Density)
	x2.SetDensityFactor(el.Dimensions.Density)
	x3.SetDensityFactor(el.Dimensions.Density)
	x4.SetDensityFactor(el.Dimensions.Density)

	y1.SetDensityFactor(el.Dimensions.Density)
	y2.SetDensityFactor(el.Dimensions.Density)
	y3.SetDensityFactor(el.Dimensions.Density)
	y4.SetDensityFactor(el.Dimensions.Density)

	// correct the line width size
	el.Dimensions.X += el.Dimensions.LineWidth / 2
	el.Dimensions.Y += el.Dimensions.LineWidth / 2
	el.Dimensions.Width -= el.Dimensions.LineWidth / 2
	el.Dimensions.Height -= el.Dimensions.LineWidth / 2

	// set coordinates from de box in draw_1
	x1.Set(el.Dimensions.X)
	x2.Set(x1.Int() + el.Dimensions.Border)
	x3.Set(x2.Int() + el.Dimensions.Width - 2*el.Dimensions.Border)
	x4.Set(x3.Int() + el.Dimensions.Border)

	y1.Set(el.Dimensions.Y)
	y2.Set(y1.Int() + el.Dimensions.Border)
	y3.Set(y2.Int() + el.Dimensions.Height - 2*el.Dimensions.Border)
	y4.Set(y3.Int() + el.Dimensions.Border)

	el.Platform.LineWidth(el.Dimensions.LineWidth)

	el.Platform.BeginPath()

	if !reflect.DeepEqual(el.Gradient, Gradient{}) {
		//todo: continua aqui
	}

	if el.Shadow.colorHasInitialized == true {
		el.Platform.ShadowColor(el.Shadow.Color)
	}

	if el.Shadow.blurHasInitialized == true {
		el.Platform.ShadowBlur(el.Shadow.Blur)
	}

	if el.Shadow.offsetXHasInitialized == true {
		el.Platform.ShadowOffsetX(el.Shadow.OffsetX)
	}

	if el.Shadow.offsetYHasInitialized == true {
		el.Platform.ShadowOffsetY(el.Shadow.OffsetY)
	}

	el.Platform.MoveTo(x2.Int(), y1.Int())                                          // a
	el.Platform.LineTo(x3.Int(), y1.Int())                                          // a->b
	el.Platform.ArcTo(x4.Int(), y1.Int(), x4.Int(), y2.Int(), el.Dimensions.Border) // c->d
	el.Platform.LineTo(x4.Int(), y3.Int())                                          // d->e
	el.Platform.ArcTo(x4.Int(), y4.Int(), x3.Int(), y4.Int(), el.Dimensions.Border) // f->g
	el.Platform.LineTo(x2.Int(), y4.Int())                                          // g->h
	el.Platform.ArcTo(x1.Int(), y4.Int(), x1.Int(), y3.Int(), el.Dimensions.Border) // i->j
	el.Platform.LineTo(x1.Int(), y2.Int())                                          // j->k
	el.Platform.ArcTo(x1.Int(), y1.Int(), x2.Int(), y1.Int(), el.Dimensions.Border) // i->j
	el.Platform.ClosePath(x2.Int(), y1.Int())                                       // a

	// pegar dados da imagem aqui

	el.Platform.Stroke()
}
