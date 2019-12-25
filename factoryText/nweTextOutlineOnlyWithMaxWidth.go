package factoryText

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform/abstractType/genericTypes"
	"github.com/helmutkemper/iotmaker.platform/abstractType/text"
)

func NewTextOutlineOnlyWithMaxWidth(platform iotmaker_platform_IDraw.IDraw, shadow iotmaker_platform_IDraw.IFilterShadowInterface, gradient iotmaker_platform_IDraw.IFilterGradientInterface, color interface{}, label string, x, y, maxWidth float64, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) text.Text {
	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.Set(x)
	x = densityCalc.Float64()

	densityCalc.Set(y)
	y = densityCalc.Float64()

	ik := genericTypes.Ink{}
	ik = genericTypes.NewInc(ik, 0, color, shadow, gradient, density, iDensity)

	tx := text.Text{
		Platform: platform,
		Ink:      ik,
		Label:    label,
		Stroke:   true,
		X:        x,
		Y:        y,
		MaxWidth: maxWidth,
	}

	tx.ConfigShadowPlatformAndFilter()
	tx.ConfigGradientPlatformAndFilter()
	tx.Create()

	return tx
}
