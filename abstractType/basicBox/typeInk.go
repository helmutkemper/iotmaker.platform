package basicBox

import iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"

type Ink struct {
	lastInteractionLineWidth int
	LineWidth                int
	Shadow                   iotmaker_platform_IDraw.IFilterShadowInterface
	Gradient                 iotmaker_platform_IDraw.IFilterGradientInterface
	hasModified              bool
}

func (el *Ink) Set(lineWidth int, shadow iotmaker_platform_IDraw.IFilterShadowInterface, gradient iotmaker_platform_IDraw.IFilterGradientInterface) {
	if el.lastInteractionLineWidth != lineWidth {
		el.lastInteractionLineWidth = lineWidth
		el.hasModified = true
	} else {
		el.hasModified = false
	}

	el.LineWidth = lineWidth
	el.Shadow = shadow
	el.Gradient = gradient
}
