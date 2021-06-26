package factoryInk

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/ink"
)

func NewTextInkWithShadow(fillColor interface{}, shadow iotmaker_platform_IDraw.IFilterShadowInterface) ink.Ink {
	var inkObj = ink.Ink{}

	inkObj.Set(1, fillColor, shadow, nil)
	return inkObj
}
