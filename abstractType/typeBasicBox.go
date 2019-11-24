package abstractType

import (
	"github.com/helmutkemper/iotmaker.platform"
	iotmaker_platform_webbrowser "github.com/helmutkemper/iotmaker.platform.webbrowser"
)

func NewBasicBox(platform iotmaker_platform.ICanvas) BasicBox {
	bb := BasicBox{
		Platform:  platform,
		X:         &iotmaker_platform.Coordinate{},
		Y:         &iotmaker_platform.Coordinate{},
		Width:     &iotmaker_platform.Coordinate{},
		Height:    &iotmaker_platform.Coordinate{},
		Border:    &iotmaker_platform.Coordinate{},
		LineWidth: &iotmaker_platform.Coordinate{},
	}
	bb.Create()

	return bb
}

type BasicBox struct {
	Platform  iotmaker_platform.ICanvas
	Density   float64
	Id        string
	X         iotmaker_platform_webbrowser.ICoordinate
	Y         iotmaker_platform_webbrowser.ICoordinate
	Width     iotmaker_platform_webbrowser.ICoordinate
	Height    iotmaker_platform_webbrowser.ICoordinate
	Border    iotmaker_platform_webbrowser.ICoordinate
	LineWidth iotmaker_platform_webbrowser.ICoordinate
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

	var x1, x2, x3, x4 iotmaker_platform_webbrowser.ICoordinate
	var y1, y2, y3, y4 iotmaker_platform_webbrowser.ICoordinate

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

	// draw image
	//el.SelfElement = pwb.NewCanvasWith2DContext(input.Id, input.Width, input.Height)

	//el.SelfContext.LineWidth(input.LineWidth)

	el.Platform.MoveTo(x2, y1)                   // a
	el.Platform.LineTo(x3, y1)                   // a->b
	el.Platform.ArcTo(x4, y1, x4, y2, el.Border) // c->d
	el.Platform.LineTo(x4, y3)                   // d->e
	el.Platform.ArcTo(x4, y4, x3, y4, el.Border) // f->g
	el.Platform.LineTo(x2, y4)                   // g->h
	el.Platform.ArcTo(x1, y4, x1, y3, el.Border) // i->j
	el.Platform.LineTo(x1, y2)                   // j->k
	el.Platform.ArcTo(x1, y1, x2, y1, el.Border) // i->j
	el.Platform.ClosePath(x2, y1)                // a
}
