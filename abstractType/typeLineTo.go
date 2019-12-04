package abstractType

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
)

func NewLineTo(platform iotmaker_platform_IDraw.IDraw, id string, density float64, x1, y1, x2, y2, lineWidth int) LineTo {

	coordinateX1 := iotmaker_platform_coordinate.Density{}
	coordinateX1.SetDensityFactor(density)
	coordinateX1.Set(x1)

	coordinateX2 := iotmaker_platform_coordinate.Density{}
	coordinateX2.SetDensityFactor(density)
	coordinateX2.Set(x2)

	coordinateY1 := iotmaker_platform_coordinate.Density{}
	coordinateY1.SetDensityFactor(density)
	coordinateY1.Set(y1)

	coordinateY2 := iotmaker_platform_coordinate.Density{}
	coordinateY2.SetDensityFactor(density)
	coordinateY2.Set(y2)

	coordinateLineWidth := iotmaker_platform_coordinate.Density{}
	coordinateLineWidth.SetDensityFactor(density)
	coordinateLineWidth.Set(lineWidth)

	lt := LineTo{
		Id:        id,
		Platform:  platform,
		Density:   density,
		X1:        coordinateX1,
		X2:        coordinateX2,
		Y1:        coordinateY1,
		Y2:        coordinateY2,
		LineWidth: coordinateLineWidth,
	}
	lt.Create()

	return lt
}

type LineTo struct {
	Platform  iotmaker_platform_IDraw.IDraw
	Density   float64
	Id        string
	X1        iotmaker_platform_coordinate.Density
	Y1        iotmaker_platform_coordinate.Density
	X2        iotmaker_platform_coordinate.Density
	Y2        iotmaker_platform_coordinate.Density
	LineWidth iotmaker_platform_coordinate.Density
}

func (el *LineTo) Create() {

	// correct the line width size
	//el.X.Add(el.SetLineWidth.Int() / 2)
	//el.Y.Add(el.SetLineWidth.Int() / 2)
	//el.Width.Sub(el.SetLineWidth.Int() / 2)
	//el.Height.Sub(el.SetLineWidth.Int() / 2)

	//el.SelfContext.SetLineWidth(input.SetLineWidth)

	el.Platform.BeginPath()
	el.Platform.SetLineWidth(el.LineWidth.Int())
	el.Platform.MoveTo(el.X1.Int(), el.Y1.Int()) // a
	el.Platform.LineTo(el.X2.Int(), el.Y2.Int()) // a->b
	el.Platform.Stroke()
}
