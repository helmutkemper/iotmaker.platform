package factoryText

import (
	global "github.com/helmutkemper/iotmaker.santa_isabel_theater.globalConfig"
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	iotmaker_platform_textMetrics "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.textMetrics"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/font"
)

func NewMeasureText(

	labelFont font.Font,
	label string,

) iotmaker_platform_textMetrics.TextMetrics {

	var platform iotmaker_platform_IDraw.IDraw
	platform = global.Global.Canvas
	platform.Save()
	platform.Font(labelFont)
	metrics := platform.MeasureText(label)
	platform.Restore()

	return metrics
}
