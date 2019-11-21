package iotmaker_platform

import (
	pwb "github.com/helmutkemper/iotmaker.platform.webbrowser"
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
)

type GraphicInput struct {
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

type Graphic struct {
	Canvas canvas
	Data   pwb.PlatformBasicType
	Input  GraphicInput
}

func (el *Graphic) SetX(x iotmaker_types.Pixel) {
	el.Input.X = x
}

func (el *Graphic) SetY(y iotmaker_types.Pixel) {
	el.Input.Y = y
}

func (el *Graphic) SetBorder(border iotmaker_types.Pixel) {
	el.Input.Y = border
}

func (el *Graphic) SetWidth(width iotmaker_types.Pixel) {
	el.Input.Width = width
}

func (el *Graphic) SetHeight(height iotmaker_types.Pixel) {
	el.Input.Height = height
}

func (el *Graphic) SetWidthHeight(width, height iotmaker_types.Pixel) {
	el.Input.Width = width
	el.Input.Height = height
}

func (el *Graphic) SetXYWidthHeight(x, y, width, height iotmaker_types.Pixel) {
	el.Input.X = x
	el.Input.Y = y
	el.Input.Width = width
	el.Input.Height = height
}

func (el *Graphic) SetXYWidthHeightLineThickness(x, y, width, height, lineThickness iotmaker_types.Pixel) {
	el.Input.X = x
	el.Input.Y = y
	el.Input.Width = width
	el.Input.Height = height
	el.Input.LineThickness = lineThickness
}

func (el *Graphic) SetLineThickness(lineThickness iotmaker_types.Pixel) {
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

func (el *Graphic) LineTo(x, y iotmaker_types.Pixel) {
	el.Canvas.LineTo(x, y)
}

func (el *Graphic) MoveTo(x, y iotmaker_types.Pixel) {
	el.Canvas.MoveTo(x, y)
}

func (el *Graphic) ArcTo(x1, y1, x2, y2, radius iotmaker_types.Pixel) {
	el.Canvas.ArcTo(x1, y1, x2, y2, radius)
}

func (el *Graphic) ClosePath(x, y iotmaker_types.Pixel) {
	el.Canvas.ClosePath(x, y)
}

func (el *Graphic) LineThickness(value iotmaker_types.Pixel) {
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

func (el *Graphic) GetImageData(x, y, width, height iotmaker_types.Pixel) pwb.PlatformBasicType {
	return el.Canvas.GetImageData(x, y, width, height)
}
