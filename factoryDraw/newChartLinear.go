package factoryDraw

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform/abstractType/draw"
)

func NewChartLinear(platform iotmaker_platform_IDraw.IDraw, x, y, width, height int, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) *draw.ChartLinear {
	d := &draw.ChartLinear{
		Platform: platform,

		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
	d.Create()

	return d
}
