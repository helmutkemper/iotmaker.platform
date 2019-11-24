package iotmaker_platform

type ICanvas interface {
	BeginPath()
	MoveTo(x, y Coordinate)
	ArcTo(x, y, radius, startAngle, endAngle Coordinate)
	LineTo(x, y Coordinate)
	ClosePath(x, y Coordinate)
}
