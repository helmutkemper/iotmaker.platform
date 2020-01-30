package factoryInk

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/ink"
)

func NewInk(ink ink.Ink, lineWidth int, fillColor interface{}, shadow iotmaker_platform_IDraw.IFilterShadowInterface, gradient iotmaker_platform_IDraw.IFilterGradientInterface, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) ink.Ink {
	densityCalc := iDensity
	densityCalc.SetInt(lineWidth)
	densityCalc.SetDensityFactor(density)

	ink.Set(densityCalc.Int(), fillColor, shadow, gradient)
	return ink
}
