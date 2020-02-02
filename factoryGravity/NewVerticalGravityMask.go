package factoryGravity

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/gravity"

// en: Binary mask to get the vertical gravity of a gravity.
func NewVerticalGravityMask() gravity.Gravity {
	return gravity.KVerticalGravityMask
}
