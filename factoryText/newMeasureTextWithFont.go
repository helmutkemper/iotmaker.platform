package factoryText

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	iotmaker_platform_textMetrics "github.com/helmutkemper/iotmaker.platform.textMetrics"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/font"
)

func NewMeasureTextWithFont(platform iotmaker_platform_IDraw.IDraw, text string, font font.Font) iotmaker_platform_textMetrics.TextMetrics {
	platform.Font(font)
	return platform.MeasureText(text)
}
