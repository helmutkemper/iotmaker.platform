package factoryInk

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/ink"
)

func NewTextInkWothGradient(fillColor interface{}, gradient iotmaker_platform_IDraw.IFilterGradientInterface) ink.Ink {
	var inkObj = ink.Ink{}

	inkObj.Set(1, fillColor, nil, gradient)
	return inkObj
}
