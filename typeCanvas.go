package iotmaker_platform

import (
	pwb "github.com/helmutkemper/iotmaker.platform.webbrowser"
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
)

type canvas struct {
	Canvas pwb.Canvas
}

func (el *canvas) LineTo(x, y iotmaker_types.Pixel) {
	el.Canvas.LineTo(x, y)
}

func (el *canvas) MoveTo(x, y iotmaker_types.Pixel) {
	el.Canvas.MoveTo(x, y)
}

func (el *canvas) ArcTo(x1, y1, x2, y2, radius iotmaker_types.Pixel) {
	el.Canvas.ArcTo(x1, y1, x2, y2, radius)
}

func (el *canvas) ClosePath(x, y iotmaker_types.Pixel) {
	el.Canvas.ClosePath(x, y)
}

func (el *canvas) LineThickness(value iotmaker_types.Pixel) {
	el.Canvas.LineWidth(value)
}

func (el *canvas) AppendToDocumentBody() {
	el.Canvas.AppendToDocumentBody()
}

func (el *canvas) BeginPath() {
	el.Canvas.BeginPath()
}

func (el *canvas) Stroke() {
	el.Canvas.Stroke()
}

func (el *canvas) GetCanvas() pwb.PlatformBasicType {
	return el.Canvas.GetCanvas()
}

func (el *canvas) GetContext() pwb.PlatformBasicType {
	return el.Canvas.GetContext()
}

func (el *canvas) GetImageData(x, y, width, height iotmaker_types.Pixel) pwb.PlatformBasicType {
	return el.Canvas.GetImageData(x, y, width, height)
}
