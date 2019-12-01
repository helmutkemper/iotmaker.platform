package gradient

func NewFactoryGradient(typeGradient Select) {
	switch typeGradient {
	case KLinearGradientFill:
	case KLinearGradientStroke:
	case KRadialGradientFill:
	case KRadialGradientStroke:
	}
}
