package factoryDimensions

import iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"

func NewDimensions(dm Dimensions, x, y, width, height, border float64, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) Dimensions {

	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.SetInt(x)
	x = densityCalc.Int()

	densityCalc.SetInt(y)
	y = densityCalc.Int()

	densityCalc.SetInt(width)
	width = densityCalc.Int()

	densityCalc.SetInt(height)
	height = densityCalc.Int()

	densityCalc.SetInt(border)
	border = densityCalc.Int()

	dm.Set(x, y, width, height, border)

	return dm
}
