package factoryText

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.interfaces/iStage"
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/font"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/text"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/basic"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/dimensions"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/ink"
)

func NewText(

	id string,
	stage iStage.IStage,
	platform,
	scratchPad iotmaker_platform_IDraw.IDraw,
	ink ink.Interface,
	font font.Font,
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

	densityCalc.SetInt(font.Size)
	font.Size = densityCalc.Int()

	tx := text.Text{
		Sprite: basic.Sprite{
			Id:               id,
			Stage:            stage,
			Platform:         platform,
			ScratchPad:       scratchPad,
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
