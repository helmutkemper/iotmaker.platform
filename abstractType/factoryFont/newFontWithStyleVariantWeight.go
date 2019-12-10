package factoryFont

import (
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/font"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/fontFamily"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/fontStyle"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/fontVariant"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/fontWeight"
)

func NewFontWithStyleVariantWeight(size int, sizeUnit string, family fontFamily.FontFamily, style fontStyle.FontStyle, variant fontVariant.FontVariant, weight fontWeight.FontWeight, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) font.Font {
	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.Set(size)
	size = densityCalc.Int()

	f := font.Font{
		Size:     size,
		SizeUnit: sizeUnit,
		Family:   family,
		Style:    style,
		Variant:  variant,
		Weight:   weight,
	}
	return f
}
