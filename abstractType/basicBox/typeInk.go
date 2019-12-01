package basicBox

import iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"

type Ink struct {
	LineWidth int
	Shadow    iotmaker_platform_IDraw.IFilterShadowInterface
	Gradient  iotmaker_platform_IDraw.IFilterGradientInterface
}
