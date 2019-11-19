package iotmaker_platform

import (
	pwb "github.com/helmutkemper/iotmaker.platform.webbrowser"
	"syscall/js"
)

type Stage struct {
	pwb.Document
	Stage          pwb.Canvas
	OriginalWidth  float64
	OriginalHeight float64
}

func (el *Stage) GetRootElementWidthAndHeight() (float64, float64) {
	doc := js.Global().Get("document")
	width := doc.Get("body").Get("clientWidth").Float()
	height := doc.Get("body").Get("clientHeight").Float()

	return width, height
}

func (el *Stage) GetRootElementWidth() float64 {
	doc := js.Global().Get("document")
	width := doc.Get("body").Get("clientWidth").Float()

	return width
}

func (el *Stage) GetRootElementHeight() float64 {
	doc := js.Global().Get("document")
	height := doc.Get("body").Get("clientHeight").Float()

	return height
}

// en: Draw a basic box with rounded edges
//
// pt_br: Desenha uma caixa básica, com bordas arrendondadas
//
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
func (el *Stage) AddBasicBox(x, y, width, height, border float64) {

	// tired programmer's rules (kiss like)
	x1 := x
	x2 := x1 + border
	x3 := x2 + width - 2*border
	x4 := x3 + border

	y1 := y
	y2 := y1 + border
	y3 := y2 + height - 2*border
	y4 := y3 + border

	el.Stage.MoveTo(x2, y1)                // a
	el.Stage.LineTo(x3, y1)                // a->b
	el.Stage.ArcTo(x4, y1, x4, y2, border) // c->d
	el.Stage.LineTo(x4, y3)                // d->e
	el.Stage.ArcTo(x4, y4, x3, y4, border) // f->g
	el.Stage.LineTo(x2, y4)                // g->h
	el.Stage.ArcTo(x1, y4, x1, y3, border) // i->j
	el.Stage.LineTo(x1, y2)                // j->k
	el.Stage.ArcTo(x1, y1, x2, y1, border) // i->j
	el.Stage.ClosePath(x2, y1)             // a
}

func (el *Stage) NewStageOnTheRoot(id string) {
	el.OriginalWidth, el.OriginalHeight = el.GetRootElementWidthAndHeight()
	el.Stage = pwb.NewCanvasWith2DContext("canvas_id", el.OriginalWidth, el.OriginalHeight)
	el.Stage.AppendToDocumentBody()

	el.Stage.BeginPath()
	el.AddBasicBox(20, 20, 100, 100, 4)
	el.Stage.Stroke()

}
