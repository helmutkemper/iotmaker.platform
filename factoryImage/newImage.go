package factoryImage

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/genericTypes"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/image"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/basic"
	stageEngine "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/engine"
)

func NewImage(id string, engine stageEngine.IEngine, platform, scratchPad iotmaker_platform_IDraw.IDraw, img interface{}, x, y, width, height int, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) *image.Image {
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
			Engine:     engine,
			Platform:   platform,
			ScratchPad: scratchPad,
			Dimensions: genericTypes.Dimensions{
				X:      float64(x),
				Y:      float64(y),
				Width:  float64(width),
				Height: float64(height),
			},
			OutBoxDimensions: genericTypes.Dimensions{
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
