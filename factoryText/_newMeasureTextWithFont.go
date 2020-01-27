package factoryText

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	iotmaker_platform_textMetrics "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.textMetrics"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/font"
)

func NewMeasureTextWithFont(platform iotmaker_platform_IDraw.IDraw, text string, font font.Font) iotmaker_platform_textMetrics.TextMetrics {
	platform.Font(font)
	return platform.MeasureText(text)
}
