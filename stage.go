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

--------
        \
         |
         |
         |
*/
func (el *Stage) AddBasicBox(x, y, width, height, border, radius float64) {

	el.Stage.BeginPath()
	el.Stage.MoveTo(x, y)
	el.Stage.LineTo(x+width-border/2, y)
	el.Stage.MoveTo(x+width+border/2, y)
	el.Stage.LineTo(x+width+border/2, y+border/2)
	//el.Stage.ArcTo(x + width - border/2, y, x + width + border/2, y + border/2, radius)
	el.Stage.LineTo(x+width+border/2, y+height+border/2)
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

	el.AddBasicBox(20, 20, 100, 100, 40, 50)
	//el.Stage.BeginPath()
	//el.Stage.StrokeStyle("#FF0000")
	//el.Stage.MoveTo(0.0, 0.0)
	//el.Stage.LineTo(el.OriginalWidth, el.OriginalHeight)
	//el.Stage.Stroke()

}
