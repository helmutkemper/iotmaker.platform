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
	el.Draw()
}

func (el *Image) Draw() {
	el.Platform.DrawImage(el.Img, el.X, el.Y, el.Width, el.Height)
}

func (el *Image) Clear() {
	el.Platform.ClearRect(el.X, el.Y, el.Width, el.Height)
}

func (el *Image) SetX(x float64) {
	el.X = x
}

func (el *Image) SetY(y float64) {
	el.Y = y
}

func (el *Image) SetXY(x, y float64) {
	el.X = x
	el.Y = y
}
