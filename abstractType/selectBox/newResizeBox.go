package selectBox

import (
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform/abstractType/basicBox"
	"github.com/helmutkemper/iotmaker.platform/abstractType/genericTypes"
)

func NewResizeBoxFromBasicBob(basicBox *basicBox.BasicBox, offsetX, offsetY, width, height, border int, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) *ResizeBoxes {

	dm := genericTypes.Dimensions{}
	dm = genericTypes.NewDimensions(dm, offsetX, offsetY, width, height, border, density, iDensity)

	rb := &ResizeBoxes{
		Platform:               basicBox.Platform,
		ScratchPad:             basicBox.ScratchPad,
		Dimensions:             dm,
		FatherOutBoxDimensions: basicBox.OutBoxDimensions,
	}
	rb.Create()

	return rb
}
