package svgDur

import "time"

type Dur time.Duration

// en: This value specifies the length of the simple duration. The value must be
// greater than 0 and can be expressed with hours (h), minutes (m), seconds (s) or
// milliseconds (ms).
// Itʼs possible to combine those time representations to create some complex
// durations like hh:mm:ss.iii or mm:ss.iii.
func (el *Dur) SetClockValue(value Dur) {
	*el = value
}

func (el *Dur) SetIndefinite() {
	*el = Dur(-1)
}

// en: This value specifies the simple duration as the intrinsic media duration.
// This is only valid for elements that define media.
// (For animation elements the attribute will be ignored if media is specified.)
func (el *Dur) SetMedia() {
	*el = Dur(-2)
}

// en: This value specifies the simple duration as indefinite.
func (el *Dur) GetDefault() Dur {
	return Dur(-1)
}
