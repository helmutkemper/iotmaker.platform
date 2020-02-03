package simple

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/point"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/basic"
)

// en: Draw a rectangle box with a rounded corners [optional]
//     Use the border equal to half the height to make the points of the drawing
//     become a semicircle
//
// pt_br: Desenha um retângulo com bordas arredondadas [opcional]
//     Use a borda igual a metade da altura para um fazer as pontas do desenho serem
//     um semicírculo
type BoxWithRoundedCorners struct {
	basic.Sprite

	x1 int
	x2 int
	x3 int
	x4 int

	y1 int
	y2 int
	y3 int
	y4 int
}

// en: Calculate a rectangle
//
// pt_br: Calcula o retângulo
func (el *BoxWithRoundedCorners) calculate() {
	el.x1 = el.Dimensions.X
	el.x2 = el.x1 + el.Dimensions.Border
	el.x3 = el.x2 + el.Dimensions.Width - 2*el.Dimensions.Border
	el.x4 = el.x3 + el.Dimensions.Border

	el.y1 = el.Dimensions.Y
	el.y2 = el.y1 + el.Dimensions.Border
	el.y3 = el.y2 + el.Dimensions.Height - 2*el.Dimensions.Border
	el.y4 = el.y3 + el.Dimensions.Border
}

func (el *BoxWithRoundedCorners) Crete() {
	el.Draw()
}

func (el *BoxWithRoundedCorners) Clear() {
	el.ScratchPad.ClearRect(el.Dimensions.X, el.Dimensions.Y, el.Dimensions.Width, el.Dimensions.Height)
}

// en: Draw function
//
// pt_br: função de desenho
func (el *BoxWithRoundedCorners) Draw() {
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

	el.ScratchPad.Save()
	el.calculate()

	el.ColorFiltersStart(el.ScratchPad)

	if el.Ink != nil {
		p0 := point.Point{el.OutBoxDimensions.X, el.OutBoxDimensions.Y}
		p1 := point.Point{el.OutBoxDimensions.X + el.OutBoxDimensions.Width, el.OutBoxDimensions.Y + el.OutBoxDimensions.Height}

		el.Ink.SetGradientLinearRectangle(p0, p1)
	}

	el.ScratchPad.BeginPath()
	el.ScratchPad.MoveTo(el.x2, el.y1)                                    // a
	el.ScratchPad.LineTo(el.x3, el.y1)                                    // a->b
	el.ScratchPad.ArcTo(el.x4, el.y1, el.x4, el.y2, el.Dimensions.Border) // c->d
	el.ScratchPad.LineTo(el.x4, el.y3)                                    // d->e
	el.ScratchPad.ArcTo(el.x4, el.y4, el.x3, el.y4, el.Dimensions.Border) // f->g
	el.ScratchPad.LineTo(el.x2, el.y4)                                    // g->h
	el.ScratchPad.ArcTo(el.x1, el.y4, el.x1, el.y3, el.Dimensions.Border) // i->j
	el.ScratchPad.LineTo(el.x1, el.y2)                                    // j->k
	el.ScratchPad.ArcTo(el.x1, el.y1, el.x2, el.y1, el.Dimensions.Border) // i->j
	el.ScratchPad.ClosePath(el.x2, el.y1)                                 // a

	el.ScratchPad.Restore()
}
