package text

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/font"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/basic"
	"reflect"
)

type Text struct {
	basic.Sprite

	//prepareShadowFilterFunctionPointer   func(iotmaker_platform_IDraw.ICanvasShadow)
	//prepareGradientFilterFunctionPointer func(iotmaker_platform_IDraw.ICanvasGradient)

	Label    string
	Font     font.Font
	Fill     bool
	Stroke   bool
	MaxWidth int
	X        int
	Y        int
}

/*func (el *Text) prepareShadowFilter() {
	if el.prepareShadowFilterFunctionPointer != nil {
		el.prepareShadowFilterFunctionPointer(el.Platform)
	}
}*/

/*func (el *Text) prepareGradientFilter(platform iotmaker_platform_IDraw.IDraw) {
	if reflect.DeepEqual(el.Ink.Color, color.RGBA{}) != true {
		platform.SetFillStyle(el.Ink.Color)
	}

	if el.prepareGradientFilterFunctionPointer != nil {
		el.prepareGradientFilterFunctionPointer(platform)
	}
}*/

/*func (el *Text) ConfigShadowPlatformAndFilter() {
	if el.Ink.Shadow == nil {
		return
	}

	el.setShadowFilter(el.Ink.Shadow.PrepareFilter)
}*/

/*func (el *Text) ConfigGradientPlatformAndFilter() {
	if el.Ink.Gradient == nil {
		return
	}

	el.setGradientAndMountColorListFilter(el.Ink.Gradient.PrepareFilter)
}*/

/*func (el *Text) setGradientAndMountColorListFilter(f func(iotmaker_platform_IDraw.ICanvasGradient)) {
	el.prepareGradientFilterFunctionPointer = f
}*/

/*func (el *Text) setShadowFilter(f func(iotmaker_platform_IDraw.ICanvasShadow)) {
	el.prepareShadowFilterFunctionPointer = f
}*/

func (el *Text) Create() {
	el.Draw()
}

func (el *Text) Clear() {
	el.Dimensions.X = 10
	el.Dimensions.Y = 10
	el.Dimensions.Width = 300
	el.Dimensions.Height = 100
	el.ScratchPad.ClearRect(el.Dimensions.X, el.Dimensions.Y, el.Dimensions.Width, el.Dimensions.Height)
}

func (el *Text) Draw() {
	el.ScratchPad.Save()

	if reflect.DeepEqual(font.Font{}, el.Font) == false {
		el.ScratchPad.Font(el.Font)
	}

	if el.Fill == true {
		if el.MaxWidth != 0 {
			el.ScratchPad.FillText(el.Label, el.X, el.Y, el.MaxWidth)
		} else {
			el.ScratchPad.FillText(el.Label, el.X, el.Y)
		}
	}

	if el.Stroke {
		if el.MaxWidth != 0 {
			el.ScratchPad.StrokeText(el.Label, el.X, el.Y, el.MaxWidth)
		} else {
			el.ScratchPad.StrokeText(el.Label, el.X, el.Y)
		}
	}

	el.ScratchPad.Restore()
}
