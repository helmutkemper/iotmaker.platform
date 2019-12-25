package independentDraw

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
)

// en: This function create a path of a box with rounded corners into the platform using density
//
// pt_br: Esta função cria o caminho de um quadrado com cantos arredondados na plataforma usando densidade
func DrawBoxWithRoundedCornersIntoThePathWithDensity(platform iotmaker_platform_IDraw.IDraw, x, y, width, height, border float64, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) {
	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.Set(x)
	x = densityCalc.Float64()

	densityCalc.Set(y)
	y = densityCalc.Float64()

	densityCalc.Set(width)
	width = densityCalc.Float64()

	densityCalc.Set(height)
	height = densityCalc.Float64()

	densityCalc.Set(border)
	border = densityCalc.Float64()

	DrawBoxWithRoundedCornersIntoThePath(platform, x, y, width, height, border)
}
