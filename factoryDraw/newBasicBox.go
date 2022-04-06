package factoryDraw

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/draw"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/dimensions"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryDimensions"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryInk"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/ink"
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

	dm := dimensions.Dimensions{}
	dm = factoryDimensions.NewDimensions(dm, x, y, width, height, border, density, iDensity)

	ik := ink.Ink{}
	ik = factoryInk.NewInkWithShadowAndGradient(int(lineWidth), color, shadow, gradient)

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
