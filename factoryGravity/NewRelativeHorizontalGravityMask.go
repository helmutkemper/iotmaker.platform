package factoryGravity

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/gravity"

// en: Binary mask for the horizontal gravity and script specific direction bit.
func NewRelativeHorizontalGravityMask() gravity.Gravity {
	return gravity.KRelativeHorizontalGravityMask
}
