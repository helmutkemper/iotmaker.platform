package iotmaker_platform

import (
	"github.com/helmutkemper/iotmaker.platform.webbrowser"
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"image/color"
)

type IDraw interface {
	SetX(x iotmaker_types.Pixel)
	SetY(y iotmaker_types.Pixel)
	SetBorder(border iotmaker_types.Pixel)
	SetWidth(width iotmaker_types.Pixel)
	SetHeight(height iotmaker_types.Pixel)
	SetWidthHeight(width, height iotmaker_types.Pixel)
	SetXYWidthHeight(x, y, width, height iotmaker_types.Pixel)
	SetXYWidthHeightLineThickness(x, y, width, height, lineThickness iotmaker_types.Pixel)
	SetLineThickness(lineThickness iotmaker_types.Pixel)
	ManageCoordinatesValuesUsedInTheCalc()
	LineTo(x, y iotmaker_types.Pixel)
	MoveTo(x, y iotmaker_types.Pixel)
	ArcTo(x1, y1, x2, y2, radius iotmaker_types.Pixel)
	ClosePath(x, y iotmaker_types.Pixel)
	LineThickness(value iotmaker_types.Pixel)
	AppendToDocumentBody()
	BeginPath()
	Stroke()
	GetCanvas() iotmaker_platform_webbrowser.PlatformBasicType
	GetContext() iotmaker_platform_webbrowser.PlatformBasicType
	GetImageData(x, y, width, height iotmaker_types.Pixel) [][]color.RGBA
}
