package abstractType

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"image/color"
	"reflect"
)

type Shadow struct {
	Color   color.RGBA
	Blur    int
	OffsetX int
	OffsetY int
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

type _BasicBox struct {
	Platform  iotmaker_platform_IDraw.IDraw
	Density   float64
	Id        string
	X         iotmaker_platform_coordinate.Coordinate
	Y         iotmaker_platform_coordinate.Coordinate
	Width     iotmaker_platform_coordinate.Coordinate
	Height    iotmaker_platform_coordinate.Coordinate
	Border    iotmaker_platform_coordinate.Coordinate
	LineWidth iotmaker_platform_coordinate.Coordinate
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

	if !reflect.DeepEqual(el.Shadow, Shadow{}) {
		el.Platform.ShadowBlur(el.Shadow.Blur)
		el.Platform.ShadowColor(el.Shadow.Color)
		el.Platform.ShadowOffsetX(el.Shadow.OffsetX)
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
