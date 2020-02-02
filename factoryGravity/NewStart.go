package factoryGravity

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/gravity"

// en: Push object to x-axis position at the start of its container, not changing
// its size.
func NewStart() gravity.Gravity {
	return gravity.KStart
}
