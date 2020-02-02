package factoryGravity

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/gravity"

// en: Push object to the right of its container, not changing its size.
func NewRight() gravity.Gravity {
	return gravity.KRight
}
