package genericTypes

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
)

func NewInc(ink Ink, lineWidth float64, fillColor interface{}, shadow iotmaker_platform_IDraw.IFilterShadowInterface, gradient iotmaker_platform_IDraw.IFilterGradientInterface, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) Ink {
	densityCalc := iDensity
	densityCalc.Set(lineWidth)
	densityCalc.SetDensityFactor(density)

	ink.Set(densityCalc.Float64(), fillColor, shadow, gradient)
	return ink
}
