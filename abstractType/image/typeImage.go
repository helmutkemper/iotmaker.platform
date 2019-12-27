package image

import (
	"github.com/helmutkemper/iotmaker.platform.webbrowser/basic"
)

type Image struct {
	basic.Sprite
	//Platform   iotmaker_platform_IDraw.IDraw
	//ScratchPad iotmaker_platform_IDraw.IDraw
	Img interface{}
	//X          float64
	//Y          float64
	//Width      float64
	//Height     float64
	data interface{}
}

func (el *Image) Crete() {
	el.Draw()
}

func (el *Image) Draw() {
	el.Platform.DrawImage(el.Img, el.OutBoxDimensions.X, el.OutBoxDimensions.Y, el.OutBoxDimensions.Width, el.OutBoxDimensions.Height)
}

func (el *Image) Clear() {
	el.Platform.ClearRect(el.OutBoxDimensions.X, el.OutBoxDimensions.Y, el.OutBoxDimensions.Width, el.OutBoxDimensions.Height)
}
