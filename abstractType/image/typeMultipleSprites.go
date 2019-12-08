package image

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	"time"
)

type MultipleSprites struct {
	Platform                iotmaker_platform_IDraw.IDraw
	Img                     interface{}
	SpriteWidth             int
	SpriteHeight            int
	SpriteFirstElementIndex int
	SpriteLastElementIndex  int
	SpriteChangeInterval    time.Duration
	X                       int
	Y                       int
	Width                   int
	Height                  int
	LifeCycleLimit          int
}

func (el *MultipleSprites) Crete() {
	el.Platform.DrawImageMultiplesSprites(el.Img, el.SpriteWidth, el.SpriteHeight, el.SpriteFirstElementIndex, el.SpriteLastElementIndex, el.SpriteChangeInterval, el.X, el.Y, el.Width, el.Height, el.LifeCycleLimit)
}
