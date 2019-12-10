package factoryText

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/font"
)

func NewText(platform iotmaker_platform_IDraw.IDraw, text string, font font.Font, x, y int, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) {
	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.Set(x)
	x = densityCalc.Int()

	densityCalc.Set(y)
	y = densityCalc.Int()

	platform.Font(font)
	platform.FillText(text, x, y)
}
