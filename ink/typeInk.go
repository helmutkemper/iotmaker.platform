package ink

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
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

func (el *Ink) ShadowPrepareFilter(platform iotmaker_platform_IDraw.ICanvasShadow) {

	if el.Shadow == nil || platform == nil {
		return
	}

	el.Shadow.PrepareFilter(platform)
}

func (el *Ink) GradientPrepareFilter(platform iotmaker_platform_IDraw.ICanvasGradient) {

	if el.Gradient == nil || platform == nil {
		return
	}

	el.Gradient.PrepareFilter(platform)
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
