package iotmaker_platform

import (
	"github.com/helmutkemper/iotmaker.platform.webbrowser"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/canvas"
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"image/color"
	"syscall/js"
)

type IDraw interface {
	GetImageData(x, y, width, height iotmaker_types.Pixel) [][]color.RGBA
	FillStyle(value string)
	StrokeStyle(value string)
	ShadowColor(value string)
	ShadowBlur(value string)
	ShadowOffsetX(value string)
	ShadowOffsetY(value string)
	CreateLinearGradient(x0, y0, x1, y1 iotmaker_types.Pixel)
	CreatePattern(image js.Value, repeatRule iotmaker_types.CanvasRepeatRule)
	CreateRadialGradient(x0, y0, r0, x1, y1, r1 iotmaker_types.Pixel)
	AddColorStop(stop iotmaker_types.Pixel, color js.Value)
	GlobalAlpha(value iotmaker_types.Pixel)
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
	LineWidth(value iotmaker_types.Pixel)
	MiterLimit(value iotmaker_types.Pixel)
	Fill()
	Stroke()
	BeginPath()
	MoveTo(x, y iotmaker_types.Pixel)
	Clip(x, y iotmaker_types.Pixel)
	QuadraticCurveTo(cpx, cpy, x, y iotmaker_types.Pixel)
	BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y iotmaker_types.Pixel)
	Arc(x, y, radius, startAngle, endAngle iotmaker_types.Pixel, anticlockwise bool)
	IsPointInPath(path js.Value, x, y iotmaker_types.Pixel, fillRule iotmaker_types.CanvasFillRule) bool
	Get(jsParam string, value ...interface{}) iotmaker_types.Pixel
	Width() iotmaker_types.Pixel
	Height() iotmaker_types.Pixel
	Data() js.Value
	CreateImageData(data js.Value)
	PutImageData(imgData js.Value, x, y, dirtyX, dirtyY, dirtyWidth, dirtyHeight iotmaker_types.Pixel)
	Rect(x, y, width, height iotmaker_types.Pixel)
	FillRect(x, y, width, height iotmaker_types.Pixel)
	StrokeRect(x, y, width, height iotmaker_types.Pixel)
	ClearRect(x, y, width, height iotmaker_types.Pixel)
	TextAlign(value iotmaker_types.CanvasFontAlignRule)
	TextBaseline(value iotmaker_types.CanvasTextBaseLineRule)
	FillText(text string, x, y, maxWidth iotmaker_types.Pixel)
	StrokeText(text string, x, y, maxWidth iotmaker_types.Pixel)
	MeasureText(text string)
	Scale(scaleWidth, scaleHeight iotmaker_types.Pixel)
	Rotate(angle iotmaker_types.Pixel)
	Translate(x, y iotmaker_types.Pixel)
	Transform(a, b, c, d, e, f iotmaker_types.Pixel)
	SetTransform(a, b, c, d, e, f iotmaker_types.Pixel)
	ClosePath(x, y iotmaker_types.Pixel)
	LineTo(x, y iotmaker_types.Pixel)
	AppendToDocumentBody()
	LineJoin(value iotmaker_types.CanvasJoinRule)
	ArcTo(x1, y1, x2, y2, radius iotmaker_types.Pixel)
	Font(font iotmaker_types.Font)
}
