package factoryShadow

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/shadow"
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
func NewShadow(

	color color.RGBA,
	blur,
	offsetX,
	offsetY float64,
	density interface{},
	iDensity iotmaker_platform_coordinate.IDensity,

) iotmaker_platform_IDraw.IFilterShadowInterface {

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

	return sd
}
