package browserCanvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
)

// en: Adds a new point and creates a line from that point to the last specified point in the canvas (this method does
//     not draw the line).
//
//     x:   The x-coordinate of where to create the line to
//     y:   The y-coordinate of where to create the line to
//     Tip: Use the stroke() method to actually draw the path on the canvas.
func (el *Draw) LineTo(x, y iotmaker_types.Coordinate) {
	el.Canvas.Browser.LineTo(x, y)
}
