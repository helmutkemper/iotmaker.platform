package factoryText

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	iotmaker_platform_textMetrics "github.com/helmutkemper/iotmaker.platform.textMetrics"
)

func NewMeasureText(platform iotmaker_platform_IDraw.IDraw, text string) iotmaker_platform_textMetrics.TextMetrics {
	return platform.MeasureText(text)
}
