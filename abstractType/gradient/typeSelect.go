package gradient

type Select int

const (
	KLinearGradientFill Select = iota + 1
	KLinearGradientStroke
	KLinearGradientFillAndStroke
	KRadialGradientFill
	KRadialGradientStroke
	KRadialGradientFillAndStroke
)
