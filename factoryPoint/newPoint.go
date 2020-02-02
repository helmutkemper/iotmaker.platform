package factoryPoint

import (
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/point"
)

// en: Creates a set of coordinates, taking into account the size of the screen.
//
// pt_br: Cria um par de coordenadas levando em consideração a densidade da tela.
//   x:        horizontal point of the screen.
//   y:        vertical point of the screen.
//   density:  density factor of the screen.
//   iDensity: density interface
func NewPoint(x, y float64, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) point.Point {

	densityX := iDensity
	densityX.Set(x)
	densityX.SetDensityFactor(density)

	densityY := iDensity
	densityY.Set(y)
	densityY.SetDensityFactor(density)

	return point.Point{
		X: densityX.Int(),
		Y: densityY.Int(),
	}
}
