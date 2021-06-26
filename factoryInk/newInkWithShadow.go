package factoryInk

import (
	global "github.com/helmutkemper/iotmaker.santa_isabel_theater.globalConfig"
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/ink"
)

func NewInkWithShadow(lineWidth int, fillColor interface{}, shadow iotmaker_platform_IDraw.IFilterShadowInterface) ink.Ink {
	var inkObj = ink.Ink{}

	densityCalc := global.Global.DensityManager
	densityCalc.SetInt(lineWidth)
	densityCalc.SetDensityFactor(global.Global.Density)

	inkObj.Set(densityCalc.Int(), fillColor, shadow, nil)
	return inkObj
}
