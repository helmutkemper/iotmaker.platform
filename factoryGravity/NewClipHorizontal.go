package factoryGravity

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/gravity"

// en: Flag to clip the edges of the object to its container along the horizontal
// axis.
func NewClipHorizontal() gravity.Gravity {
	return gravity.KClipHorizontal
}
