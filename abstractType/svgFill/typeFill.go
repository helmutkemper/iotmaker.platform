package svgFill

// https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/fill
type Fill int

// en: For <animate>, fill defines the final state of the animation.
// Value	freeze (Keep the state of the last animation frame) | remove (Keep the
// state of the first animation frame)
// Default value	remove
// Animatable	No
func (el *Fill) Animate() {}

// en: For <animateMotion>, fill defines the final state of the animation.
// Value	freeze (Keep the state of the last animation frame) | remove (Keep the
// state of the first animation frame)
//Default value	remove
//Animatable	No
func (el *Fill) AnimateMotion() {}

// en: For <animateTransform>, fill defines the final state of the animation.
// Value	freeze (Keep the state of the last animation frame) | remove (Keep the
// state of the first animation frame)
// Default value	remove
// Animatable	No
func (el *Fill) AnimateTransform() {}

// en: For <circle>, fill is a presentation attribute that defines the color of the
// circle.
// Value	<paint>
// Default value	black
// Animatable	Yes
// Note: As a presentation attribute fill can be used as a CSS property.
func (el *Fill) Circle() {}

// en: For <ellipse>, fill is a presentation attribute that defines the color of the
// ellipse.
// Value	<paint>
// Default value	black
// Animatable	Yes
// Note: As a presentation attribute fill can be used as a CSS property.
func (el *Fill) Ellipse() {}

// en: For <path>, fill is a presentation attribute that defines the color of the
// interior of the shape. (Interior is define by the fill-rule attribute)
// Value	<paint>
// Default value	black
// Animatable	Yes
// Note: As a presentation attribute fill can be used as a CSS property.
func (el *Fill) Path() {}

// en: For <polygon>, fill is a presentation attribute that defines the color of the
// interior of the shape. (Interior is define by the fill-rule attribute)
// Value	<paint>
// Default value	black
// Animatable	Yes
// Note: As a presentation attribute fill can be used as a CSS property.
func (el *Fill) Polygon() {}

// en: For <polyline>, fill is a presentation attribute that defines the color of the
// interior of the shape. (Interior is define by the fill-rule attribute)
// Value	<paint>
// Default value	black
// Animatable	Yes
// Note: As a presentation attribute fill can be used as a CSS property.
func (el *Fill) Polyline() {}

// en: For <rect>, fill is a presentation attribute that defines the color of the
// rectangle.
// Value	<paint>
// Default value	black
// Animatable	Yes
// Note: As a presentation attribute fill can be used as a CSS property.
func (el *Fill) Rect() {}

// en: For <set>, fill defines the final state of the animation.
// Value	freeze (Keep the state of the last animation frame) | remove (Keep the
// state of the first animation frame)
// Default value	remove
// Animatable	No
func (el *Fill) Set() {}

// en: For <text>, fill is a presentation attribute that defines what the color of
// the text.
// Value	<paint>
// Default value	black
// Animatable	Yes
// Note: As a presentation attribute fill can be used as a CSS property.
func (el *Fill) Text() {}

// en: For <textPath>, fill is a presentation attribute that defines the color of the
// text.
// Value	<paint>
// Default value	black
// Animatable	Yes
// Note: As a presentation attribute fill can be used as a CSS property.
func (el *Fill) TextPath() {}

// en: For <tspan>, fill is a presentation attribute that defines the color of the
// text.
// Value	<paint>
// Default value	black
// Animatable	Yes
// Note: As a presentation attribute fill can be used as a CSS property.
func (el *Fill) TSpan() {}
