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
//   radius:   radius is a width radius, not a degrees angle
//   density:  density factor of the screen.
//   iDensity: density interface
func NewPointWithRadius(x, y, radius float64, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) point.PointWithRadius {

	densityX := iDensity
	densityX.Set(x)
	densityX.SetDensityFactor(density)

	densityY := iDensity
	densityY.Set(y)
	densityY.SetDensityFactor(density)

	densityR := iDensity
	densityR.Set(radius)
	densityR.SetDensityFactor(density)

	return point.PointWithRadius{
		X:      densityX.Int(),
		Y:      densityY.Int(),
		Radius: densityR.Int(),
	}
}
