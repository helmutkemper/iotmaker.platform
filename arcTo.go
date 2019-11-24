package iotmaker_platform

// en: Creates an arc/curve (used to create circles, or parts of circles)
//     x:             The horizontal coordinate of the arc's center.
//     y:             The vertical coordinate of the arc's center.
//     radius:        The arc's radius. Must be positive.
//     startAngle:    The angle at which the arc starts in radians, measured from the positive x-axis.
//     endAngle:      The angle at which the arc ends in radians, measured from the positive x-axis.
func (el *Draw) ArcTo(x, y, radius, startAngle, endAngle Coordinate) {
	el.Canvas.ArcTo(x.Int(), y.Int(), radius.Int(), startAngle.Int(), endAngle.Int())
}
