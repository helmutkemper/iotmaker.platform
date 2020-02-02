package factoryGravity

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/gravity"

// en: Raw bit controlling whether the layout direction is relative or not
// (START/END instead of absolute LEFT/RIGHT).
func NewRelativeLayoutDirection() gravity.Gravity {
	return gravity.KRelativeLayoutDirection
}
