package gradient

type Select int

const (
	KLinearGradientFill Select = iota + 1
	KLinearGradientStroke
	KRadialGradientFill
	KRadialGradientStroke
)
