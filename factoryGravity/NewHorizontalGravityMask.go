package factoryGravity

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/gravity"

// en: Binary mask to get the absolute horizontal gravity of a gravity.
func NewHorizontalGravityMask() gravity.Gravity {
	return gravity.KHorizontalGravityMask
}
