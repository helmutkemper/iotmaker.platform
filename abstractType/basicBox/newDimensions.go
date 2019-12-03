package basicBox

import iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"

func NewDimensions(dm Dimensions, x, y, width, height, border int, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) Dimensions {

	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.Set(x)
	x = densityCalc.Int()

	densityCalc.Set(y)
	y = densityCalc.Int()

	densityCalc.Set(width)
	width = densityCalc.Int()

	densityCalc.Set(height)
	height = densityCalc.Int()

	densityCalc.Set(border)
	border = densityCalc.Int()

	dm.Set(x, y, width, height, border)

	return dm
}
