package factoryInk

import (
	global "github.com/helmutkemper/iotmaker.santa_isabel_theater.globalConfig"
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/ink"
)

func NewInkWothGradient(lineWidth int, fillColor interface{}, gradient iotmaker_platform_IDraw.IFilterGradientInterface) ink.Ink {
	var inkObj = ink.Ink{}

	densityCalc := global.Global.DensityManager
	densityCalc.SetInt(lineWidth)
	densityCalc.SetDensityFactor(global.Global.Density)

	inkObj.Set(densityCalc.Int(), fillColor, nil, gradient)
	return inkObj
}
