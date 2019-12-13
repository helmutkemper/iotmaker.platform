package factoryFont

import (
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/font"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/fontFamily"
)

func NewFont(size int, sizeUnit string, family fontFamily.FontFamily, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) font.Font {
	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.Set(size)
	size = densityCalc.Int()

	f := font.Font{
		Size:     size,
		SizeUnit: sizeUnit,
		Family:   family,
	}
	return f
}
