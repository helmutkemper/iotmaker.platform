package factoryShadow

import iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"

func ResetStylesGlobal(platform iotmaker_platform_IDraw.IDraw) {
	platform.ResetFillStyle()
	platform.ResetStrokeStyle()
	platform.ResetShadow()
}
