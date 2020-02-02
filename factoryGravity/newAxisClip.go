package factoryGravity

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/gravity"

// en: Raw bit controlling whether the right/bottom edge is clipped to its
// container, based on the gravity direction being applied.
func NewAxisClip() gravity.Gravity {
	return gravity.KAxisClip
}
