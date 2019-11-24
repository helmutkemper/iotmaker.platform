package iotmaker_platform

import iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"

// en: Moves the path to the specified point in the canvas, without creating a line
//     x: The x-coordinate of where to move the path to
//     y: The y-coordinate of where to move the path to
//     The moveTo() method moves the path to the specified point in the canvas, without creating a line.
//     Tip: Use the stroke() method to actually draw the path on the canvas.
func (el *Draw) MoveTo(x, y iotmaker_platform_coordinate.Coordinate) {
	el.Canvas.MoveTo(x.Int(), y.Int())
}
