package basicBox

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
)

func NewBasicBox(platform iotmaker_platform_IDraw.IDraw, id string, x, y, width, height, border, lineWidth int, shadow iotmaker_platform_IDraw.IFilterShadowInterface, gradient iotmaker_platform_IDraw.IFilterGradientInterface, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) *BasicBox {
	dm := NewDimensions(x, y, width, height, border, density, iDensity)
	ik := NewInc(lineWidth, shadow, gradient, density, iDensity)

	bb := &BasicBox{
		Platform:   platform,
		Id:         id,
		Dimensions: dm,
		Ink:        ik,
	}

	bb.configShadowPlatformAndFilter()
	bb.configGradientPlatformAndFilter()
	bb.Create()

	return bb
}
