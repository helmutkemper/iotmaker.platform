package factoryDimensions

import iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"

func NewDimensions(dm Dimensions, x, y, width, height, border float64, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) Dimensions {

	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.Set(x)
	x = densityCalc.Float64()

	densityCalc.Set(y)
	y = densityCalc.Float64()

	densityCalc.Set(width)
	width = densityCalc.Float64()

	densityCalc.Set(height)
	height = densityCalc.Float64()

	densityCalc.Set(border)
	border = densityCalc.Float64()

	dm.Set(x, y, width, height, border)

	return dm
}
