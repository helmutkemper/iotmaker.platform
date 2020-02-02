package factoryText

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.interfaces/iStage"
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/font"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/text"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/basic"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/ink"
)

func NewTextWithMaxWidth(

	id string,
	stage iStage.IStage,
	platform,
	scratchPad iotmaker_platform_IDraw.IDraw,
	ink ink.Interface,
	labelFont font.Font,
	label string,
	x,
	y int,
	maxWidth int,
	density interface{},
	iDensity iotmaker_platform_coordinate.IDensity,

) *text.Text {

	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.SetInt(x)
	x = densityCalc.Int()

	densityCalc.SetInt(y)
	y = densityCalc.Int()

	tx := text.Text{
		Sprite: basic.Sprite{
			Id:         id,
			Stage:      stage,
			Platform:   platform,
			ScratchPad: scratchPad,
			Ink:        ink,
		},
		Label:    label,
		Font:     labelFont,
		Fill:     true,
		MaxWidth: maxWidth,
	}

	tx.Create()

	return &tx
}
