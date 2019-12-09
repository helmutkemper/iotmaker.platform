package factoryColor

import (
	"github.com/helmutkemper/iotmaker.platform/abstractType/gradient"
	"image/color"
)

func NewColorPosition(color color.RGBA, percentPosition float64) gradient.ColorStop {
	return gradient.ColorStop{
		Color: color,
		Stop:  percentPosition,
	}
}
