package gradient

import iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"

type Gradient struct {
	Platform   iotmaker_platform_IDraw.IGradient
	Type       Select
	Coordinate Coordinate
	ColorList  []ColorStop
}

func (el *Gradient) PrepareGradientAndMountColorListFilter() {
	x0 := el.Coordinate.X
	x1 := el.Coordinate.X + el.Coordinate.Width

	y0 := el.Coordinate.Y
	y1 := el.Coordinate.Y + el.Coordinate.Height

	var gradient interface{}

	if el.Type == KLinearGradientFill || el.Type == KLinearGradientStroke {
		gradient = el.Platform.CreateLinearGradient(x0, y0, x1, y1)
	} else if el.Type == KRadialGradientFill || el.Type == KRadialGradientStroke {
		//TODO: radial gradient
		gradient = el.Platform.CreateLinearGradient(x0, y0, x1, y1)
	}

	for _, value := range el.ColorList {
		el.Platform.AddColorStop(gradient, value.Stop, value.Color)
	}

	if el.Type == KLinearGradientFill || el.Type == KRadialGradientFill {
		el.Platform.FillStyle(gradient)
	} else if el.Type == KLinearGradientStroke || el.Type == KRadialGradientStroke {
		el.Platform.StrokeStyle(gradient)
	}

}

func (el *Gradient) Fill(gradient interface{}) {
	el.Platform.FillStyle(gradient)
}

func (el *Gradient) Stroke(gradient interface{}) {
	el.Platform.StrokeStyle(gradient)
}
