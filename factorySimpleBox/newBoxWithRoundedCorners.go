package factorySimpleBox

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.interfaces/iStage"
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/simple"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/basic"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/dimensions"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/ink"
)

// en: This function create a path of a box with rounded corners into the platform using density
//
// pt_br: Esta função cria o caminho de um quadrado com cantos arredondados na plataforma usando densidade
func NewBoxWithRoundedCorners(

	id string,
	stage iStage.IStage,
	platform,
	scratchPad iotmaker_platform_IDraw.IDraw,
	ink ink.Interface,
	x,
	y,
	width,
	height,
	border int,
	density interface{},
	iDensity iotmaker_platform_coordinate.IDensity,

) *simple.BoxWithRoundedCorners {

	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.SetInt(x)
	x = densityCalc.Int()

	densityCalc.SetInt(y)
	y = densityCalc.Int()

	densityCalc.SetInt(width)
	width = densityCalc.Int()

	densityCalc.SetInt(height)
	height = densityCalc.Int()

	densityCalc.SetInt(border)
	border = densityCalc.Int()

	rect := &simple.BoxWithRoundedCorners{
		Sprite: basic.Sprite{
			Id:         id,
			Stage:      stage,
			Platform:   platform,
			ScratchPad: scratchPad,
			Ink:        ink,
			Dimensions: dimensions.Dimensions{
				X:      x,
				Y:      y,
				Width:  width,
				Height: height,
				Border: border,
			},
			OutBoxDimensions: dimensions.Dimensions{
				X:      x,
				Y:      y,
				Width:  width,
				Height: height,
			},
		},
	}
	rect.Crete()

	return rect
}
