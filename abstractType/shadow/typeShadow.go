package shadow

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	"image/color"
	"reflect"
)

// en: Please, usa a  function shadow.NewShadowFilter() to make a new filter
//
// pt_br: Por favor, use a função shadow.NewShadowFilter() para montar um novo filtro
type Shadow struct {
	Platform iotmaker_platform_IDraw.ICanvasShadow
	Color    color.RGBA
	Blur     float64
	OffsetX  int
	OffsetY  int
}

// en: Please, usa a  function shadow.NewShadowFilter() to make a new filter.
// receive a canvas specific platform draw shadow functions.
//
// pt_br: Por favor, use a função shadow.NewShadowFilter() para montar um novo filtro
// recebe as funções de desenho específicas do elemento canvas da plataforma
func (el *Shadow) PrepareFilter(platform iotmaker_platform_IDraw.ICanvasShadow) {

	el.Platform = platform

	if el.Platform == nil || reflect.DeepEqual(color.RGBA{}, el.Color) == true {
		return
	}

	el.Platform.SetShadowColor(el.Color)
	el.Platform.SetShadowBlur(el.Blur)
	el.Platform.ShadowOffsetX(el.OffsetX)
	el.Platform.ShadowOffsetY(el.OffsetY)
}
