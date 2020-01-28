package factoryFont

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/font"
)

// en: set a global font for canvas element (based on web browser canvas)
//     Note: this method don't set font color for text
//
// pt_br: define uma fonte global para o elemento canvas (baseado no comportamento do
//     navegador web)
func SetGlobal(platform iotmaker_platform_IDraw.IDraw, font font.Font) {
	platform.Font(font)
}
