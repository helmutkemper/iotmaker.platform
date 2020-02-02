package factoryGravity

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/gravity"

// en: Grow the vertical size of the object if needed so it completely fills its
// container.
func NewFillVertical() gravity.Gravity {
	return gravity.KFillVertical
}
