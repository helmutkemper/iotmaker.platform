package gradient

type Select int

const (
	KLinearFill Select = iota + 1
	KLinearStroke
	KLinearFillAndStroke
	KRadialFill
	KRadialStroke
	KRadialFillAndStroke
)
