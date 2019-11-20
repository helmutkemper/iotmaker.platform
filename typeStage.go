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
	XUsedInTheCalc             float64
	YUsedInTheCalc             float64
	BorderUsedInTheCalc        float64
	WidthUsedInTheCalc         float64
	HeightUsedInTheCalc        float64
	LineThicknessUsedInTheCalc float64
}

type Graphic struct {
	Canvas pwb.Canvas
	Data   pwb.Value
	Input  GraphicInput
}

func (el *Graphic) SetX(x float64) {
	el.Input.X = x
	el.setNewValuesUsedInTheCalc()
}

func (el *Graphic) SetY(y float64) {
	el.Input.Y = y
	el.setNewValuesUsedInTheCalc()
}

func (el *Graphic) SetBorder(border float64) {
	el.Input.Y = border
	el.setNewValuesUsedInTheCalc()
}

func (el *Graphic) SetWidth(width float64) {
	el.Input.Width = width
	el.setNewValuesUsedInTheCalc()
}

func (el *Graphic) SetHeight(height float64) {
	el.Input.Height = height
	el.setNewValuesUsedInTheCalc()
}

func (el *Graphic) SetWidthHeight(width, height float64) {
	el.Input.Width = width
	el.Input.Height = height
	el.setNewValuesUsedInTheCalc()
}

func (el *Graphic) SetXYWidthHeight(x, y, width, height float64) {
	el.Input.X = x
	el.Input.Y = y
	el.Input.Width = width
	el.Input.Height = height
	el.setNewValuesUsedInTheCalc()
}

func (el *Graphic) SetXYWidthHeightLineThickness(x, y, width, height, lineThickness float64) {
	el.Input.X = x
	el.Input.Y = y
	el.Input.Width = width
	el.Input.Height = height
	el.Input.LineThickness = lineThickness
	el.setNewValuesUsedInTheCalc()
}

func (el *Graphic) SetLineThickness(lineThickness float64) {
	el.Input.LineThickness = lineThickness
	el.setNewValuesUsedInTheCalc()
}

func (el *Graphic) setNewValuesUsedInTheCalc() {
	el.Input.XUsedInTheCalc = el.Input.X
	el.Input.YUsedInTheCalc = el.Input.Y
	el.Input.BorderUsedInTheCalc = el.Input.Border
	el.Input.WidthUsedInTheCalc = el.Input.Width
	el.Input.HeightUsedInTheCalc = el.Input.Height
	el.Input.LineThicknessUsedInTheCalc = el.Input.LineThickness
}

func (el *Graphic) ManagerCoordinateAndSizeToAdjustLineThicknessExtraSpacesIntoLine() {
	adjust := el.Input.LineThickness / 2
	el.Input.XUsedInTheCalc = el.Input.X - adjust
	el.Input.YUsedInTheCalc = el.Input.Y - adjust
	el.Input.WidthUsedInTheCalc = el.Input.Width + adjust
	el.Input.HeightUsedInTheCalc = el.Input.Height + adjust
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
	el.Canvas.LineWidth(value)
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
