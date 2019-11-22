package iotmaker_platform

import (
	"github.com/helmutkemper/iotmaker.platform/browserCanvas"
	"github.com/helmutkemper/iotmaker.platform/canvas"
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"image/color"
)

type DrawInput struct {
	X                          iotmaker_types.Coordinate
	Y                          iotmaker_types.Coordinate
	Border                     iotmaker_types.Coordinate
	Width                      iotmaker_types.Coordinate
	Height                     iotmaker_types.Coordinate
	LineThickness              iotmaker_types.Coordinate
	UsedInTheCalcX             iotmaker_types.Coordinate
	UsedInTheCalcY             iotmaker_types.Coordinate
	UsedInTheCalcBorder        iotmaker_types.Coordinate
	UsedInTheCalcWidth         iotmaker_types.Coordinate
	UsedInTheCalcHeight        iotmaker_types.Coordinate
	UsedInTheCalcLineThickness iotmaker_types.Coordinate
}

type Draw struct {
	Canvas browserCanvas.Canvas
	Data   iotmaker_types.PlatformBasicType
	Input  DrawInput
}

func (el *Draw) SetX(x iotmaker_types.Coordinate) {
	el.Input.X = x
}

func (el *Draw) SetY(y iotmaker_types.Coordinate) {
	el.Input.Y = y
}

func (el *Draw) SetBorder(border iotmaker_types.Coordinate) {
	el.Input.Y = border
}

func (el *Draw) SetWidth(width iotmaker_types.Coordinate) {
	el.Input.Width = width
}

func (el *Draw) SetHeight(height iotmaker_types.Coordinate) {
	el.Input.Height = height
}

func (el *Draw) SetWidthHeight(width, height iotmaker_types.Coordinate) {
	el.Input.Width = width
	el.Input.Height = height
}

func (el *Draw) SetXYWidthHeight(x, y, width, height iotmaker_types.Coordinate) {
	el.Input.X = x
	el.Input.Y = y
	el.Input.Width = width
	el.Input.Height = height
}

func (el *Draw) SetXYWidthHeightLineThickness(x, y, width, height, lineThickness iotmaker_types.Coordinate) {
	el.Input.X = x
	el.Input.Y = y
	el.Input.Width = width
	el.Input.Height = height
	el.Input.LineThickness = lineThickness
}

func (el *Draw) SetLineThickness(lineThickness iotmaker_types.Coordinate) {
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

func (el *Draw) LineTo(x, y iotmaker_types.Coordinate) {
	el.Canvas.LineTo(x, y)
}

func (el *Draw) MoveTo(x, y iotmaker_types.Coordinate) {
	el.Canvas.MoveTo(x, y)
}

func (el *Draw) ArcTo(x1, y1, x2, y2, radius iotmaker_types.Coordinate) {
	el.Canvas.ArcTo(x1, y1, x2, y2, radius)
}

func (el *Draw) ClosePath(x, y iotmaker_types.Coordinate) {
	el.Canvas.ClosePath(x, y)
}

func (el *Draw) LineThickness(value iotmaker_types.Coordinate) {
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

func (el *Draw) GetCanvas() iotmaker_types.PlatformBasicType {
	return el.Canvas.GetCanvas()
}

func (el *Draw) GetContext() iotmaker_types.PlatformBasicType {
	return el.Canvas.GetContext()
}

func (el *Draw) GetImageData(x, y, width, height iotmaker_types.Coordinate) [][]color.RGBA {
	return el.Canvas.GetImageData(x, y, width, height)
}

func (el *Draw) AddBasicBox(x, y, width, height, border, lineThickness iotmaker_types.Coordinate) {
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

	var i ICanvas = &Draw{}

	el.Canvas.LineThickness(lineThickness)

	i.MoveTo(x2, y1)                // a
	i.LineTo(x3, y1)                // a->b
	i.ArcTo(x4, y1, x4, y2, border) // c->d
	i.LineTo(x4, y3)                // d->e
	i.ArcTo(x4, y4, x3, y4, border) // f->g
	i.LineTo(x2, y4)                // g->h
	i.ArcTo(x1, y4, x1, y3, border) // i->j
	i.LineTo(x1, y2)                // j->k
	i.ArcTo(x1, y1, x2, y1, border) // i->j
	i.ClosePath(x2, y1)             // a
}
