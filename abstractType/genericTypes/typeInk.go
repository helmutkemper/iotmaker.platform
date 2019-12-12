package genericTypes

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	"image/color"
)

type Ink struct {
	lastInteractionLineWidth int
	LineWidth                int
	Shadow                   iotmaker_platform_IDraw.IFilterShadowInterface
	Gradient                 iotmaker_platform_IDraw.IFilterGradientInterface
	Color                    color.RGBA
	hasModified              bool
}

func (el *Ink) Set(lineWidth int, fillColor interface{}, shadow iotmaker_platform_IDraw.IFilterShadowInterface, gradient iotmaker_platform_IDraw.IFilterGradientInterface) {
	if el.lastInteractionLineWidth != lineWidth {
		el.lastInteractionLineWidth = lineWidth
		el.hasModified = true
	} else {
		el.hasModified = false
	}

	if fillColor != nil {
		el.Color = fillColor.(color.RGBA)
	}
	el.LineWidth = lineWidth
	el.Shadow = shadow
	el.Gradient = gradient
}
