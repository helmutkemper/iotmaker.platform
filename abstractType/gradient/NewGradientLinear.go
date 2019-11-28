package gradient

func NewGradientLinear(coordinate Coordinate, colorList []ColorStop) Gradient {
	return Gradient{
		Type:       0,
		Coordinate: coordinate,
		ColorList:  colorList,
	}
}
