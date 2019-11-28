package gradient

import iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"

type Gradient struct {
	Platform   iotmaker_platform_IDraw.IGradient
	Type       int
	Coordinate Coordinate
	ColorList  []ColorStop
}

func (el *Gradient) PrepareGradientAndMountColorListFilter() {
	x0 := el.Coordinate.X
	x1 := el.Coordinate.X + el.Coordinate.Width

	y0 := el.Coordinate.Y
	y1 := el.Coordinate.Y + el.Coordinate.Height

	var gradient interface{}

	//TODO: radial gradient
	if el.Type == 0 {
		gradient = el.Platform.CreateLinearGradient(x0, y0, x1, y1)
	} else {
		//FIXME: mudar para radial
		gradient = el.Platform.CreateLinearGradient(x0, y0, x1, y1)
	}

	for _, value := range el.ColorList {
		el.Platform.AddColorStop(gradient, value.Stop, value.Color)
	}
}

func (el *Gradient) Fill(gradient interface{}) {
	el.Platform.FillStyle(gradient)
}

func (el *Gradient) Stroke(gradient interface{}) {
	el.Platform.StrokeStyle(gradient)
}
