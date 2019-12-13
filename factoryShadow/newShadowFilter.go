package factoryShadow

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform/abstractType/shadow"
	"image/color"
	"reflect"
)

// en: creates a shadow effect over the canvas draw image
//
// pt_br: cria o efeito de sombra sobre o desenho
//
// Example:
//    color        := colornames.DarkredHalfTransparent
//    blur         := 5
//    offsetX      := 2
//    offsetY      := 2
//    density      := 1.0
//    iDensity     := &Density{}
//    shadowFilter := shadow.NewShadowFilter(color, blur, offsetX, offsetY)
func NewShadowFilter(color color.RGBA, blur, offsetX, offsetY int, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) iotmaker_platform_IDraw.IFilterShadowInterface {

	if reflect.DeepEqual(color, color.RGBA) {
		return &shadow.Shadow{}
	}

	densityOffsetX := iDensity
	densityOffsetX.Set(offsetX)
	densityOffsetX.SetDensityFactor(density)

	densityOffsetY := iDensity
	densityOffsetY.Set(offsetY)
	densityOffsetY.SetDensityFactor(density)

	sd := &shadow.Shadow{
		Color:   color,
		Blur:    blur,
		OffsetX: densityOffsetX.Int(),
		OffsetY: densityOffsetY.Int(),
	}

	sd.ColorEnable = true

	if blur != 0 {
		sd.BlurEnable = true
	}

	if offsetX != 0 {
		sd.OffsetXEnable = true
	}

	if offsetY != 0 {
		sd.OffsetYEnable = true
	}

	return sd
}
