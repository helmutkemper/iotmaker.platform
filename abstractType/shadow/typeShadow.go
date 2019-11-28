package shadow

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	"image/color"
)

type Shadow struct {
	Platform      iotmaker_platform_IDraw.IShadow
	Color         color.RGBA
	ColorEnable   bool
	Blur          int
	BlurEnable    bool
	OffsetX       int
	OffsetXEnable bool
	OffsetY       int
	OffsetYEnable bool
}

func (el *Shadow) PrepareShadowFilter() {
	// the feature of the javascript itself
	if el.ColorEnable == false || el.Platform == nil {
		return
	}

	el.Platform.ShadowColor(el.Color)

	if el.BlurEnable == true {
		el.Platform.ShadowBlur(el.Blur)
	}

	if el.OffsetXEnable == true {
		el.Platform.ShadowOffsetX(el.OffsetX)
	}

	if el.OffsetYEnable == true {
		el.Platform.ShadowOffsetY(el.OffsetY)
	}
}
