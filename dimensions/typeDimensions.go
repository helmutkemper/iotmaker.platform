package dimensions

type Overflow int

type Dimensions struct {
	X      int
	Y      int
	Width  int
	Height int
	Border int

	SpaceLeft   int
	SpaceRight  int
	SpaceTop    int
	SpaceBottom int

	Parent   *Dimensions
	Overflow Overflow
}
