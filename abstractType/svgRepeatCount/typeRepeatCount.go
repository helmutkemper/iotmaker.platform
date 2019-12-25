package svgRepeatCount

type RepeatCount int

// en: This value specifies the number of iterations. It can include partial
// iterations expressed as fraction values. A fractional value describes a portion
// of the simple duration. Values must be greater than 0.
func (el *RepeatCount) SetNumber(number RepeatCount) {
	*el = number
}

// en: This value indicates that the animation will be repeated indefinitely (i.e.
// until the document ends).
func (el *RepeatCount) SetIndefinite() {
	*el = -1
}
