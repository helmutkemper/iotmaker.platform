package shadow

import (
	"image/color"
	"reflect"
)

func NewShadow(color color.RGBA, blur, offsetX, offsetY int) Shadow {

	if reflect.DeepEqual(color, color.RGBA) {
		return Shadow{}
	}

	sd := Shadow{
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
