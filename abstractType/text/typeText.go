package text

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/font"
	"github.com/helmutkemper/iotmaker.platform/abstractType/genericTypes"
	"image/color"
	"reflect"
)

type Text struct {
	Platform iotmaker_platform_IDraw.IDraw
	Ink      genericTypes.Ink

	prepareShadowFilterFunctionPointer   func(iotmaker_platform_IDraw.ICanvasShadow)
	prepareGradientFilterFunctionPointer func(iotmaker_platform_IDraw.ICanvasGradient)

	Label    string
	Font     font.Font
	Fill     bool
	Stroke   bool
	MaxWidth int
	X        int
	Y        int
}

func (el *Text) prepareShadowFilter() {
	if el.prepareShadowFilterFunctionPointer != nil {
		el.prepareShadowFilterFunctionPointer(el.Platform)
	}
}

func (el *Text) prepareGradientFilter(platform iotmaker_platform_IDraw.IDraw) {
	if reflect.DeepEqual(el.Ink.Color, color.RGBA{}) != true {
		platform.SetFillStyle(el.Ink.Color)
	}

	if el.prepareGradientFilterFunctionPointer != nil {
		el.prepareGradientFilterFunctionPointer(platform)
	}
}

func (el *Text) ConfigShadowPlatformAndFilter() {
	if el.Ink.Shadow == nil {
		return
	}

	el.setShadowFilter(el.Ink.Shadow.PrepareFilter)
}

func (el *Text) ConfigGradientPlatformAndFilter() {
	if el.Ink.Gradient == nil {
		return
	}

	el.setGradientAndMountColorListFilter(el.Ink.Gradient.PrepareFilter)
}

func (el *Text) setGradientAndMountColorListFilter(f func(iotmaker_platform_IDraw.ICanvasGradient)) {
	el.prepareGradientFilterFunctionPointer = f
}

func (el *Text) setShadowFilter(f func(iotmaker_platform_IDraw.ICanvasShadow)) {
	el.prepareShadowFilterFunctionPointer = f
}

func (el *Text) Create() {
	el.prepareGradientFilter(el.Platform)
	el.prepareShadowFilter()

	if reflect.DeepEqual(font.Font{}, el.Font) == false {
		el.Platform.Font(el.Font)
	}

	if el.Fill == true {
		if el.MaxWidth != 0 {
			el.Platform.FillText(el.Label, el.X, el.Y, el.MaxWidth)
		} else {
			el.Platform.FillText(el.Label, el.X, el.Y)
		}
	} else if el.Stroke {
		if el.MaxWidth != 0 {
			el.Platform.StrokeText(el.Label, el.X, el.Y, el.MaxWidth)
		} else {
			el.Platform.StrokeText(el.Label, el.X, el.Y)
		}
	}

	el.Platform.ResetStrokeStyle()
	el.Platform.ResetFillStyle()
	el.Platform.ResetShadow()

}
