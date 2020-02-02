package factoryGravity

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/gravity"

// en: Grow the horizontal size of the object if needed so it completely fills its
// container.
func NewFillHorizontal() gravity.Gravity {
	return gravity.KFillHorizontal
}
