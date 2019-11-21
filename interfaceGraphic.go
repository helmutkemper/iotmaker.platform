package iotmaker_platform

import (
	pwb "github.com/helmutkemper/iotmaker.platform.webbrowser"
)

type InterfaceGraphic interface {
	SetX(x float64)
	SetY(y float64)
	SetBorder(border float64)
	SetWidth(width float64)
	SetHeight(height float64)
	SetWidthHeight(width, height float64)
	SetXYWidthHeight(x, y, width, height float64)
	SetXYWidthHeightLineThickness(x, y, width, height, lineThickness float64)
	SetLineThickness(lineThickness float64)
	ManageCoordinatesValuesUsedInTheCalc()
	LineTo(x, y float64)
	MoveTo(x, y float64)
	ArcTo(x1, y1, x2, y2, radius float64)
	ClosePath(x, y float64)
	LineThickness(value float64)
	AppendToDocumentBody()
	BeginPath()
	Stroke()
	GetCanvas() pwb.PlatformBasicType
	GetContext() pwb.PlatformBasicType
	GetImageData(x, y, width, height float64) pwb.PlatformBasicType
}
