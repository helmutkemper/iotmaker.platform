package factoryInk

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/ink"
)

func NewTextInkWothShadowAndGradient(fillColor interface{}, shadow iotmaker_platform_IDraw.IFilterShadowInterface, gradient iotmaker_platform_IDraw.IFilterGradientInterface) ink.Ink {
	var inkObj = ink.Ink{}

	inkObj.Set(1, fillColor, shadow, gradient)
	return inkObj
}
