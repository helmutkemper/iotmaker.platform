package independentDraw

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
)

// esta função deve ser apagada (obsoleta)

// en: This function create a path of a box with rounded corners into the platform
//
// pt_br: Esta função cria o caminho de um quadrado com cantos arredondados na plataforma
func DrawBoxWithRoundedCornersIntoThePath(platform iotmaker_platform_IDraw.IDraw, x, y, width, height, border float64) {
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

	if platform == nil {
		return
	}

	x1 := x
	x2 := x1 + border
	x3 := x2 + width - 2*border
	x4 := x3 + border

	y1 := y
	y2 := y1 + border
	y3 := y2 + height - 2*border
	y4 := y3 + border

	platform.BeginPath()
	platform.MoveTo(x2, y1)                // a
	platform.LineTo(x3, y1)                // a->b
	platform.ArcTo(x4, y1, x4, y2, border) // c->d
	platform.LineTo(x4, y3)                // d->e
	platform.ArcTo(x4, y4, x3, y4, border) // f->g
	platform.LineTo(x2, y4)                // g->h
	platform.ArcTo(x1, y4, x1, y3, border) // i->j
	platform.LineTo(x1, y2)                // j->k
	platform.ArcTo(x1, y1, x2, y1, border) // i->j
	platform.ClosePath(x2, y1)             // a

}
