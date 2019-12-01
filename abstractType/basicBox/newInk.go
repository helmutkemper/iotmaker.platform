package basicBox

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
)

func NewInc(lineWidth int, shadow iotmaker_platform_IDraw.IFilterShadowInterface, gradient iotmaker_platform_IDraw.IFilterGradientInterface, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) Ink {
	densityCalc := iDensity
	densityCalc.Set(lineWidth)
	densityCalc.SetDensityFactor(density)

	return Ink{
		LineWidth: densityCalc.Int(),
		Shadow:    shadow,
		Gradient:  gradient,
	}
}
