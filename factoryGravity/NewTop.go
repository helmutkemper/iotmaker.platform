package factoryGravity

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/gravity"

// en: Push object to the top of its container, not changing its size.
func NewTop() gravity.Gravity {
	return gravity.KTop
}
