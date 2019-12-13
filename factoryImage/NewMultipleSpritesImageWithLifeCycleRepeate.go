package factoryImage

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform/abstractType/image"
	"time"
)

func NewMultipleSpritesImageWithLifeCycleRepeat(platform iotmaker_platform_IDraw.IDraw, img interface{}, spriteWidth, spriteHeight, spriteFirstElementIndex, spriteLastElementIndex int, spriteChangeInterval time.Duration, xImageOut, yImageOut, widthImageOut, heightImageOut, lifeCycleLimit int, lifeCycleRepeatInterval time.Duration, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) *image.MultipleSprites {
	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.Set(xImageOut)
	xImageOut = densityCalc.Int()

	densityCalc.Set(yImageOut)
	yImageOut = densityCalc.Int()

	densityCalc.Set(widthImageOut)
	widthImageOut = densityCalc.Int()

	densityCalc.Set(heightImageOut)
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
		LifeCycleLimit:          lifeCycleLimit,
		LifeCycleRepeatInterval: lifeCycleRepeatInterval,
	}
	ret.Crete()

	return ret
}
