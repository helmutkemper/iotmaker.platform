package iotmaker_platform

import (
	pwb "github.com/helmutkemper/iotmaker.platform.webbrowser"
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"image/color"
)

type DrawInput struct {
	X                          iotmaker_types.Pixel
	Y                          iotmaker_types.Pixel
	Border                     iotmaker_types.Pixel
	Width                      iotmaker_types.Pixel
	Height                     iotmaker_types.Pixel
	LineThickness              iotmaker_types.Pixel
	UsedInTheCalcX             iotmaker_types.Pixel
	UsedInTheCalcY             iotmaker_types.Pixel
	UsedInTheCalcBorder        iotmaker_types.Pixel
	UsedInTheCalcWidth         iotmaker_types.Pixel
	UsedInTheCalcHeight        iotmaker_types.Pixel
	UsedInTheCalcLineThickness iotmaker_types.Pixel
}

type Draw struct {
	Canvas canvas
	Data   pwb.PlatformBasicType
	Input  DrawInput
}

func (el *Draw) SetX(x iotmaker_types.Pixel) {
	el.Input.X = x
}

func (el *Draw) SetY(y iotmaker_types.Pixel) {
	el.Input.Y = y
}

func (el *Draw) SetBorder(border iotmaker_types.Pixel) {
	el.Input.Y = border
}

func (el *Draw) SetWidth(width iotmaker_types.Pixel) {
	el.Input.Width = width
}

func (el *Draw) SetHeight(height iotmaker_types.Pixel) {
	el.Input.Height = height
}

func (el *Draw) SetWidthHeight(width, height iotmaker_types.Pixel) {
	el.Input.Width = width
	el.Input.Height = height
}

func (el *Draw) SetXYWidthHeight(x, y, width, height iotmaker_types.Pixel) {
	el.Input.X = x
	el.Input.Y = y
	el.Input.Width = width
	el.Input.Height = height
}

func (el *Draw) SetXYWidthHeightLineThickness(x, y, width, height, lineThickness iotmaker_types.Pixel) {
	el.Input.X = x
	el.Input.Y = y
	el.Input.Width = width
	el.Input.Height = height
	el.Input.LineThickness = lineThickness
}

func (el *Draw) SetLineThickness(lineThickness iotmaker_types.Pixel) {
	el.Input.LineThickness = lineThickness
}

func (el *Draw) ManageCoordinatesValuesUsedInTheCalc() {
	el.setNewValuesUsedInTheCalc()
	el.managerCoordinateAndSizeToAdjustLineThicknessExtraSpacesIntoLine()
}

func (el *Draw) setNewValuesUsedInTheCalc() {
	el.Input.UsedInTheCalcX = el.Input.X
	el.Input.UsedInTheCalcY = el.Input.Y
	el.Input.UsedInTheCalcBorder = el.Input.Border
	el.Input.UsedInTheCalcWidth = el.Input.Width
	el.Input.UsedInTheCalcHeight = el.Input.Height
	el.Input.UsedInTheCalcLineThickness = el.Input.LineThickness
}

func (el *Draw) managerCoordinateAndSizeToAdjustLineThicknessExtraSpacesIntoLine() {
	adjust := el.Input.LineThickness / 2
	el.Input.UsedInTheCalcX = el.Input.X - adjust
	el.Input.UsedInTheCalcY = el.Input.Y - adjust
	el.Input.UsedInTheCalcWidth = el.Input.Width + adjust
	el.Input.UsedInTheCalcHeight = el.Input.Height + adjust
}

func (el *Draw) LineTo(x, y iotmaker_types.Pixel) {
	el.Canvas.LineTo(x, y)
}

func (el *Draw) MoveTo(x, y iotmaker_types.Pixel) {
	el.Canvas.MoveTo(x, y)
}

func (el *Draw) ArcTo(x1, y1, x2, y2, radius iotmaker_types.Pixel) {
	el.Canvas.ArcTo(x1, y1, x2, y2, radius)
}

func (el *Draw) ClosePath(x, y iotmaker_types.Pixel) {
	el.Canvas.ClosePath(x, y)
}

func (el *Draw) LineThickness(value iotmaker_types.Pixel) {
	el.Canvas.LineThickness(value)
}

func (el *Draw) AppendToDocumentBody() {
	el.Canvas.AppendToDocumentBody()
}

func (el *Draw) BeginPath() {
	el.Canvas.BeginPath()
}

func (el *Draw) Stroke() {
	el.Canvas.Stroke()
}

func (el *Draw) GetCanvas() pwb.PlatformBasicType {
	return el.Canvas.GetCanvas()
}

func (el *Draw) GetContext() pwb.PlatformBasicType {
	return el.Canvas.GetContext()
}

func (el *Draw) GetImageData(x, y, width, height iotmaker_types.Pixel) [][]color.RGBA {
	return el.Canvas.GetImageData(x, y, width, height)
}

func (el *Draw) AddBasicBox(x, y, width, height, border, lineThickness iotmaker_types.Pixel) {
	//              border        border
	//             x1  x2         x3 x4
	//           l     a          b     c
	//        y1    +--|----------|--+    y1
	//  border      |                |      border
	//        y2 k ---              --- d y2
	//              |                |
	//              |                |
	//        y3 j ---              --- e y3
	//  border      |                |      border
	//        y4    +--|----------|--+    y4
	//           i     h          g     f
	//             x1  x2         x3 x4
	//              border        border

	// tired programmer's rules (kiss like)
	x += lineThickness / 2
	width -= lineThickness / 2
	height -= lineThickness / 2

	x1 := x
	x2 := x1 + border
	x3 := x2 + width - 2*border
	x4 := x3 + border

	y1 := y
	y2 := y1 + border
	y3 := y2 + height - 2*border
	y4 := y3 + border

	el.Canvas.LineThickness(lineThickness)

	el.Canvas.MoveTo(x2, y1)                // a
	el.Canvas.LineTo(x3, y1)                // a->b
	el.Canvas.ArcTo(x4, y1, x4, y2, border) // c->d
	el.Canvas.LineTo(x4, y3)                // d->e
	el.Canvas.ArcTo(x4, y4, x3, y4, border) // f->g
	el.Canvas.LineTo(x2, y4)                // g->h
	el.Canvas.ArcTo(x1, y4, x1, y3, border) // i->j
	el.Canvas.LineTo(x1, y2)                // j->k
	el.Canvas.ArcTo(x1, y1, x2, y1, border) // i->j
	el.Canvas.ClosePath(x2, y1)             // a
}
