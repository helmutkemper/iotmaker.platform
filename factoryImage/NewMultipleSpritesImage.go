package factoryImage

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/image"
	"time"
)

func NewMultipleSpritesImage(
	platform iotmaker_platform_IDraw.IDraw,
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
		Platform:                platform,
		Img:                     img,
		SpriteWidth:             spriteWidth,
		SpriteHeight:            spriteHeight,
		SpriteFirstElementIndex: spriteFirstElementIndex,
		SpriteLastElementIndex:  spriteLastElementIndex,
		SpriteChangeInterval:    spriteChangeInterval,
		X:                       xImageOut,
		Y:                       yImageOut,
		Width:                   widthImageOut,
		Height:                  heightImageOut,
	}
	ret.Crete()

	return ret
}
