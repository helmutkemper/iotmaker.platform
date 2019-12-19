package draw

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"image/color"
)

type ChartLinear struct {
	Platform iotmaker_platform_IDraw.IDraw

	X      int
	Y      int
	Width  int
	Height int

	LineColor color.RGBA

	XAxis []int
	YAxis []int

	XAxisLine  float64
	YAxisLine  float64
	XAxisColor color.RGBA
	YAxisColor color.RGBA

	XSubAxisLine  float64
	YSubAxisLine  float64
	XSubAxisColor color.RGBA
	YSubAxisColor color.RGBA

	Density  interface{}
	IDensity iotmaker_platform_coordinate.IDensity
}

func (el *ChartLinear) Create() {

	// axis x
	el.lineTo(el.X, el.Y, el.X, el.Y+el.Height, el.XAxisLine, el.XAxisColor)

	// axis y
	el.lineTo(el.X, el.Y+el.Height, el.X+el.Width, el.Y+el.Height, el.YAxisLine, el.YAxisColor)

	for _, v := range el.XAxis {
		el.lineTo(v, el.Y, v, el.Y+el.Height, el.XSubAxisLine, el.XSubAxisColor)
	}

	for _, v := range el.YAxis {
		el.lineTo(el.X, v, el.X+el.Width, v, el.YSubAxisLine, el.YSubAxisColor)
	}

	el.Platform.ResetLineWidth()
	el.Platform.ResetFillStyle()
	el.Platform.ResetStrokeStyle()
	el.Platform.ResetShadow()
}

func (el *ChartLinear) lineTo(x1, y1, x2, y2 int, lineWidth interface{}, color interface{}) {
	el.Platform.BeginPath()
	el.Platform.SetStrokeStyle(color)
	el.Platform.SetLineWidth(lineWidth)
	el.Platform.MoveTo(x1, y1)
	el.Platform.LineTo(x2, y2)
	el.Platform.ClosePath(x1, y1)
	el.Platform.Stroke()
}

func (el *ChartLinear) Begin(x, y int) {
	el.Platform.BeginPath()
	el.Platform.SetLineWidth(0.005)
	el.Platform.MoveTo(x, y)
}

func (el *ChartLinear) Pixel(x, y int) {
	el.Platform.SetLineWidth(0.01)
	el.Platform.MoveTo(x-1, y-1)
	el.Platform.LineTo(x, y)
	el.Platform.Stroke()
}

func (el *ChartLinear) End() {
	el.Platform.Stroke()
}
