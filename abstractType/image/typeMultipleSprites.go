package image

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	"time"
)

type MultipleSprites struct {
	Platform                iotmaker_platform_IDraw.IDraw
	Img                     interface{}
	SpriteWidth             float64
	SpriteHeight            float64
	SpriteFirstElementIndex float64
	SpriteLastElementIndex  float64
	SpriteChangeInterval    time.Duration
	X                       float64
	Y                       float64
	Width                   float64
	Height                  float64

	ClearRectDeltaX      float64
	ClearRectDeltaY      float64
	ClearRectDeltaWidth  float64
	ClearRectDeltaHeight float64

	LifeCycleLimit          float64
	LifeCycleRepeatLimit    float64
	LifeCycleRepeatInterval time.Duration
}

func (el *MultipleSprites) Crete() {
	el.Platform.DrawImageMultiplesSprites(el.Img, el.SpriteWidth, el.SpriteHeight, el.SpriteFirstElementIndex, el.SpriteLastElementIndex, el.SpriteChangeInterval,
		el.X, el.Y, el.Width, el.Height,
		el.ClearRectDeltaX, el.ClearRectDeltaY, el.ClearRectDeltaWidth, el.ClearRectDeltaHeight,
		el.LifeCycleLimit, el.LifeCycleRepeatLimit, el.LifeCycleRepeatInterval)
}
