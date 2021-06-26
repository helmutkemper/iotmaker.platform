package factoryText

import (
	global "github.com/helmutkemper/iotmaker.santa_isabel_theater.globalConfig"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/font"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/text"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/basic"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/ink"
)

func NewTextToButton(
	id string,
	ink ink.Interface,
	font font.Font,
	label string,
) *text.Text {

	densityCalc := global.Global.DensityManager
	densityCalc.SetDensityFactor(global.Global.Density)

	densityCalc.SetInt(font.Size)
	font.Size = densityCalc.Int()

	tx := text.Text{
		Sprite: basic.Sprite{
			Id:         id,
			Stage:      global.Global.Stage,
			Platform:   global.Global.Canvas,
			ScratchPad: global.Global.ScratchPad,
			Ink:        ink,
		},
		Label: label,
		Font:  font,
		Fill:  true,
	}

	tx.Create()

	return &tx
}
