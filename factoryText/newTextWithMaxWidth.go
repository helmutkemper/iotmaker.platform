package factoryText

import (
	global "github.com/helmutkemper/iotmaker.santa_isabel_theater.globalConfig"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/font"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/text"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/basic"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/ink"
)

func NewTextWithMaxWidth(

	id string,
	ink ink.Interface,
	labelFont font.Font,
	label string,
	x,
	y int,
	maxWidth int,

) *text.Text {

	densityCalc := global.Global.DensityManager
	densityCalc.SetDensityFactor(global.Global.Density)

	densityCalc.SetInt(x)
	x = densityCalc.Int()

	densityCalc.SetInt(y)
	y = densityCalc.Int()

	tx := text.Text{
		Sprite: basic.Sprite{
			Id:         id,
			Stage:      global.Global.Stage,
			Platform:   global.Global.Canvas,
			ScratchPad: global.Global.ScratchPad,
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
