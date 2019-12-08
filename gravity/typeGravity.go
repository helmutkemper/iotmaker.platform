package gravity

import "strings"

type Gravity int

func (el Gravity) String() string {
	return gravitiesString[el]
}

func (el Gravity) ToType(value string) Gravity {
	return gravitiesStringToGravityMap[strings.ToLower(value)]
}

var gravitiesString = [...]string{
	"No Gravity",
	"Axis Clip",
	"Axis Pull After",
	"Axis Pull Before",
	"Axis Specified",
	"Axis X Shift",
	"Axis Y Shift",
	"Bottom",
	"Center",
	"Center Horizontal",
	"Center Vertical",
	"Clip Horizontal",
	"Clip Vertical",
	"Display Clip Horizontal",
	"Display Clip Vertical",
	"End",
	"Fill",
	"Fill Horizontal",
	"Fill Vertical",
	"Horizontal Gravity Mask",
	"Left",
	"Relative Horizontal Gravity Mask",
	"Relative Layout Direction",
	"Right",
	"Start",
	"Top",
	"Vertical Gravity Mask",
}

var gravitiesStringToGravityMap = map[string]Gravity{
	"no gravity":                       KNoGravity,
	"axis clip":                        KAxisClip,
	"axis pull after":                  KAxisPullAfter,
	"axis pull before":                 KAxisPullBefore,
	"axis specified":                   KAxisSpecified,
	"axis x shift":                     KAxisXShift,
	"axis y shift":                     KAxisYShift,
	"bottom":                           KBottom,
	"center":                           KCenter,
	"center horizontal":                KCenterHorizontal,
	"center vertical":                  KCenterVertical,
	"clip horizontal":                  KClipHorizontal,
	"clip vertical":                    KClipVertical,
	"display clip horizontal":          KDisplayClipHorizontal,
	"display clip vertical":            KDisplayClipVertical,
	"end":                              KEnd,
	"fill":                             KFill,
	"fill horizontal":                  KFillHorizontal,
	"fill vertical":                    KFillVertical,
	"horizontal gravity mask":          KHorizontalGravityMask,
	"left":                             KLeft,
	"relative horizontal gravity mask": KRelativeHorizontalGravityMask,
	"relative layout direction":        KRelativeLayoutDirection,
	"right":                            KRight,
	"start":                            KStart,
	"top":                              KTop,
	"vertical gravity mask":            KVerticalGravityMask,
}

const (
	// en: Constant indicating that no gravity has been set *
	KNoGravity Gravity = iota

	// en: Raw bit controlling whether the right/bottom edge is clipped to its
	// container, based on the gravity direction being applied.
	KAxisClip

	// en: Raw bit controlling how the right/bottom edge is placed.
	KAxisPullAfter

	// en: Raw bit controlling how the left/top edge is placed.
	KAxisPullBefore

	// en: Raw bit indicating the gravity for an axis has been specified.
	KAxisSpecified

	// en: Bits defining the horizontal axis.
	KAxisXShift

	// en: Bits defining the vertical axis.
	KAxisYShift

	// en: Push object to the bottom of its container, not changing its size.
	KBottom

	// en: Place the object in the center of its container in both the vertical and
	// horizontal axis, not changing its size.
	KCenter

	// en: Place object in the horizontal center of its container, not changing its
	// size.
	KCenterHorizontal

	// en: Place object in the vertical center of its container, not changing its size.
	KCenterVertical

	// en: Flag to clip the edges of the object to its container along the horizontal
	// axis.
	KClipHorizontal

	// en: Flag to clip the edges of the object to its container along the vertical
	// axis.
	KClipVertical

	// en: Special constant to enable clipping to an overall display along the
	// horizontal dimension.
	KDisplayClipHorizontal

	// en: Special constant to enable clipping to an overall display along the vertical
	// dimension.
	KDisplayClipVertical

	// en: Push object to x-axis position at the end of its container, not changing its
	// size.
	KEnd

	// en: Grow the horizontal and vertical size of the object if needed so it
	// completely fills its container.
	KFill

	// en: Grow the horizontal size of the object if needed so it completely fills its
	// container.
	KFillHorizontal

	// en: Grow the vertical size of the object if needed so it completely fills its
	// container.
	KFillVertical

	// en: Binary mask to get the absolute horizontal gravity of a gravity.
	KHorizontalGravityMask

	// en: Push object to the left of its container, not changing its size.
	KLeft

	// en: Binary mask for the horizontal gravity and script specific direction bit.
	KRelativeHorizontalGravityMask

	// en: Raw bit controlling whether the layout direction is relative or not
	// (START/END instead of absolute LEFT/RIGHT).
	KRelativeLayoutDirection

	// en: Push object to the right of its container, not changing its size.
	KRight

	// en: Push object to x-axis position at the start of its container, not changing
	// its size.
	KStart

	// en: Push object to the top of its container, not changing its size.
	KTop

	// en: Binary mask to get the vertical gravity of a gravity.
	KVerticalGravityMask
)
