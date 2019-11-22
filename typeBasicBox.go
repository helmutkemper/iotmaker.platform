package iotmaker_platform

import (
	canvas2 "github.com/helmutkemper/iotmaker.platform.webbrowser/canvas"
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
)

type BasicBox struct {
	Platform ICanvas
	canvas2.Document
}

type Input struct {
	Density   float64
	Id        string
	X         iotmaker_types.Coordinate
	Y         iotmaker_types.Coordinate
	Width     iotmaker_types.Coordinate
	Height    iotmaker_types.Coordinate
	Border    iotmaker_types.Coordinate
	LineWidth iotmaker_types.Coordinate
}

func (el *BasicBox) Create(input Input) {
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

	var x1, x2, x3, x4 iotmaker_types.Coordinate
	var y1, y2, y3, y4 iotmaker_types.Coordinate

	input.X.Add(input.LineWidth.Int() / 2)
	input.Width.Sub(input.LineWidth.Int() / 2)
	input.Height.Sub(input.LineWidth.Int() / 2)

	x1.Set(input.X.Int())
	x2.Set(x1.Int() + input.Border.Int())
	x3.Set(x2.Int() + input.Width.Int() - 2*input.Border.Int())
	x4.Set(x3.Int() + input.Border.Int())

	y1.Set(input.Y.Int())
	y2.Set(y1.Int() + input.Border.Int())
	y3.Set(y2.Int() + input.Height.Int() - 2*input.Border.Int())
	y4.Set(y3.Int() + input.Border.Int())

	//el.SelfElement = pwb.NewCanvasWith2DContext(input.Id, input.Width, input.Height)

	//el.SelfContext.LineWidth(input.LineWidth)

	el.Platform.MoveTo(x2, y1)                      // a
	el.Platform.LineTo(x3, y1)                      // a->b
	el.Platform.ArcTo(x4, y1, x4, y2, input.Border) // c->d
	el.Platform.LineTo(x4, y3)                      // d->e
	el.Platform.ArcTo(x4, y4, x3, y4, input.Border) // f->g
	el.Platform.LineTo(x2, y4)                      // g->h
	el.Platform.ArcTo(x1, y4, x1, y3, input.Border) // i->j
	el.Platform.LineTo(x1, y2)                      // j->k
	el.Platform.ArcTo(x1, y1, x2, y1, input.Border) // i->j
	el.Platform.ClosePath(x2, y1)                   // a
}
