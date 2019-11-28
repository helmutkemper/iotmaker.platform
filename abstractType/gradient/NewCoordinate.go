package gradient

// fixme: make density compliance
func NewCoordinate(x, y, width, height int) Coordinate {
	return Coordinate{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}
