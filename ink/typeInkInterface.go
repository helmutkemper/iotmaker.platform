package ink

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/point"
)

type Interface interface {
	Set(lineWidth int, fillColor interface{}, shadow iotmaker_platform_IDraw.IFilterShadowInterface, gradient iotmaker_platform_IDraw.IFilterGradientInterface)
	SetGradientLinearRectangle(p0, p1 point.Point)
	ShadowPrepareFilter(platform iotmaker_platform_IDraw.ICanvasShadow)
	GradientPrepareFilter(platform iotmaker_platform_IDraw.ICanvasGradient)
}
