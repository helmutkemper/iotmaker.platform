package factoryInk

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/ink"
)

func NewTextInk(fillColor interface{}) ink.Ink {
	var inkObj = ink.Ink{}

	inkObj.Set(1, fillColor, nil, nil)
	return inkObj
}
