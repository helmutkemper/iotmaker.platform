package iotmaker_platform

import (
  iotmaker_types "github.com/helmutkemper/iotmaker.types"
  "image/color"
  "syscall/js"
)

type IDraw interface {
	GetImageData(x, y, width, height iotmaker_types.Coordinate) [][]color.RGBA
	FillStyle(value string)
	StrokeStyle(value string)
	ShadowColor(value string)
	ShadowBlur(value string)
	ShadowOffsetX(value string)
	ShadowOffsetY(value string)
	CreateLinearGradient(x0, y0, x1, y1 iotmaker_types.Coordinate)
	CreatePattern(image js.Value, repeatRule iotmaker_types.CanvasRepeatRule)
	CreateRadialGradient(x0, y0, r0, x1, y1, r1 iotmaker_types.Coordinate)
	AddColorStop(stop iotmaker_types.Coordinate, color js.Value)
	GlobalAlpha(value iotmaker_types.Coordinate)
	GlobalCompositeOperation(value iotmaker_types.CanvasCompositeOperationsRule)
	GetCanvas() js.Value
	GetContext() js.Value
	InitializeContext2DById(id string)
	InitializeContext3DById(id string)
	Save()
	Restore()
	CreateEvent()
	DrawImage(value iotmaker_types.DrawImage)
	LineCap(value iotmaker_types.CanvasCapRule)
	LineWidth(value iotmaker_types.Coordinate)
	MiterLimit(value iotmaker_types.Coordinate)
	Fill()
	Stroke()
	BeginPath()
	MoveTo(x, y iotmaker_types.Coordinate)
	Clip(x, y iotmaker_types.Coordinate)
	QuadraticCurveTo(cpx, cpy, x, y iotmaker_types.Coordinate)
	BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y iotmaker_types.Coordinate)
	Arc(x, y, radius, startAngle, endAngle iotmaker_types.Coordinate, anticlockwise bool)
	IsPointInPath(path js.Value, x, y iotmaker_types.Coordinate, fillRule iotmaker_types.CanvasFillRule) bool
	Get(jsParam string, value ...interface{}) iotmaker_types.Coordinate
	Width() iotmaker_types.Coordinate
	Height() iotmaker_types.Coordinate
	Data() js.Value
	CreateImageData(data js.Value)
	PutImageData(imgData js.Value, x, y, dirtyX, dirtyY, dirtyWidth, dirtyHeight iotmaker_types.Coordinate)
	Rect(x, y, width, height iotmaker_types.Coordinate)
	FillRect(x, y, width, height iotmaker_types.Coordinate)
	StrokeRect(x, y, width, height iotmaker_types.Coordinate)
	ClearRect(x, y, width, height iotmaker_types.Coordinate)
	TextAlign(value iotmaker_types.CanvasFontAlignRule)
	TextBaseline(value iotmaker_types.CanvasTextBaseLineRule)
	FillText(text string, x, y, maxWidth iotmaker_types.Coordinate)
	StrokeText(text string, x, y, maxWidth iotmaker_types.Coordinate)
	MeasureText(text string)
	Scale(scaleWidth, scaleHeight iotmaker_types.Coordinate)
	Rotate(angle iotmaker_types.Coordinate)
	Translate(x, y iotmaker_types.Coordinate)
	Transform(a, b, c, d, e, f iotmaker_types.Coordinate)
	SetTransform(a, b, c, d, e, f iotmaker_types.Coordinate)
	ClosePath(x, y iotmaker_types.Coordinate)
	LineTo(x, y iotmaker_types.Coordinate)
	AppendToDocumentBody()
	LineJoin(value iotmaker_types.CanvasJoinRule)
	ArcTo(x1, y1, x2, y2, radius iotmaker_types.Coordinate)
	Font(font iotmaker_types.Font)
}
