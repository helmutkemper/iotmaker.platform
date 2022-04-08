package factoryImage

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.interfaces/iStage"
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/image"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/basic"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/dimensions"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/ink"
	"time"
)

func NewMultipleSpritesImage(

	id string,
	stage iStage.IStage,
	platform,
	scratchPad iotmaker_platform_IDraw.IDraw,
	ink ink.Interface,
	img interface{},
	spriteWidth,
	spriteHeight,
	spriteFirstElementIndex,
	spriteLastElementIndex int,
	spriteChangeInterval time.Duration,
	xImageOut,
	yImageOut,
	widthImageOut,
	heightImageOut int,
	density interface{},
	iDensity iotmaker_platform_coordinate.IDensity,
) *image.MultipleSprites {

	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.SetInt(xImageOut)
	xImageOut = densityCalc.Int()

	densityCalc.SetInt(yImageOut)
	yImageOut = densityCalc.Int()

	densityCalc.SetInt(widthImageOut)
	widthImageOut = densityCalc.Int()

	densityCalc.SetInt(heightImageOut)
	heightImageOut = densityCalc.Int()

	ret := &image.MultipleSprites{
		Sprite: basic.Sprite{
			Id:         id,
			Stage:      stage,
			Platform:   platform,
			ScratchPad: scratchPad,
			Ink:        ink,
			Dimensions: dimensions.Dimensions{
				X:      xImageOut,
				Y:      yImageOut,
				Width:  widthImageOut,
				Height: heightImageOut,
			},
			OutBoxDimensions: dimensions.Dimensions{
				X:      xImageOut,
				Y:      yImageOut,
				Width:  widthImageOut,
				Height: heightImageOut,
			},
		},
		Platform:                platform,
		Img:                     img,
		SpriteFirstElementIndex: spriteFirstElementIndex,
		SpriteLastElementIndex:  spriteLastElementIndex,
		SpriteChangeInterval:    spriteChangeInterval,
	}
	ret.Crete()

	return ret
}
