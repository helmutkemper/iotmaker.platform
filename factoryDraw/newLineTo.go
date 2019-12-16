package factoryDraw

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform/abstractType/draw"
)

func NewLineTo(platform iotmaker_platform_IDraw.IDraw, x1, y1, x2, y2, lineWidth int, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) draw.LineTo {
	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.Set(x1)
	x1 = densityCalc.Int()

	densityCalc.Set(x2)
	x2 = densityCalc.Int()

	densityCalc.Set(y1)
	y1 = densityCalc.Int()

	densityCalc.Set(y2)
	y2 = densityCalc.Int()

	densityCalc.Set(lineWidth)
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
