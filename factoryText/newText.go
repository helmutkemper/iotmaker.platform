package factoryText

import (
	global "github.com/helmutkemper/iotmaker.santa_isabel_theater.globalConfig"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/font"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/text"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/basic"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/dimensions"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/ink"
)

func NewText(

	id string,
	ink ink.Interface,
	font font.Font,
	label string,
	x,
	y int,
) *text.Text {

	densityCalc := global.Global.DensityManager
	densityCalc.SetDensityFactor(global.Global.Density)

	densityCalc.SetInt(x)
	x = densityCalc.Int()

	densityCalc.SetInt(y)
	y = densityCalc.Int()

	densityCalc.SetInt(font.Size)
	font.Size = densityCalc.Int()

	tx := text.Text{
		Sprite: basic.Sprite{
			Id:               id,
			Stage:            global.Global.Stage,
			Platform:         global.Global.Canvas,
			ScratchPad:       global.Global.ScratchPad,
			Ink:              ink,
			Dimensions:       dimensions.Dimensions{X: x, Y: y},
			OutBoxDimensions: dimensions.Dimensions{X: x, Y: y},
		},
		Label: label,
		Font:  font,
		Fill:  true,
	}

	tx.Create()

	return &tx
}
