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

func NewTextToButton(

	id string,
	stage iStage.IStage,
	platform,
	scratchPad iotmaker_platform_IDraw.IDraw,
	ink ink.Ink,
	font font.Font,
	label string,
	density interface{},
	iDensity iotmaker_platform_coordinate.IDensity,

) *text.Text {

	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.SetInt(font.Size)
	font.Size = densityCalc.Int()

	tx := text.Text{
		Sprite: basic.Sprite{
			Id:         id,
			Stage:      stage,
			Platform:   platform,
			ScratchPad: scratchPad,
			Ink:        ink,
		},
		Label: label,
		Font:  font,
		Fill:  true,
	}

	tx.Create()

	return &tx
}
