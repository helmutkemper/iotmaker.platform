package gradient

import iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"

// en: Creates a set of coordinates, taking into account the size of the screen.
//
// pt_br: Cria um par de coordenadas levando em consideração a densidade da tela.
//   x:        horizontal point of the screen.
//   y:        vertical point of the screen.
//   radius:   radius is a width radius, not a degrees angle
//   density:  density factor of the screen.
//   iDensity: density interface
func NewPointWithRadius(x, y, radius int, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) PointWithRadius {

	densityX := iDensity
	densityX.Set(x)
	densityX.SetDensityFactor(density)

	densityY := iDensity
	densityY.Set(y)
	densityY.SetDensityFactor(density)

	densityR := iDensity
	densityR.Set(radius)
	densityR.SetDensityFactor(density)

	return PointWithRadius{
		X:      densityX.Int(),
		Y:      densityY.Int(),
		Radius: densityR.Int(),
	}
}
