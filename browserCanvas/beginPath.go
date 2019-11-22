package browserCanvas

//	en: Begins a path, or resets the current path
//      Tip: Use moveTo(), lineTo(), quadricCurveTo(), bezierCurveTo(), arcTo(), and arc(), to create paths.
//      Tip: Use the stroke() method to actually draw the path on the canvas.
func (el *Draw) BeginPath() {
	el.Canvas.Browser.BeginPath()
}
