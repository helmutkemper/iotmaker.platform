package draw

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
)

type LineTo struct {
	Platform  iotmaker_platform_IDraw.IDraw
	Density   interface{}
	X1        float64
	Y1        float64
	X2        float64
	Y2        float64
	LineWidth float64
}

func (el *LineTo) Create() {

	el.Platform.BeginPath()
	el.Platform.SetLineWidth(el.LineWidth)
	el.Platform.MoveTo(el.X1, el.Y1) // a
	el.Platform.LineTo(el.X2, el.Y2) // a->b
	el.Platform.Stroke()
}
