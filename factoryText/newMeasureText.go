package factoryText

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	iotmaker_platform_textMetrics "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.textMetrics"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/font"
)

func NewMeasureText(

	platform iotmaker_platform_IDraw.IDraw,

	labelFont font.Font,
	label string,

) iotmaker_platform_textMetrics.TextMetrics {

	platform.Save()
	platform.Font(labelFont)
	metrics := platform.MeasureText(label)
	platform.Restore()

	return metrics
}
