package gradient

import iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"

// en: Make a new filter radial gradient with stroke for use with the canvas elements.
//
// Creates a radial gradient. The parameters represent two circles, one with its center at (x0, y0) and a radius of r0,
// and the other with its center at (x1, y1) with a radius of r1.
//
// pt_bt: Monta um novo filtro linear gradient para ser usado com os elementos do canvas.
//
// O gradiente radial cria um gradiente representado por dois circulos, um centralizado no ponto p0 (x0, y0) com raio r0
// e outro círculo centrado no ponto p1 (x1, y1) e com raio r1.
//    coordinateP0: Coordinate of the start point of the gradient. Please, use a NewPointWithRadius() function to set a point.
//    coordinateP1: Coordinate of the end point of the gradient. Please, use a NewPointWithRadius() function to set a point.
//    colorList:    Color position list. Please, use a NewColorPosition() and NewColorList() functions to set a color
//                  list. Example: NewColorList(NewColorPosition(), NewColorPosition(), ...)
func NewGradientRadialToStroke(coordinateP0, coordinateP1 PointWithRadius, colorList []ColorStop) iotmaker_platform_IDraw.IFilterGradientInterface {
	return &Gradient{
		Type:         KLinearGradientFill,
		CoordinateP0: coordinateP0,
		CoordinateP1: coordinateP1,
		ColorList:    colorList,
	}
}
