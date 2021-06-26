package factoryInk

import (
	global "github.com/helmutkemper/iotmaker.santa_isabel_theater.globalConfig"
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/ink"
)

func NewInkWothShadowAndGradient(lineWidth int, fillColor interface{}, shadow iotmaker_platform_IDraw.IFilterShadowInterface, gradient iotmaker_platform_IDraw.IFilterGradientInterface) ink.Ink {
	var inkObj = ink.Ink{}

	densityCalc := global.Global.DensityManager
	densityCalc.SetInt(lineWidth)
	densityCalc.SetDensityFactor(global.Global.Density)

	inkObj.Set(densityCalc.Int(), fillColor, shadow, gradient)
	return inkObj
}
