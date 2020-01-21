package factoryImage

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/genericTypes"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/image"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/basic"
	stageEngine "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/engine"
)

func NewImageWithDelta(id string, engine stageEngine.IEngine, platform, scratchPad iotmaker_platform_IDraw.IDraw, img interface{}, x, y, width, height, xDelta, yDelta float64, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) *image.Image {
	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.Set(x)
	x = densityCalc.Float64()

	densityCalc.Set(y)
	y = densityCalc.Float64()

	densityCalc.Set(width)
	width = densityCalc.Float64()

	densityCalc.Set(height)
	height = densityCalc.Float64()

	densityCalc.Set(xDelta)
	xDelta = densityCalc.Float64()

	densityCalc.Set(yDelta)
	yDelta = densityCalc.Float64()

	ret := &image.Image{
		Sprite: basic.Sprite{
			Id:         id,
			Engine:     engine,
			Platform:   platform,
			ScratchPad: scratchPad,
			Dimensions: genericTypes.Dimensions{
				X:      x,
				Y:      y,
				Width:  width,
				Height: height,
			},
			OutBoxDimensions: genericTypes.Dimensions{
				X:      x,
				Y:      y,
				Width:  width,
				Height: height,
			},

			MovieDeltaX: int(xDelta),
			MovieDeltaY: int(yDelta),
		},
		Img: img,
	}
	ret.Crete()

	return ret
}
