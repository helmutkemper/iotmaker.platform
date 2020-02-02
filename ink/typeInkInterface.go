package ink

import iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"

type Interface interface {
	Set(lineWidth int, fillColor interface{}, shadow iotmaker_platform_IDraw.IFilterShadowInterface, gradient iotmaker_platform_IDraw.IFilterGradientInterface)
	ShadowPrepareFilter(platform iotmaker_platform_IDraw.ICanvasShadow)
	GradientPrepareFilter(platform iotmaker_platform_IDraw.ICanvasGradient)
}
