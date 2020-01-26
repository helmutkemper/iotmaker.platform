package factoryImage

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.interfaces/iStage"
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/image"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/basic"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/dimensions"
)

func NewImage(id string, stage iStage.IStage, platform, scratchPad iotmaker_platform_IDraw.IDraw, img interface{}, x, y, width, height int, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) *image.Image {
	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.Set(float64(x))
	x = densityCalc.Int()

	densityCalc.Set(float64(y))
	y = densityCalc.Int()

	densityCalc.Set(float64(width))
	width = densityCalc.Int()

	densityCalc.Set(float64(height))
	height = densityCalc.Int()

	ret := &image.Image{
		Sprite: basic.Sprite{
			Id:         id,
			Stage:      stage,
			Platform:   platform,
			ScratchPad: scratchPad,
			Dimensions: dimensions.Dimensions{
				X:      float64(x),
				Y:      float64(y),
				Width:  float64(width),
				Height: float64(height),
			},
			OutBoxDimensions: dimensions.Dimensions{
				X:      float64(x),
				Y:      float64(y),
				Width:  float64(width),
				Height: float64(height),
			},
		},
		Img: img,
	}
	ret.Crete()

	return ret
}
