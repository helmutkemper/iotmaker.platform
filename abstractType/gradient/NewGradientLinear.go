package gradient

func NewFillGradientLinear(coordinate Coordinate, colorList []ColorStop) Gradient {
	return Gradient{
		Type:       KLinearGradientFill,
		Coordinate: coordinate,
		ColorList:  colorList,
	}
}
