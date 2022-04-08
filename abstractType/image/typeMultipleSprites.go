package image

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/basic"
	"time"
)

type MultipleSprites struct {
	basic.Sprite

	Platform                iotmaker_platform_IDraw.IDraw
	Img                     interface{}
	SpriteFirstElementIndex int
	SpriteLastElementIndex  int
	SpriteChangeInterval    time.Duration

	ClearRectDeltaX      int
	ClearRectDeltaY      int
	ClearRectDeltaWidth  int
	ClearRectDeltaHeight int

	LifeCycleLimit          int
	LifeCycleRepeatLimit    int
	LifeCycleRepeatInterval time.Duration
}

func (el *MultipleSprites) Crete() {
	el.Platform.DrawImageMultiplesSprites(
		el.Img,
		el.Dimensions.Width, el.Dimensions.Height,
		el.SpriteFirstElementIndex, el.SpriteLastElementIndex,
		el.SpriteChangeInterval,
		el.Dimensions.X, el.Dimensions.Y, el.Dimensions.Width, el.Dimensions.Height,
		el.ClearRectDeltaX, el.ClearRectDeltaY, el.ClearRectDeltaWidth, el.ClearRectDeltaHeight,
		el.LifeCycleLimit, el.LifeCycleRepeatLimit, el.LifeCycleRepeatInterval,
	)
}

func (el *MultipleSprites) Draw() {
	el.ScratchPad.Save()
	el.ColorFiltersStart(el.ScratchPad)
	el.ScratchPad.DrawImage(
		el.Img,
		el.Dimensions.X,
		el.Dimensions.Y,
		el.Dimensions.Width,
		el.Dimensions.Height,
	)
	el.ScratchPad.Restore()
}

func (el *MultipleSprites) Clear() {
	el.ScratchPad.ClearRect(
		el.Dimensions.X,
		el.Dimensions.Y,
		el.Dimensions.Width,
		el.Dimensions.Height,
	)
}
