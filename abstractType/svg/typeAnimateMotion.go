package svg

// en: The SVG <animateMotion> element let define how an element moves along a motion
// path.
// Note: To reuse an existing path, it will be necessary to use an <mpath> element
// inside the <animateMotion> element instead of the path attribute.
// Note: For <animateMotion> the default value for the calcMode attribute is paced
type AnimateMotion struct {
	// en: This attribute indicate, in the range [0,1], how far is the object along the
	// path for each keyTimes associated values.
	// Value type: <number>*; Default value: none; Animatable: no
	KeyPoints int

	// en: This attribute defines the path of the motion, using the same syntaxe as the
	// d attribute.
	// Value type: <string>; Default value: none; Animatable: no
	Path string

	//en: This attribute defines a rotation applied to the elment animated along a
	// path, usualy to make it pointing in the direction of the animation.
	// Value type: <number>|auto|auto-reverse; Default value: 0; Animatable: no
	Rotate int
}
