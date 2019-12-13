package factoryBasicBox

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform/abstractType/basicBox"
	"github.com/helmutkemper/iotmaker.platform/abstractType/genericTypes"
)

func NewBasicBox(platform, scratchPad iotmaker_platform_IDraw.IDraw, color interface{}, id string, x, y, width, height, border, lineWidth int, shadow iotmaker_platform_IDraw.IFilterShadowInterface, gradient iotmaker_platform_IDraw.IFilterGradientInterface, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) *basicBox.BasicBox {
	dm := genericTypes.Dimensions{}
	dm = genericTypes.NewDimensions(dm, x, y, width, height, border, density, iDensity)

	ik := genericTypes.Ink{}
	ik = genericTypes.NewInc(ik, lineWidth, color, shadow, gradient, density, iDensity)

	bb := &basicBox.BasicBox{
		Platform:   platform,
		ScratchPad: scratchPad,
		Id:         id,
		Dimensions: dm,
		Ink:        ik,
	}

	bb.ConfigShadowPlatformAndFilter()
	bb.ConfigGradientPlatformAndFilter()
	bb.Create()

	return bb
}
