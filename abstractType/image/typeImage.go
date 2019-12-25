package image

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
)

type Image struct {
	Platform   iotmaker_platform_IDraw.IDraw
	ScratchPad iotmaker_platform_IDraw.IDraw
	Img        interface{}
	X          int
	Y          int
	Width      int
	Height     int
	data       interface{}
}

func (el *Image) Crete() {
	el.Platform.DrawImage(el.Img, el.X, el.Y, el.Width, el.Height)
}

func (el *Image) SetX(x int) {
	el.Platform.DrawImage(el.Img, x, el.Y, el.Width, el.Height)
}

func (el *Image) SetY(y int) {
	el.Platform.DrawImage(el.Img, el.X, y, el.Width, el.Height)
}

func (el *Image) SetXY(x, y int) {
	el.Platform.DrawImage(el.Img, x, y, el.Width, el.Height)
}
