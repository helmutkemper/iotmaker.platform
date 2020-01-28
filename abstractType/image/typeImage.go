package image

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/basic"
)

type Image struct {
	basic.Sprite
	Img  interface{}
	data interface{}
}

func (el *Image) Crete() {
	el.Draw()
}

func (el *Image) Draw() {
	el.ScratchPad.Save()
	el.ScratchPad.DrawImage(el.Img, el.Dimensions.X, el.Dimensions.Y, el.Dimensions.Width, el.Dimensions.Height)
	el.ScratchPad.Restore()
}

func (el *Image) Clear() {
	el.ScratchPad.ClearRect(el.Dimensions.X, el.Dimensions.Y, el.Dimensions.Width, el.Dimensions.Height)
}
