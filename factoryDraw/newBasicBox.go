package factoryDraw

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform/abstractType/draw"
	"github.com/helmutkemper/iotmaker.platform/abstractType/genericTypes"
)

func NewBasicBox(
	platform,
	scratchPad iotmaker_platform_IDraw.IDraw,
	color interface{},
	id string,
	x,
	y,
	width,
	height,
	border,
	lineWidth float64,
	shadow iotmaker_platform_IDraw.IFilterShadowInterface,
	gradient iotmaker_platform_IDraw.IFilterGradientInterface,
	density interface{},
	iDensity iotmaker_platform_coordinate.IDensity,
) *draw.BasicBox {

	dm := genericTypes.Dimensions{}
	dm = genericTypes.NewDimensions(dm, x, y, width, height, border, density, iDensity)

	ik := genericTypes.Ink{}
	ik = genericTypes.NewInc(ik, lineWidth, color, shadow, gradient, density, iDensity)

	bb := &draw.BasicBox{
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
