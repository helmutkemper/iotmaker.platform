package factoryShadow

import iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"

func ResetShadowGlobal(platform iotmaker_platform_IDraw.IDraw) {
	platform.ShadowReset()
}
