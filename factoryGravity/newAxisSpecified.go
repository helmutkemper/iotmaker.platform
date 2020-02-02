package factoryGravity

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/gravity"

// en: Raw bit indicating the gravity for an axis has been specified.
func NewAxisSpecified() gravity.Gravity {
	return gravity.KAxisSpecified
}
