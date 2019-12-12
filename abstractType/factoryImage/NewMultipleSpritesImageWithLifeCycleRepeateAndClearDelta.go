package factoryImage

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform/abstractType/image"
	"time"
)

func NewMultipleSpritesImageWithLifeCycleRepeatAndClearDelta(platform iotmaker_platform_IDraw.IDraw, img interface{}, spriteWidth, spriteHeight, spriteFirstElementIndex, spriteLastElementIndex int, spriteChangeInterval time.Duration, xImageOut, yImageOut, widthImageOut, heightImageOut, clearRectDeltaX, clearRectDeltaY, clearRectDeltaWidth, clearRectDeltaHeight, lifeCycleLimit int, lifeCycleRepeatInterval time.Duration, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) *image.MultipleSprites {
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

	densityCalc.Set(clearRectDeltaX)
	clearRectDeltaX = densityCalc.Int()

	densityCalc.Set(clearRectDeltaY)
	clearRectDeltaY = densityCalc.Int()

	densityCalc.Set(clearRectDeltaWidth)
	clearRectDeltaWidth = densityCalc.Int()

	densityCalc.Set(clearRectDeltaHeight)
	clearRectDeltaHeight = densityCalc.Int()

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
		ClearRectDeltaX:         clearRectDeltaX,
		ClearRectDeltaY:         clearRectDeltaY,
		ClearRectDeltaWidth:     clearRectDeltaWidth,
		ClearRectDeltaHeight:    clearRectDeltaHeight,
		LifeCycleLimit:          lifeCycleLimit,
		LifeCycleRepeatInterval: lifeCycleRepeatInterval,
	}
	ret.Crete()

	return ret
}
