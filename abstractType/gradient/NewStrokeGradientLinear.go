package gradient

func NewStrokeGradientLinear(coordinate Coordinate, colorList []ColorStop) Gradient {
	return Gradient{
		Type:       KLinearGradientStroke,
		Coordinate: coordinate,
		ColorList:  colorList,
	}
}
