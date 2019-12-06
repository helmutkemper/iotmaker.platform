package gradient

import iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"

type Gradient struct {
	Platform     iotmaker_platform_IDraw.ICanvasGradient
	Type         Select
	CoordinateP0 interface{}
	CoordinateP1 interface{}
	ColorList    []ColorStop
}

func (el *Gradient) PrepareFilter(platform iotmaker_platform_IDraw.ICanvasGradient) {
	el.Platform = platform

	if el.Platform == nil {
		return
	}

	var x0, x1, y0, y1, r0, r1 interface{}

	switch pConvertedType := el.CoordinateP0.(type) {
	case Point:
		x0 = pConvertedType.X
		y0 = pConvertedType.Y

	case PointWithRadius:
		x0 = pConvertedType.X
		y0 = pConvertedType.Y
		r0 = pConvertedType.Radius
	}

	switch pConvertedType := el.CoordinateP1.(type) {
	case Point:
		x1 = pConvertedType.X
		y1 = pConvertedType.Y

	case PointWithRadius:
		x1 = pConvertedType.X
		y1 = pConvertedType.Y
		r1 = pConvertedType.Radius
	}

	var gradient interface{}

	if el.Type == KLinearGradientFill || el.Type == KLinearGradientStroke || el.Type == KLinearGradientFillAndStroke {
		gradient = el.Platform.CreateLinearGradient(x0, y0, x1, y1)
	} else if el.Type == KRadialGradientFill || el.Type == KRadialGradientStroke || el.Type == KRadialGradientFillAndStroke {
		gradient = el.Platform.CreateRadialGradient(x0, y0, r0, x1, y1, r1)
	}

	for _, value := range el.ColorList {
		el.Platform.AddColorStopPosition(gradient, value.Stop, value.Color)
	}

	if el.Type == KLinearGradientFill || el.Type == KLinearGradientFillAndStroke || el.Type == KRadialGradientFill || el.Type == KRadialGradientFillAndStroke {
		el.Platform.SetFillStyle(gradient)
		el.Platform.Fill()
	}

	if el.Type == KLinearGradientStroke || el.Type == KLinearGradientFillAndStroke || el.Type == KRadialGradientStroke || el.Type == KRadialGradientFillAndStroke {
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
