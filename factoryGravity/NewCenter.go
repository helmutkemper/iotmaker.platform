package factoryGravity

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/gravity"

// en: Place the object in the center of its container in both the vertical and
// horizontal axis, not changing its size.
func NewCenter() gravity.Gravity {
	return gravity.KCenter
}
