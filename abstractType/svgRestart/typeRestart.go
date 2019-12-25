package svgRestart

type Restart int

const (
	// en: This value indicates that the animation can be restarted at any time.
	KAlways Restart = iota

	// en: This value indicates that the animation can only be restarted when it is not
	// active (i.e. after the active end). Attempts to restart the animation during its
	// active duration are ignored.
	KWhenNotActive

	// en: This value indicates that the animation cannot be restarted for the time the
	// document is loaded.
	KNever
)

func (el Restart) String() string {
	return restartString[el]
}

var restartString = [...]string{
	"always",
	"whenNotActive",
	"never",
}
