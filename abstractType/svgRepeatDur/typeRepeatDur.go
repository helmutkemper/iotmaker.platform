package svgRepeatDur

import "time"

type RepeatDur time.Duration

// en: This value specifies the duration in presentation time to repeat the
// animation.
func (el *RepeatDur) SetClockValue(value RepeatDur) {
	*el = value
}

// en: This value indicates that the animation will be repeated indefinitely (i.e.
// until the document ends).
func (el *RepeatDur) SetIndefinite() {
	*el = RepeatDur(-1)
}
