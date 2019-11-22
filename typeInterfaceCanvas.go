package iotmaker_platform

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
)

type ICanvas interface {
	BeginPath()
	MoveTo(x, y iotmaker_types.Coordinate)
	ArcTo(x, y, radius, startAngle, endAngle iotmaker_types.Coordinate)
	LineTo(x, y iotmaker_types.Coordinate)
	ClosePath(x, y iotmaker_types.Coordinate)
}
