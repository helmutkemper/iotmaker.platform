package text

import (
	iotmakerPlatformTextMetrics "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.textMetrics"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/font"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/basic"
	"image/color"
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
}

func (el *Text) Create() {
	el.Draw()
}

func (el *Text) Clear() {
	el.ScratchPad.ClearRect(el.Dimensions.X, el.Dimensions.Y, el.Dimensions.Width, el.Dimensions.Height)
}

func (el *Text) Draw() {
	el.ScratchPad.Save()

	metrics := iotmakerPlatformTextMetrics.TextMetrics{}

	if reflect.DeepEqual(font.Font{}, el.Font) == false {
		el.ScratchPad.Font(el.Font)
		metrics = el.ScratchPad.MeasureText(el.Label)

		if reflect.DeepEqual(color.RGBA{}, el.Font.Color) == false && el.Fill == true {
			el.ScratchPad.SetFillStyle(el.Font.Color)
		}

		if reflect.DeepEqual(color.RGBA{}, el.Font.Color) == false && el.Stroke == true {
			el.ScratchPad.SetStrokeStyle(el.Font.Color)
		}
	}

	el.Dimensions.Width = int(metrics.Width)
	el.Dimensions.Height = el.Font.Size

	el.OutBoxDimensions.Width = int(metrics.Width)
	el.OutBoxDimensions.Height = el.Font.Size

	if el.Fill == true {
		if el.MaxWidth != 0 {
			el.ScratchPad.FillText(el.Label, el.Dimensions.X, el.Dimensions.Y+el.OutBoxDimensions.Height, el.MaxWidth)
		} else {
			el.ScratchPad.FillText(el.Label, el.Dimensions.X, el.Dimensions.Y+el.OutBoxDimensions.Height)
		}
	}

	if el.Stroke {
		if el.MaxWidth != 0 {
			el.ScratchPad.StrokeText(el.Label, el.Dimensions.X, el.Dimensions.Y+el.OutBoxDimensions.Height, el.MaxWidth)
		} else {
			el.ScratchPad.StrokeText(el.Label, el.Dimensions.X, el.Dimensions.Y+el.OutBoxDimensions.Height)
		}
	}

	el.ScratchPad.Restore()
}
