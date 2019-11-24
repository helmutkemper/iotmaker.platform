package iotmaker_platform

import iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"

// en: Adds a new point and creates a line from that point to the last specified point in the canvas (this method does
//     not draw the line).
//
//     x:   The x-coordinate of where to create the line to
//     y:   The y-coordinate of where to create the line to
//     Tip: Use the stroke() method to actually draw the path on the canvas.
func (el *Draw) LineTo(x, y iotmaker_platform_coordinate.Coordinate) {
	el.Canvas.LineTo(x.Int(), y.Int())
}
