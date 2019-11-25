package abstractType

import (
	"github.com/helmutkemper/iotmaker.platform"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
)

func NewBasicBox(platform iotmaker_platform.ICanvas, id string, density float64, x, y, width, height, border, lineWidth int) BasicBox {

	coordinateX := iotmaker_platform_coordinate.Coordinate{}
	coordinateX.SetDensityFactor(density)
	coordinateX.Set(x)

	coordinateY := iotmaker_platform_coordinate.Coordinate{}
	coordinateY.SetDensityFactor(density)
	coordinateY.Set(y)

	coordinateWidth := iotmaker_platform_coordinate.Coordinate{}
	coordinateWidth.SetDensityFactor(density)
	coordinateWidth.Set(width)

	coordinateHeight := iotmaker_platform_coordinate.Coordinate{}
	coordinateHeight.SetDensityFactor(density)
	coordinateHeight.Set(height)

	coordinateBorder := iotmaker_platform_coordinate.Coordinate{}
	coordinateBorder.SetDensityFactor(density)
	coordinateBorder.Set(border)

	coordinateLineWidth := iotmaker_platform_coordinate.Coordinate{}
	coordinateLineWidth.SetDensityFactor(density)
	coordinateLineWidth.Set(lineWidth)

	bb := BasicBox{
		Platform:  platform,
		Density:   density,
		Id:        id,
		X:         coordinateX,
		Y:         coordinateY,
		Width:     coordinateWidth,
		Height:    coordinateHeight,
		Border:    coordinateBorder,
		LineWidth: coordinateLineWidth,
	}
	bb.Create()

	return bb
}

type BasicBox struct {
	Platform  iotmaker_platform.ICanvas
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
	x1.SetDensityFactor(el.Density)
	x2.SetDensityFactor(el.Density)
	x3.SetDensityFactor(el.Density)
	x4.SetDensityFactor(el.Density)

	y1.SetDensityFactor(el.Density)
	y2.SetDensityFactor(el.Density)
	y3.SetDensityFactor(el.Density)
	y4.SetDensityFactor(el.Density)

	// correct the line width size
	el.X.Add(el.LineWidth.Int() / 2)
	el.Y.Add(el.LineWidth.Int() / 2)
	el.Width.Sub(el.LineWidth.Int() / 2)
	el.Height.Sub(el.LineWidth.Int() / 2)

	// set coordinates from de box in draw_1
	x1.Set(el.X.Int())
	x2.Set(x1.Int() + el.Border.Int())
	x3.Set(x2.Int() + el.Width.Int() - 2*el.Border.Int())
	x4.Set(x3.Int() + el.Border.Int())

	y1.Set(el.Y.Int())
	y2.Set(y1.Int() + el.Border.Int())
	y3.Set(y2.Int() + el.Height.Int() - 2*el.Border.Int())
	y4.Set(y3.Int() + el.Border.Int())

	el.Platform.LineWidth(el.LineWidth.Int())

	el.Platform.BeginPath()
	el.Platform.MoveTo(x2.Int(), y1.Int())                                     // a
	el.Platform.LineTo(x3.Int(), y1.Int())                                     // a->b
	el.Platform.ArcTo(x4.Int(), y1.Int(), x4.Int(), y2.Int(), el.Border.Int()) // c->d
	el.Platform.LineTo(x4.Int(), y3.Int())                                     // d->e
	el.Platform.ArcTo(x4.Int(), y4.Int(), x3.Int(), y4.Int(), el.Border.Int()) // f->g
	el.Platform.LineTo(x2.Int(), y4.Int())                                     // g->h
	el.Platform.ArcTo(x1.Int(), y4.Int(), x1.Int(), y3.Int(), el.Border.Int()) // i->j
	el.Platform.LineTo(x1.Int(), y2.Int())                                     // j->k
	el.Platform.ArcTo(x1.Int(), y1.Int(), x2.Int(), y1.Int(), el.Border.Int()) // i->j
	el.Platform.ClosePath(x2.Int(), y1.Int())                                  // a
	el.Platform.Stroke()
}
