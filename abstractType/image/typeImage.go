package image

import (
	"github.com/helmutkemper/iotmaker.platform.webbrowser/basic"
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
	el.Platform.DrawImage(el.Img, el.Dimensions.X, el.Dimensions.Y, el.Dimensions.Width, el.Dimensions.Height)
}

func (el *Image) Clear() {
	el.Platform.ClearRect(el.Dimensions.X, el.Dimensions.Y, el.Dimensions.Width, el.Dimensions.Height)
}
