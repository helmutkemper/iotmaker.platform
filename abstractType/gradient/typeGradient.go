package gradient

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/point"
)

type Gradient struct {
	Platform     iotmaker_platform_IDraw.ICanvasGradient
	Type         Select
	CoordinateP0 interface{}
	CoordinateP1 interface{}
	ColorList    []ColorStop
}

func (el *Gradient) SetP0(point interface{}) {
	el.CoordinateP0 = point
}

func (el *Gradient) SetP1(point interface{}) {
	el.CoordinateP1 = point
}

func (el *Gradient) PrepareFilter(platform iotmaker_platform_IDraw.ICanvasGradient) {
	el.Platform = platform

	if el.Platform == nil {
		return
	}

	var x0, x1, y0, y1, r0, r1 interface{}

	switch pConvertedType := el.CoordinateP0.(type) {
	case point.Point:
		x0 = pConvertedType.X
		y0 = pConvertedType.Y

	case point.PointWithRadius:
		x0 = pConvertedType.X
		y0 = pConvertedType.Y
		r0 = pConvertedType.Radius
	}

	switch pConvertedType := el.CoordinateP1.(type) {
	case point.Point:
		x1 = pConvertedType.X
		y1 = pConvertedType.Y

	case point.PointWithRadius:
		x1 = pConvertedType.X
		y1 = pConvertedType.Y
		r1 = pConvertedType.Radius
	}

	var gradient interface{}

	if el.Type == KLinearFill || el.Type == KLinearStroke || el.Type == KLinearFillAndStroke {
		gradient = el.Platform.CreateLinearGradient(x0, y0, x1, y1)
	} else if el.Type == KRadialFill || el.Type == KRadialStroke || el.Type == KRadialFillAndStroke {
		gradient = el.Platform.CreateRadialGradient(x0, y0, r0, x1, y1, r1)
	}

	for _, value := range el.ColorList {
		el.Platform.AddColorStopPosition(gradient, value.Stop, value.Color)
	}

	if el.Type == KLinearFill || el.Type == KLinearFillAndStroke || el.Type == KRadialFill || el.Type == KRadialFillAndStroke {
		el.Platform.SetFillStyle(gradient)
		el.Platform.Fill()
	}

	if el.Type == KLinearStroke || el.Type == KLinearFillAndStroke || el.Type == KRadialStroke || el.Type == KRadialFillAndStroke {
		el.Platform.SetStrokeStyle(gradient)
		el.Platform.Stroke()
	}

}

func (el *Gradient) Fill(gradient interface{}) {
	el.Platform.SetFillStyle(gradient)
}

func (el *Gradient) Stroke(gradient interface{}) {
	el.Platform.SetStrokeStyle(gradient)
}
