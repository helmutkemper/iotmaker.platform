package factoryDimensions

import (
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/dimensions"
)

// fixme:

func NewDimensions(dm dimensions.Dimensions, x, y, width, height, border float64, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) dimensions.Dimensions {

	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.Set(x)
	x = densityCalc.Float()

	densityCalc.Set(y)
	y = densityCalc.Float()

	densityCalc.Set(width)
	width = densityCalc.Float()

	densityCalc.Set(height)
	height = densityCalc.Float()

	densityCalc.Set(border)
	border = densityCalc.Float()

	//dm.Set(x, y, width, height, border)

	return dm
}
