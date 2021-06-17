package svg

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/svgDur"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/svgRepeatCount"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/svgRepeatDur"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/svgRestart"
	"time"
)

type Animate struct {
	// en: The begin attribute defines when an animation should begin or when an
	// element should be discarded.
	// The attribute value is a semicolon separated list of values. The interpretation
	// of a list of start times is detailed in the SMIL specification in "Evaluation of
	// begin and end time lists". Each individual value can be one of the following :
	// <offset-value>, <syncbase-value>, <event-value>, <repeat-value>,
	// <accessKey-value>, <wallclock-sync-value> or the keyword indefinite.
	Begin svgDur.Dur

	// en: The dur attribute indicates the simple duration of an animation.
	Dur svgDur.Dur

	// en: The end attribute defines an end value for the animation that can constrain
	// the active duration.
	End svgDur.Dur

	// en: The min attribute specifies the minimum value of the active animation
	// duration.
	Min time.Duration

	// en: The max attribute specifies the maximum value of the active animation
	// duration.
	Max time.Duration

	// en: The restart attribute specifies whether or not an animation can restart.
	Restart svgRestart.Restart

	// en: The repeatCount attribute indicates the number of times an animation will
	// take place.
	RepeatCount svgRepeatCount.RepeatCount

	// en: The repeatDur attribute specifies the total duration for repeating an
	// animation.
	RepeatDur svgRepeatDur.RepeatDur

	// en: The fill attribute has two different meanings. For shapes and text it's a
	// presentation attribute that defines the color (or any SVG paint servers like
	// gradients or patterns) used to paint the element; for animation it defines the
	// final state of the animation.
	Fill string
}
