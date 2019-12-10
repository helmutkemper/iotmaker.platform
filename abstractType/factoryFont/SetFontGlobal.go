package factoryFont

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/font"
)

func SetFontGlobal(platform iotmaker_platform_IDraw.IDraw, font font.Font) {
	platform.Font(font)
}
