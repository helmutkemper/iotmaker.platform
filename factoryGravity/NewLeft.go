package factoryGravity

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/gravity"

// en: Push object to the left of its container, not changing its size.
func NewLeft() gravity.Gravity {
	return gravity.KLeft
}
