package factoryText

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.interfaces/iStage"
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/font"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/text"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/basic"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryInk"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/ink"
)

func NewTextWithFont(

	id string,
	stage iStage.IStage,
	platform,
	scratchPad iotmaker_platform_IDraw.IDraw,
	shadow iotmaker_platform_IDraw.IFilterShadowInterface,
	gradient iotmaker_platform_IDraw.IFilterGradientInterface,
	color interface{},
	labelFont font.Font,
	label string,
	x,
	y int,
	density interface{},
	iDensity iotmaker_platform_coordinate.IDensity,

) *text.Text {

	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.SetInt(x)
	x = densityCalc.Int()

	densityCalc.SetInt(y)
	y = densityCalc.Int()

	ik := ink.Ink{}
	ik = factoryInk.NewInk(ik, 0, color, shadow, gradient, density, iDensity)

	tx := text.Text{
		Sprite: basic.Sprite{
			Id:         id,
			Stage:      stage,
			Platform:   platform,
			ScratchPad: scratchPad,
			Ink:        ik,
		},
		Label: label,
		Font:  labelFont,
		Fill:  true,
		X:     x,
		Y:     y,
	}

	tx.Create()

	return &tx
}
