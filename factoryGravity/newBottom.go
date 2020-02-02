package factoryGravity

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/gravity"

// en: Push object to the bottom of its container, not changing its size.
func NewBottom() gravity.Gravity {
	return gravity.KBottom
}
