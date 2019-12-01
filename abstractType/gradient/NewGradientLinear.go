package gradient

func NewFillGradientLinear(coordinateP0, coordinateP1 Point, colorList []ColorStop) Gradient {
	return Gradient{
		Type:         KLinearGradientFill,
		CoordinateP0: coordinateP0,
		CoordinateP1: coordinateP1,
		ColorList:    colorList,
	}
}
