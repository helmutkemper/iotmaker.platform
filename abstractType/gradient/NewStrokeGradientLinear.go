package gradient

import iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"

func NewStrokeGradientLinear(coordinate Coordinate, colorList []ColorStop) iotmaker_platform_IDraw.IFilterGradientInterface {
	return &Gradient{
		Type:       KLinearGradientStroke,
		Coordinate: coordinate,
		ColorList:  colorList,
	}
}
