package factoryDraw

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/draw"
)

func NewLineTo(platform iotmaker_platform_IDraw.IDraw, x1, y1, x2, y2, lineWidth float64, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) draw.LineTo {
	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.Set(x1)
	x1 = densityCalc.Float()

	densityCalc.Set(x2)
	x2 = densityCalc.Float()

	densityCalc.Set(y1)
	y1 = densityCalc.Float()

	densityCalc.Set(y2)
	y2 = densityCalc.Float()

	densityCalc.Set(lineWidth)
	lineWidth = densityCalc.Float()

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
