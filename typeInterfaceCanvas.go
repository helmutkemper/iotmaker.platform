package iotmaker_platform

import iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"

type ICanvas interface {
	BeginPath()
	MoveTo(x, y iotmaker_platform_coordinate.Coordinate)
	ArcTo(x, y, radius, startAngle, endAngle iotmaker_platform_coordinate.Coordinate)
	LineTo(x, y iotmaker_platform_coordinate.Coordinate)
	ClosePath(x, y iotmaker_platform_coordinate.Coordinate)
}
