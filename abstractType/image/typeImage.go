package image

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
)

type Image struct {
	Platform   iotmaker_platform_IDraw.IDraw
	ScratchPad iotmaker_platform_IDraw.IDraw
	Img        interface{}
	X          float64
	Y          float64
	Width      float64
	Height     float64
	data       interface{}
}

func (el *Image) Crete() {
	el.Platform.DrawImage(el.Img, el.X, el.Y, el.Width, el.Height)
}

func (el *Image) SetX(x float64) {
	el.Platform.DrawImage(el.Img, x, el.Y, el.Width, el.Height)
}

func (el *Image) SetY(y float64) {
	el.Platform.DrawImage(el.Img, el.X, y, el.Width, el.Height)
}

func (el *Image) SetXY(x, y float64) {
	el.Platform.DrawImage(el.Img, x, y, el.Width, el.Height)
}
