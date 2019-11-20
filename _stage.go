package iotmaker_platform

import (
	pwb "github.com/helmutkemper/iotmaker.platform.webbrowser"
	"strconv"
	"sync"
	"syscall/js"
)

var mouseX, mouseY float64

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

var gx1, gx4, gy1, gy4 float64

// en: Draw a basic box with rounded edges
//
// pt_br: Desenha uma caixa básica, com bordas arrendondadas
func (el *Stage) AddBasicBox(x, y, width, height, border, LineWidth float64) {
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
	x += LineWidth / 2
	width -= LineWidth / 2
	height -= LineWidth / 2

	x1 := x
	x2 := x1 + border
	x3 := x2 + width - 2*border
	x4 := x3 + border

	y1 := y
	y2 := y1 + border
	y3 := y2 + height - 2*border
	y4 := y3 + border

	gx1 = x1
	gx4 = x4
	gy1 = y1
	gy4 = y4

	el.Stage.LineWidth(LineWidth)

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
	el.Stage = pwb.NewCanvasWith2DContext("canvas_id", el.OriginalWidth/2, el.OriginalHeight/2)
	el.Stage.AppendToDocumentBody()

	el.Stage.BeginPath()
	el.AddBasicBox(50, 50, 180, 180, 10, 20)
	el.Stage.Stroke()

	ele := pwb.NewElement()
	ele.Create("p", "id")
	ele.AppendElementToDocumentBody()

	/*
	     012
	   0 012
	   1 345
	   2 678

	   (y*w)+x


	      0    1    2    3
	    0 0000.0000.0000.0000
	      4    5    6    7
	    1 0000.0000.0000.0000
	      8    9    10   11
	    2 0000.0000.0000.0000
	      12   13   14   15
	    3 0000.0000.0000.0000
	*/
	mouseMoveEvt := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		mouseX = e.Get("clientX").Float()
		mouseY = e.Get("clientY").Float()
		test := js.Global().Get("document").Call("getElementById", "id")

		mouseX += 0
		if mouseX < 0 {
			mouseX = 0
		}
		mouseY += 0
		if mouseY < 0 {
			mouseY = 0
		}

		//dest := make([]byte, int((gx4-gx1)*(gy4-gy1)*8))

		data := el.Stage.SelfContext.Call("getImageData", 0, 0, 300, 300)
		output := data.Get("data")
		coor := int((mouseY*300 + mouseX) * 4)
		//if mouseX >= gx1 && mouseX <= gx4 && mouseY >= gy1 && mouseY <= gy4 {
		//if len( dest ) <= coor {

		if coor > 300*300*4 {
			test.Set("innerHTML", "0:0:0:0")

			return nil
		}

		test.Set("innerHTML",
			//strconv.FormatInt(int64(mouseX), 10)+
			//":"+
			//strconv.FormatInt(int64(mouseY), 10)+
			//":"+
			strconv.FormatInt(int64(output.Index(coor+0).Int()), 16)+
				":"+
				strconv.FormatInt(int64(output.Index(coor+1).Int()), 16)+
				":"+
				strconv.FormatInt(int64(output.Index(coor+2).Int()), 16)+
				":"+
				strconv.FormatInt(int64(output.Index(coor+3).Int()), 16),
		)
		return nil
		for i := 0; i != coor; i += 1 {
			if output.Index(i).Int() != 0 {
				test.Set("innerHTML", i)
				break
			}
		}

		//}
		//}

		return nil
	})
	defer mouseMoveEvt.Release()
	js.Global().Get("document").Call("addEventListener", "mousemove", mouseMoveEvt)

	var ws sync.WaitGroup
	ws.Add(1)
	ws.Wait()
}
