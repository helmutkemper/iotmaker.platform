package iotmaker_platform

type ICanvas interface {
	BeginPath()
	MoveTo(x, y int)
	ArcTo(x, y, radius, startAngle, endAngle int)
	LineTo(x, y int)
	ClosePath(x, y int)
	Stroke()
}
