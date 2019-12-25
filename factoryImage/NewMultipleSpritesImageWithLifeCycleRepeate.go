package factoryImage

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform/abstractType/image"
	"time"
)

func NewMultipleSpritesImageWithLifeCycleRepeat(platform iotmaker_platform_IDraw.IDraw, img interface{}, spriteWidth, spriteHeight, spriteFirstElementIndex, spriteLastElementIndex float64, spriteChangeInterval time.Duration, xImageOut, yImageOut, widthImageOut, heightImageOut, lifeCycleLimit float64, lifeCycleRepeatInterval time.Duration, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) *image.MultipleSprites {
	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.Set(xImageOut)
	xImageOut = densityCalc.Float64()

	densityCalc.Set(yImageOut)
	yImageOut = densityCalc.Float64()

	densityCalc.Set(widthImageOut)
	widthImageOut = densityCalc.Float64()

	densityCalc.Set(heightImageOut)
	heightImageOut = densityCalc.Float64()

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
