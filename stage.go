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

/*
ctx.beginPath();
ctx.moveTo(x, y);
ctx.lineTo(x+width, y);
ctx.arcTo(x+width+border, y, x+width+border, y+border, radius);
ctx.lineTo(x+width+border, y+border+height);
ctx.stroke();
*/
// pt_br: Desenha uma caixa básica, com bordas arrendondadas
//
//       border        border
//       x1 x2         x3 x4
//       +--|----------|--+  y1
//       |                |    border
//      ---              --- y2
//       |                |
//       |                |
//      ---              --- y3
//       |                |    border
//       +--|----------|--+  y4
func (el *Stage) AddBasicBox(x, y, width, height, border float64) {

	// tired programmer's rules (kiss)
	x1 := x
	x2 := x1 + border
	x3 := x2 + width - 2*border
	x4 := x3 + border

	y1 := y
	y2 := y1 + border
	y3 := y2 + height - 2*border
	y4 := y3 + border

	el.Stage.BeginPath()
	el.Stage.MoveTo(x1, y1)
	el.Stage.LineTo(x2, y1)

	el.Stage.ArcTo(x3, y1, x4, y2, border)
	el.Stage.LineTo(x4, y3)

	el.Stage.Stroke()

	//el.Stage.MoveTo(20, 20)
	//el.Stage.LineTo(100, 20)
	//el.Stage.ArcTo(150, 20, 150, 70, 50)
	//el.Stage.LineTo(150, 120)
	//el.Stage.Stroke()

	//el.Stage.Rect( 10, 10, 300, 300 )
	//el.Stage.Stroke()
}

func (el *Stage) NewStageOnTheRoot(id string) {
	el.OriginalWidth, el.OriginalHeight = el.GetRootElementWidthAndHeight()
	el.Stage = pwb.NewCanvasWith2DContext("canvas_id", el.OriginalWidth, el.OriginalHeight)
	el.Stage.AppendToDocumentBody()

	el.AddBasicBox(20, 20, 100, 100, 10)
	//el.Stage.BeginPath()
	//el.Stage.StrokeStyle("#FF0000")
	//el.Stage.MoveTo(0.0, 0.0)
	//el.Stage.LineTo(el.OriginalWidth, el.OriginalHeight)
	//el.Stage.Stroke()

}
