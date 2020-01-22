package simple

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
)

// en: Draw a rectangle box with a rounded corners [optional]
//     Use the border equal to half the height to make the points of the drawing
//     become a semicircle
//
// pt_br: Desenha um retângulo com bordas arredondadas [opcional]
//     Use a borda igual a metade da altura para um fazer as pontas do desenho serem
//     um semicírculo
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
//
type BoxWithRoundedCorners struct {
	Platform iotmaker_platform_IDraw.IDraw

	X      int
	Y      int
	Width  int
	Height int
	Border int

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
func (el *BoxWithRoundedCorners) Init() {
	el.x1 = el.X
	el.x2 = el.x1 + el.Border
	el.x3 = el.x2 + el.Width - 2*el.Border
	el.x4 = el.x3 + el.Border

	el.y1 = el.Y
	el.y2 = el.y1 + el.Border
	el.y3 = el.y2 + el.Height - 2*el.Border
	el.y4 = el.y3 + el.Border
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

	if el.Platform == nil {
		return
	}

	el.Platform.BeginPath()
	el.Platform.MoveTo(el.x2, el.y1)                         // a
	el.Platform.LineTo(el.x3, el.y1)                         // a->b
	el.Platform.ArcTo(el.x4, el.y1, el.x4, el.y2, el.Border) // c->d
	el.Platform.LineTo(el.x4, el.y3)                         // d->e
	el.Platform.ArcTo(el.x4, el.y4, el.x3, el.y4, el.Border) // f->g
	el.Platform.LineTo(el.x2, el.y4)                         // g->h
	el.Platform.ArcTo(el.x1, el.y4, el.x1, el.y3, el.Border) // i->j
	el.Platform.LineTo(el.x1, el.y2)                         // j->k
	el.Platform.ArcTo(el.x1, el.y1, el.x2, el.y1, el.Border) // i->j
	el.Platform.ClosePath(el.x2, el.y1)                      // a
}
