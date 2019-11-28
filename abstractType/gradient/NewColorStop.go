package gradient

import "image/color"

func NewColorPosition(color color.RGBA, percentPosition float64) ColorStop {
	return ColorStop{
		Color: color,
		Stop:  percentPosition,
	}
}
