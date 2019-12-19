package factoryDraw

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform/abstractType/draw"
	"image/color"
)

func NewChartLinear(platform iotmaker_platform_IDraw.IDraw, x, y, width, height int, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) *draw.ChartLinear {
	d := &draw.ChartLinear{
		Platform: platform,

		X:      x,
		Y:      y,
		Width:  width,
		Height: height,

		LineColor: color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xFF},

		XAxisLine:  0.4,
		XAxisColor: color.RGBA{R: 0x00, G: 0xFF, B: 0x00, A: 0x10},

		XAxis: []int{110, 210, 310, 410, 510},
		YAxis: []int{110, 210, 310, 410},

		YAxisLine:  0.4,
		YAxisColor: color.RGBA{R: 0x00, G: 0xFF, B: 0x00, A: 0x10},

		XSubAxisLine:  0.4,
		XSubAxisColor: color.RGBA{R: 0xBF, G: 0xBF, B: 0xFF, A: 0x01},

		YSubAxisLine:  0.4,
		YSubAxisColor: color.RGBA{R: 0xBF, G: 0xBF, B: 0xFF, A: 0x01},
	}
	d.Create()

	return d
}
