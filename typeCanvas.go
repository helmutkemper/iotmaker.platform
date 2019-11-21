package iotmaker_platform

import (
	pwb "github.com/helmutkemper/iotmaker.platform.webbrowser"
)

type canvas struct {
	Canvas pwb.Canvas
}

func (el *canvas) LineTo(x, y float64) {
	el.Canvas.LineTo(x, y)
}

func (el *canvas) MoveTo(x, y float64) {
	el.Canvas.MoveTo(x, y)
}

func (el *canvas) ArcTo(x1, y1, x2, y2, radius float64) {
	el.Canvas.ArcTo(x1, y1, x2, y2, radius)
}

func (el *canvas) ClosePath(x, y float64) {
	el.Canvas.ClosePath(x, y)
}

func (el *canvas) LineThickness(value float64) {
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

func (el *canvas) GetImageData(x, y, width, height float64) pwb.PlatformBasicType {
	return el.Canvas.GetImageData(x, y, width, height)
}
