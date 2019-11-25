package abstractType

import (
	iotmaker_platform "github.com/helmutkemper/iotmaker.platform"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
)

func NewLineTo(platform iotmaker_platform.ICanvas, id string, density float64, x1, y1, x2, y2 int) LineTo {

	coordinateX1 := iotmaker_platform_coordinate.Coordinate{}
	coordinateX1.SetDensityFactor(density)
	coordinateX1.Set(x1)

	coordinateX2 := iotmaker_platform_coordinate.Coordinate{}
	coordinateX2.SetDensityFactor(density)
	coordinateX2.Set(x2)

	coordinateY1 := iotmaker_platform_coordinate.Coordinate{}
	coordinateY1.SetDensityFactor(density)
	coordinateY1.Set(y1)

	coordinateY2 := iotmaker_platform_coordinate.Coordinate{}
	coordinateY2.SetDensityFactor(density)
	coordinateY2.Set(y2)

	lt := LineTo{
		Id:       id,
		Platform: platform,
		Density:  density,
		X1:       coordinateX1,
		X2:       coordinateX2,
		Y1:       coordinateY1,
		Y2:       coordinateY2,
	}
	lt.Create()

	return lt
}

type LineTo struct {
	Platform iotmaker_platform.ICanvas
	Density  float64
	Id       string
	X1       iotmaker_platform_coordinate.Coordinate
	Y1       iotmaker_platform_coordinate.Coordinate
	X2       iotmaker_platform_coordinate.Coordinate
	Y2       iotmaker_platform_coordinate.Coordinate
}

func (el *LineTo) Create() {

	// correct the line width size
	//el.X.Add(el.LineWidth.Int() / 2)
	//el.Y.Add(el.LineWidth.Int() / 2)
	//el.Width.Sub(el.LineWidth.Int() / 2)
	//el.Height.Sub(el.LineWidth.Int() / 2)

	//el.SelfContext.LineWidth(input.LineWidth)

	el.Platform.BeginPath()
	el.Platform.MoveTo(el.X1.Int(), el.Y1.Int()) // a
	el.Platform.LineTo(el.X2.Int(), el.Y2.Int()) // a->b
	el.Platform.Stroke()
}
