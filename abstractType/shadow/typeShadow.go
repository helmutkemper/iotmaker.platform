package shadow

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	"image/color"
)

// en: Please, usa a  function shadow.NewShadowFilter() to make a new filter
//
// pt_br: Por favor, use a função shadow.NewShadowFilter() para montar um novo filtro
type Shadow struct {
	Platform      iotmaker_platform_IDraw.ICanvasShadow
	Color         color.RGBA
	ColorEnable   bool
	Blur          int
	BlurEnable    bool
	OffsetX       int
	OffsetXEnable bool
	OffsetY       int
	OffsetYEnable bool
}

// en: Please, usa a  function shadow.NewShadowFilter() to make a new filter.
// receive a canvas specific platform draw shadow functions.
//
// pt_br: Por favor, use a função shadow.NewShadowFilter() para montar um novo filtro
// recebe as funções de desenho específicas do elemento canvas da plataforma
func (el *Shadow) PrepareFilter(platform iotmaker_platform_IDraw.ICanvasShadow) {

	el.Platform = platform

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
