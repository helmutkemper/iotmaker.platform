package iotmaker_platform

import (
	pwb "github.com/helmutkemper/iotmaker.platform.webbrowser"
	canvas2 "github.com/helmutkemper/iotmaker.platform.webbrowser/canvas"
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
)

type BasicBox struct {
	canvas2.Document
	Draw
}

type Input struct {
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

	// tired programmer's rules (kiss like)
	input.X += input.LineWidth / 2
	input.Width -= input.LineWidth / 2
	input.Height -= input.LineWidth / 2

	x1 := input.X
	x2 := x1 + input.Border
	x3 := x2 + input.Width - 2*input.Border
	x4 := x3 + input.Border

	y1 := input.Y
	y2 := y1 + input.Border
	y3 := y2 + input.Height - 2*input.Border
	y4 := y3 + input.Border

	el.SelfElement = pwb.NewCanvasWith2DContext(input.Id, input.Width, input.Height)

	el.SelfContext.LineWidth(input.LineWidth)

	el.MoveTo(x2, y1)                      // a
	el.LineTo(x3, y1)                      // a->b
	el.ArcTo(x4, y1, x4, y2, input.Border) // c->d
	el.LineTo(x4, y3)                      // d->e
	el.ArcTo(x4, y4, x3, y4, input.Border) // f->g
	el.LineTo(x2, y4)                      // g->h
	el.ArcTo(x1, y4, x1, y3, input.Border) // i->j
	el.LineTo(x1, y2)                      // j->k
	el.ArcTo(x1, y1, x2, y1, input.Border) // i->j
	el.ClosePath(x2, y1)                   // a
}
