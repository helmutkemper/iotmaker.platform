package factorySimpleBox

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/simple"
)

// en: This function create a path of a box with rounded corners into the platform using density
//
// pt_br: Esta função cria o caminho de um quadrado com cantos arredondados na plataforma usando densidade
func NewBoxWithRoundedCorners(platform iotmaker_platform_IDraw.IDraw, x, y, width, height, border int, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) simple.BoxWithRoundedCorners {
	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.SetInt(x)
	x = densityCalc.Int()

	densityCalc.SetInt(y)
	y = densityCalc.Int()

	densityCalc.SetInt(width)
	width = densityCalc.Int()

	densityCalc.SetInt(height)
	height = densityCalc.Int()

	densityCalc.SetInt(border)
	border = densityCalc.Int()

	rect := simple.BoxWithRoundedCorners{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
		Border: border,
	}
	rect.Init()
	rect.Draw()

	return rect
}
