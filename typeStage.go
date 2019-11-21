package iotmaker_platform

import (
	pwb "github.com/helmutkemper/iotmaker.platform.webbrowser"
)

type GraphicInput struct {
	X                          float64
	Y                          float64
	Border                     float64
	Width                      float64
	Height                     float64
	LineThickness              float64
	UsedInTheCalcX             float64
	UsedInTheCalcY             float64
	UsedInTheCalcBorder        float64
	UsedInTheCalcWidth         float64
	UsedInTheCalcHeight        float64
	UsedInTheCalcLineThickness float64
}

type Graphic struct {
	Canvas canvas
	Data   pwb.PlatformBasicType
	Input  GraphicInput
}

func (el *Graphic) SetX(x float64) {
	el.Input.X = x
}

func (el *Graphic) SetY(y float64) {
	el.Input.Y = y
}

func (el *Graphic) SetBorder(border float64) {
	el.Input.Y = border
}

func (el *Graphic) SetWidth(width float64) {
	el.Input.Width = width
}

func (el *Graphic) SetHeight(height float64) {
	el.Input.Height = height
}

func (el *Graphic) SetWidthHeight(width, height float64) {
	el.Input.Width = width
	el.Input.Height = height
}

func (el *Graphic) SetXYWidthHeight(x, y, width, height float64) {
	el.Input.X = x
	el.Input.Y = y
	el.Input.Width = width
	el.Input.Height = height
}

func (el *Graphic) SetXYWidthHeightLineThickness(x, y, width, height, lineThickness float64) {
	el.Input.X = x
	el.Input.Y = y
	el.Input.Width = width
	el.Input.Height = height
	el.Input.LineThickness = lineThickness
}

func (el *Graphic) SetLineThickness(lineThickness float64) {
	el.Input.LineThickness = lineThickness
}

func (el *Graphic) ManageCoordinatesValuesUsedInTheCalc() {
	el.setNewValuesUsedInTheCalc()
	el.managerCoordinateAndSizeToAdjustLineThicknessExtraSpacesIntoLine()
}

func (el *Graphic) setNewValuesUsedInTheCalc() {
	el.Input.UsedInTheCalcX = el.Input.X
	el.Input.UsedInTheCalcY = el.Input.Y
	el.Input.UsedInTheCalcBorder = el.Input.Border
	el.Input.UsedInTheCalcWidth = el.Input.Width
	el.Input.UsedInTheCalcHeight = el.Input.Height
	el.Input.UsedInTheCalcLineThickness = el.Input.LineThickness
}

func (el *Graphic) managerCoordinateAndSizeToAdjustLineThicknessExtraSpacesIntoLine() {
	adjust := el.Input.LineThickness / 2
	el.Input.UsedInTheCalcX = el.Input.X - adjust
	el.Input.UsedInTheCalcY = el.Input.Y - adjust
	el.Input.UsedInTheCalcWidth = el.Input.Width + adjust
	el.Input.UsedInTheCalcHeight = el.Input.Height + adjust
}

func (el *Graphic) LineTo(x, y float64) {
	el.Canvas.LineTo(x, y)
}

func (el *Graphic) MoveTo(x, y float64) {
	el.Canvas.MoveTo(x, y)
}

func (el *Graphic) ArcTo(x1, y1, x2, y2, radius float64) {
	el.Canvas.ArcTo(x1, y1, x2, y2, radius)
}

func (el *Graphic) ClosePath(x, y float64) {
	el.Canvas.ClosePath(x, y)
}

func (el *Graphic) LineThickness(value float64) {
	el.Canvas.LineThickness(value)
}

func (el *Graphic) AppendToDocumentBody() {
	el.Canvas.AppendToDocumentBody()
}

func (el *Graphic) BeginPath() {
	el.Canvas.BeginPath()
}

func (el *Graphic) Stroke() {
	el.Canvas.Stroke()
}

func (el *Graphic) GetCanvas() pwb.PlatformBasicType {
	return el.Canvas.GetCanvas()
}

func (el *Graphic) GetContext() pwb.PlatformBasicType {
	return el.Canvas.GetContext()
}

func (el *Graphic) GetImageData(x, y, width, height float64) pwb.PlatformBasicType {
	return el.Canvas.GetImageData(x, y, width, height)
}
