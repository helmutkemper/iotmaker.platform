package factoryDraw

import (
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform/abstractType/draw"
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
)

func NewLineTo(platform iotmaker_platform_IDraw.IDraw, x1, y1, x2, y2, lineWidth float64, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) draw.LineTo {
	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.SetInt(x1)
	x1 = densityCalc.Int()

	densityCalc.SetInt(x2)
	x2 = densityCalc.Int()

	densityCalc.SetInt(y1)
	y1 = densityCalc.Int()

	densityCalc.SetInt(y2)
	y2 = densityCalc.Int()

	densityCalc.SetInt(lineWidth)
	lineWidth = densityCalc.Int()

	lt := draw.LineTo{
		Platform:  platform,
		Density:   density,
		X1:        x1,
		X2:        x2,
		Y1:        y1,
		Y2:        y2,
		LineWidth: lineWidth,
	}
	lt.Create()

	return lt
}
