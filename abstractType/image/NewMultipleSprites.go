package image

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"time"
)

func NewMultipleSprites(platform iotmaker_platform_IDraw.IDraw, image interface{}, spriteWidth, spriteHeight, spriteFirstElementIndex, spriteLastElementIndex int, spriteChangeInterval time.Duration, xImageOut, yImageOut, widthImageOut, heightImageOut, lifeCycleLimit int, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) *MultipleSprites {
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

	ret := &MultipleSprites{
		Platform:                platform,
		Img:                     image,
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
	}
	ret.Crete()

	return ret
}
