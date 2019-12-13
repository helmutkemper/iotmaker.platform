package factoryGradient

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	"github.com/helmutkemper/iotmaker.platform/abstractType/gradient"
	"github.com/helmutkemper/iotmaker.platform/abstractType/point"
)

func SetGlobalGradientLinearToStroke(platform iotmaker_platform_IDraw.IDraw, coordinateP0, coordinateP1 point.Point, colorList []gradient.ColorStop) iotmaker_platform_IDraw.IFilterGradientInterface {
	gd := &gradient.Gradient{
		Type:         gradient.KLinearStroke,
		CoordinateP0: coordinateP0,
		CoordinateP1: coordinateP1,
		ColorList:    colorList,
	}
	gd.SetGlobal(platform)

	return gd
}
