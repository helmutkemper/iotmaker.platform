package iotmaker_platform

import iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"

// en: Creates a path from the current point back to the starting point
//     The closePath() method creates a path from the current point back to the starting point.
//     Tip: Use the stroke() method to actually draw the path on the canvas.
//     Tip: Use the fill() method to fill the drawing (black is default). Use the fillStyle property to fill with
//     another color/gradient.
func (el *Draw) ClosePath(x, y iotmaker_platform_coordinate.Density) {
	el.Canvas.ClosePath(x.Int(), y.Int())
}
