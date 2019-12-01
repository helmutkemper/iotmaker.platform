package shadow

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
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
//    shadowFilter := shadow.NewShadowFilter(color, blur, offsetX, offsetY)
func NewShadowFilter(color color.RGBA, blur, offsetX, offsetY int) iotmaker_platform_IDraw.IFilterShadowInterface {

	if reflect.DeepEqual(color, color.RGBA) {
		return &Shadow{}
	}

	sd := &Shadow{
		Color:   color,
		Blur:    blur,
		OffsetX: offsetX,
		OffsetY: offsetY,
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
