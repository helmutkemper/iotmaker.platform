package factoryColor

import "github.com/helmutkemper/iotmaker.platform/abstractType/gradient"

func NewColorList(list ...gradient.ColorStop) []gradient.ColorStop {
	return list
}
