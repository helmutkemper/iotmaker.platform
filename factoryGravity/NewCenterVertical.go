package factoryGravity

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/gravity"

// en: Place object in the vertical center of its container, not changing its size.
func NewCenterVertical() gravity.Gravity {
	return gravity.KCenterVertical
}
