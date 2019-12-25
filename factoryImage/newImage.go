package factoryImage

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform/abstractType/image"
)

func NewImage(platform, scratchPad iotmaker_platform_IDraw.IDraw, img interface{}, x, y, width, height int, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) *image.Image {
	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.Set(x)
	x = densityCalc.Int()

	densityCalc.Set(y)
	y = densityCalc.Int()

	densityCalc.Set(width)
	width = densityCalc.Int()

	densityCalc.Set(height)
	height = densityCalc.Int()

	ret := &image.Image{
		Platform:   platform,
		ScratchPad: scratchPad,
		Img:        img,
		X:          x,
		Y:          y,
		Width:      width,
		Height:     height,
	}
	ret.Crete()

	return ret
}
