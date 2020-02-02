package factoryGravity

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/gravity"

// en: Place object in the horizontal center of its container, not changing its
// size.
func NewCenterHorizontal() gravity.Gravity {
	return gravity.KCenterHorizontal
}
