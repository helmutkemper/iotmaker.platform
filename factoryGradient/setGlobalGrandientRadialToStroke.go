package factoryGradient

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	"github.com/helmutkemper/iotmaker.platform/abstractType/gradient"
	"github.com/helmutkemper/iotmaker.platform/abstractType/point"
)

func SetGlobalGradientRadialToStroke(platform iotmaker_platform_IDraw.IDraw, coordinateP0, coordinateP1 point.PointWithRadius, colorList []gradient.ColorStop) iotmaker_platform_IDraw.IFilterGradientInterface {
	gd := &gradient.Gradient{
		Type:         gradient.KLinearFill,
		CoordinateP0: coordinateP0,
		CoordinateP1: coordinateP1,
		ColorList:    colorList,
	}
	gd.SetGlobal(platform)

	return gd
}
