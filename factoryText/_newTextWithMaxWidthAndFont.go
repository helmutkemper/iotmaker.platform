package factoryText

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/font"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/text"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryInk"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/ink"
)

func NewTextWithMaxWidthAndFont(platform iotmaker_platform_IDraw.IDraw, shadow iotmaker_platform_IDraw.IFilterShadowInterface, color interface{}, gradient iotmaker_platform_IDraw.IFilterGradientInterface, labelFont font.Font, label string, x, y, maxWidth float64, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) text.Text {
	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.Set(x)
	x = densityCalc.Float64()

	densityCalc.Set(y)
	y = densityCalc.Float64()

	ik := ink.Ink{}
	ik = factoryInk.NewInk(ik, 0, color, shadow, gradient, density, iDensity)

	tx := text.Text{
		Platform: platform,
		Font:     labelFont,
		Ink:      ik,
		Label:    label,
		Fill:     true,
		X:        x,
		Y:        y,
		MaxWidth: maxWidth,
	}

	tx.ConfigShadowPlatformAndFilter()
	tx.ConfigGradientPlatformAndFilter()
	tx.Create()

	return tx
}
