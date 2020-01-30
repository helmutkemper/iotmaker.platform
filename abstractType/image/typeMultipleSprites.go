package image

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
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

	ClearRectDeltaX      int
	ClearRectDeltaY      int
	ClearRectDeltaWidth  int
	ClearRectDeltaHeight int

	LifeCycleLimit          int
	LifeCycleRepeatLimit    int
	LifeCycleRepeatInterval time.Duration
}

func (el *MultipleSprites) Crete() {
	el.Platform.DrawImageMultiplesSprites(el.Img, el.SpriteWidth, el.SpriteHeight, el.SpriteFirstElementIndex, el.SpriteLastElementIndex, el.SpriteChangeInterval,
		el.X, el.Y, el.Width, el.Height,
		el.ClearRectDeltaX, el.ClearRectDeltaY, el.ClearRectDeltaWidth, el.ClearRectDeltaHeight,
		el.LifeCycleLimit, el.LifeCycleRepeatLimit, el.LifeCycleRepeatInterval)
}
