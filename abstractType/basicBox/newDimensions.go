package basicBox

import iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"

func NewDimensions(x, y, width, height, border int, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) Dimensions {

	dm := Dimensions{}

	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.Set(x)
	dm.X = densityCalc.Int()

	densityCalc.Set(y)
	dm.Y = densityCalc.Int()

	densityCalc.Set(width)
	dm.Width = densityCalc.Int()

	densityCalc.Set(height)
	dm.Height = densityCalc.Int()

	densityCalc.Set(border)
	dm.Border = densityCalc.Int()

	return dm
}
